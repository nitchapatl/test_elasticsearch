package main

import (
  "time"
  "fmt"
)

const (
    layoutISO = "2006-01-02"
    layoutUS  = "January 2, 2006"
    layout_mmddyy = "01/02/06"
)

func main() {
//date := "1999-12-31"
/*date := "12/31/69"
t, _ := time.Parse(layout_mmddyy, date)
fmt.Println(t)                  // 1999-12-31 00:00:00 +0000 UTC
fmt.Println(t.Format(layoutUS)) // December 31, 1999*/

    p := fmt.Println

    t := time.Now()
    p(t.Format(time.RFC3339))

    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05-0700"))

}
