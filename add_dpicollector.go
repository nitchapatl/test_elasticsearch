package main

import (
  "context"
  "fmt"
  "github.com/olivere/elastic"
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

func main() {

  ctx := context.Background()

	// init Elastic client
	elasticClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

  add_dpi := []dpicollector{}

  //setting data
  add_dpi = append(add_dpi, dpicollector{1, "dpicollector_1", "dpicollector_190101"})
  add_dpi = append(add_dpi, dpicollector{2, "dpicollector_2", "dpicollector_180224"})
  add_dpi = append(add_dpi, dpicollector{3, "dpicollector_3", "dpicollector_190512"})
  add_dpi = append(add_dpi, dpicollector{4, "dpicollector_4", "dpicollector_190413"})
  add_dpi = append(add_dpi, dpicollector{5, "dpicollector_5", "dpicollector_190401"})
  //add_dpi = append(add_dpi, dpicollector{6, "dpicollector_6", "dpicollector_151112"})
  //add_dpi = append(add_dpi, dpicollector{7, "dpicollector_7", "dpicollector_190326"})
  //add_dpi = append(add_dpi, dpicollector{8, "dpicollector_8", "dpicollector_190420"})
  //add_dpi = append(add_dpi, dpicollector{9, "dpicollector_9", "dpicollector_190226"})
  //add_dpi = append(add_dpi, dpicollector{10, "dpicollector_10", "dpicollector_180627"})

  //fmt.Println(add_dpi)


  for _,dpicollector := range add_dpi{
        //fmt.Println(index)
        //fmt.Println(dpicollector)

        // Insert dpicollector
        fmt.Println("+++++ Insert %s \n", dpicollector.Dpicollector_name)
        InsertDPI(ctx, elasticClient, dpicollector)

    }
}

// Insert dpicollector
func InsertDPI(ctx context.Context, elasticClient *elastic.Client, data dpicollector) error {


    _, err := elasticClient.Index().
			Index(indexName).
			Type(docType).
			BodyJson(data).
			Do(ctx)

		if err != nil {
			return err
		}

    return nil
}
