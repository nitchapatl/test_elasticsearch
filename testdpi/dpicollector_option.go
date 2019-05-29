package main

import (
  "flag"
  "fmt"
  "context"
  "github.com/olivere/elastic"
  "strings"
  "net/http"
  "io/ioutil"
  "strconv"
  "math"
  "reflect"
  "time"
)

type dpicollector struct {
  ID   int64                    `json:"id"`
  Dpicollector_name string      `json:"dpicollector_name"`
  Dpicollector_yymmdd string    `json:"dpicollector_yymmdd"`
  Dpicollector_content string   `json:"dpicollector_content"`
}

const (
	docType   = "default"
)

func main() {

  // receive data retention option
  var day *int
  var size *int

  day = flag.Int("day", 0, "Please specify day for collect information")
  size = flag.Int("size", 0, "Please specify storage size for collect information")

  // receive flag
  flag.Parse()

  dpicollector_day := *day
  dpicollector_size := *size

  fmt.Println("day:", dpicollector_day)
  fmt.Println("size:", dpicollector_size)

  var sIndexNames []string

  now := time.Now()

  ctx := context.Background()

  // init client
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

  index_names, err := client.IndexNames()
  if err != nil {
    // Handle error
    panic(err)
  }

  for _, index_name := range index_names {

    if (strings.Contains(index_name, "dpicollector")) {
      fmt.Printf("%s\n", index_name)

      sIndexNames = append(sIndexNames, index_name)

    }

  }

  // data retention condition
  if dpicollector_day == 0 && dpicollector_size == 0 {
    // nothing


  } else if dpicollector_day != 0 && dpicollector_size == 0 {

    fmt.Printf("data retention option: %d day\n", dpicollector_day)

    for _, index_name := range sIndexNames {
      //fmt.Println(index_name)

      if err := deleteDataRetention(client, ctx, index_name, dpicollector_day, now); err != nil {
        fmt.Println("Error : ", err.Error())
      }

    }

  } else if dpicollector_day == 0 && dpicollector_size != 0 {

    fmt.Printf("data retention option: %d MB\n", dpicollector_size)

    for _, index_name := range sIndexNames {

      // index storage size
      store_size_mb, err := getStoreSizeMB(index_name)
      if err != nil {
        fmt.Println("Error : ", err.Error())
      }

      //fmt.Println(store_size_mb)

      if store_size_mb > float64(dpicollector_size) {
        if err := deleteDataRetention(client, ctx, index_name, 30, now); err != nil {
          fmt.Println("Error : ", err.Error())
        }


      }

    }

  } else if dpicollector_day != 0 && dpicollector_size != 0 {

    fmt.Printf("data retention option: %d days, %d MB\n", dpicollector_day, dpicollector_size)

    for _, index_name := range sIndexNames {

      // index storage size
      store_size_mb, err := getStoreSizeMB(index_name)
      if err != nil {
        fmt.Println("Error : ", err.Error())
      }

      //fmt.Println(store_size_mb)

      if store_size_mb > float64(dpicollector_size) {
        if err := deleteDataRetention(client, ctx, index_name, dpicollector_day, now); err != nil {
          fmt.Println("Error : ", err.Error())
        }


      }

    }

  }

}

func buildURL(index_name string) string {
   return "http://localhost:9200/_cat/indices/" + index_name + "?h=store.size&bytes=b"
}

func buildURLDocsCnt(index_name string) string {
   return "http://localhost:9200/_cat/indices/" + index_name + "?h=docs.count"
}

 func mByteFormat(inputNum float64, precision int) float64 {

         if precision <= 0 {
                 precision = 1
         }

         //var unit string
         var returnVal float64

         returnVal = round((inputNum / 1048576), precision)
         //unit = " MB" // megabyte*/

         return returnVal

 }

 func getStoreSizeMB(index_name string) (store_size_mb float64, err error) {

   //URL Encoding
   URL := buildURL(index_name)

   resp, _ := http.Get(URL)

   defer resp.Body.Close()

   // reads html as a slice of bytes
   html, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      return 0, err
   }

   htmls := string(html)

   s := strings.TrimSuffix(htmls, "\n")

   store_size,err := strconv.ParseFloat(s, 64);
   if err != nil {
     return 0, err
   }
   store_size_mb = mByteFormat(store_size, 1)

   return store_size_mb, nil
 }

 func getDocsCount(index_name string) (docs_count int, err error) {

   //URL Encoding
   URL := buildURLDocsCnt(index_name)

   resp, _ := http.Get(URL)

   defer resp.Body.Close()

   // reads html as a slice of bytes
   html, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      return 0, err
   }

   htmls := string(html)

   s := strings.TrimSuffix(htmls, "\n")

   docs_count, err = strconv.Atoi(s)
   if err != nil {
     return 0, err
   }

   return docs_count, nil
 }

 func round(input float64, places int) (newVal float64) {
         var round float64
         pow := math.Pow(10, float64(places))
         digit := pow * input
         round = math.Floor(digit)
         newVal = round / pow
         return newVal
 }

 func deleteDataRetention(client *elastic.Client, ctx context.Context, index_name string, dpicollector_day int, now time.Time) error {

   db_size, _ := getDocsCount(index_name)

   searchResult, err := client.Search().
     Index(index_name).
     Type(docType).
     Sort("id", true).
     Pretty(true).
     Size(db_size).
     Do(ctx)
   if err != nil {
     return err
   }

   // over iterating the hits, see below.
   var d dpicollector
   for _, item := range searchResult.Each(reflect.TypeOf(d)) {
     if t, ok := item.(dpicollector); ok {
       //fmt.Printf("%d) %s day %s\n",t.ID, t.Dpicollector_name, t.Dpicollector_yymmdd)
       //fmt.Println(t)

       //fmt.Printf("%q\n", strings.Split(t.Dpicollector_yymmdd, "_"))
       s := strings.Split(t.Dpicollector_yymmdd, "_")
       //fmt.Println(s[1])

       s_day := s[1]
       //fmt.Println(s_day)

       yy := s_day[0:2]
       //fmt.Printf("%s\n", yy)

       mm := s_day[2:4]
       //fmt.Printf("%s\n", mm)

       dd := s_day[4:6]
       //fmt.Printf("%s\n", dd)

       since_day := mm + "/" + dd + "/" + yy
       fmt.Printf("%s\n", since_day)
       p, _ := time.Parse("01/02/06", since_day)

       diff_days := now.Sub(p).Hours() / 24
       diff := int(diff_days)
       fmt.Printf("%f %d\n", diff_days, diff)


       if (dpicollector_day != 0) {
         if (diff > dpicollector_day) {

           //delete dpi
           fmt.Printf("Delete %s\n", t.Dpicollector_name)

           query := elastic.NewBoolQuery()
           query = query.Must(elastic.NewTermQuery("id", t.ID))

           _, err := elastic.NewDeleteByQueryService(client).
           Index(index_name).
           Type(docType).
           Query(query).
           Do(ctx)

           if err != nil {
             // Handle error
             return err
           }

         }
       }
     }
   }

   return nil
 }
