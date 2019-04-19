package main

import (
  "context"
  "fmt"
  "github.com/olivere/elastic"
)

const (
	indexName = "book"
	docType   = "default"
)

type Book struct {
  book_name  string  `json:"book_name"`
}

// InsertProduct ...
func InsertProduct(ctx context.Context, elasticClient *elastic.Client) error {

    b := Book {
      book_name : "The GO Programming Language",
    }

    _, err := elasticClient.Index().
			Index(indexName).
			Type(docType).
			BodyJson(b).
			Do(ctx)

		if err != nil {
			return err
		}

    return nil
}

func main() {

  ctx := context.Background()

  // init Elastic client
  elasticClient, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
  if err != nil {
    fmt.Println("Error : ", err.Error())
  }

  // Insert product
  fmt.Println("+++++ Insert +++++")
  InsertProduct(ctx, elasticClient)

}
