package main

import (
  "flag"
  "fmt"
  "context"
  "github.com/olivere/elastic"
  //"encoding/json"
  "reflect"
  "strings"
  "time"
)

type dpicollector struct {
  ID   int64                    `json:"id"`
  Dpicollector_name string      `json:"dpicollector_name"`
  Dpicollector_yymmdd string    `json:"dpicollector_yymmdd"`
}

const (
	indexName = "testdpi"
	docType   = "default"
)

const (
    layoutISO = "2006-01-02"
    layoutUS  = "January 2, 2006"
    layout_mmddyy = "01/02/06"
    layout_RFC3339 = "2006-01-02T15:04:05-0700"
)


func init() {



}

func main() {

  // receive day from user
  var day *int

  day = flag.Int("day", 0, "Please specify day for delete older information")

  // receive flag
  flag.Parse()

  dpicollector_day := *day

  fmt.Println("day:", dpicollector_day)

  ctx := context.Background()

	// init Elastic client
	elasticClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}


  // get dpicollector data from Elastic search
  // Search
	//termQuery := elastic.NewTermQuery("dpicollector_yymmdd", "dpicollector")
	searchResult, err := elasticClient.Search().
		Index(indexName).
    Type(docType).
		//Query(termQuery).
		Sort("id", true).
		Pretty(true).       // pretty print request and response JSON
		Do(ctx)             // execute
	if err != nil {
		// Handle error
		panic(err)
	}


  // over iterating the hits, see below.
	var d dpicollector
	for _, item := range searchResult.Each(reflect.TypeOf(d)) {
		if t, ok := item.(dpicollector); ok {
			//fmt.Printf("%d) %s day %s\n",t.ID, t.Dpicollector_name, t.Dpicollector_yymmdd)
      fmt.Println(t)

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

      now := time.Now()
      //nt := now.Format(layout_RFC3339)
      //fmt.Printf("%T", nt)

      diff_days := now.Sub(p).Hours() / 24
      diff := int(diff_days)
      fmt.Printf("%f %d\n", diff_days, diff)


      if (dpicollector_day != 0) {
        if (diff > dpicollector_day) {

          //delete dpi
          fmt.Printf("Delete %s\n", t.Dpicollector_name)

          query := elastic.NewBoolQuery()
          query = query.Must(elastic.NewTermQuery("id", t.ID))

          _, err := elastic.NewDeleteByQueryService(elasticClient).
          Index(indexName).
          Type(docType).
          Query(query).
          Do(ctx)

          if err != nil {
            // Handle error
            panic(err)
          }

        }
      }

		}
	}
	// TotalHits is another convenience function that works even when something goes wrong.
	//fmt.Printf("Found a total of %d \n", searchResult.TotalHits())


}
