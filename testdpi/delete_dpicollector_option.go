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
  Dpicollector_content string   `json:"dpicollector_content"`
}

func main () {

  ctx := context.Background()

  // init Elastic client
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

  var index_name string

  for i := 1 ; i <= 5; i ++ {


    index_name = fmt.Sprintf("dpicollector%d", i)

    //fmt.Printf("%s\n", index_name)

    // Delete an index
	  deleteIndex, err := client.DeleteIndex(index_name).Do(ctx)
	  if err != nil {
		  // Handle error
		  panic(err)
	  }
	  if !deleteIndex.Acknowledged {
		  // Not acknowledged
      fmt.Printf("Can not delete index %s\n", index_name)
	  }

  }

}
