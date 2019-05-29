package main

import (
  "context"
  "fmt"
  "github.com/olivere/elastic"
  "strings"
  "github.com/Pallinder/go-randomdata"
  "time"
  "log"
)

type dpicollector struct {
  ID   int64                    `json:"id"`
  Dpicollector_name string      `json:"dpicollector_name"`
  Dpicollector_yymmdd string    `json:"dpicollector_yymmdd"`
  Dpicollector_content string   `json:"dpicollector_content"`
}

func main () {

  start := time.Now()

  ctx := context.Background()

  // init Elastic client
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"))
	if err != nil {
		fmt.Println("Error : ", err.Error())
	}

  var index_name string

  var content string = ""

  content  = content + "6fEN4PmlAVPwvFBd5jVo"
  content  = content + "hY6IyUZPbUbU21Gk6QEe"
  content  = content + "9aLwsDJ88y8HkvnHZjMt"
  content  = content + "J8L615mGdLO7QZcsuvvx"
  content  = content + "PlKOO6GuIt3bfSG009Rg"
  content  = content + "IJGEhVYQLJLXemQIYVrL"
  content  = content + "4GGgpLSxDGcBH7S5LRd6"
  content  = content + "xNhnNmX1dOgSH4crtV7b"
  content  = content + "VQvAHJVFublQkEBsIihu"
  content  = content + "q2kHTooG6LzUA0eySYKE"
  content  = content + "C4yxgI7QTFQ2yxmv6PcC"
  content  = content + "wIZQPzfZnkr1HUwmbAvF"
  content  = content + "S6UOHeUk8clkkG1e5djC"
  content  = content + "swTwLm4miEiDzKFrgj1s"
  content  = content + "DSOfuXATse2d0tisrDMr"
  content  = content + "gGl81Yc7jhvDKAsGCQIb"
  content  = content + "U0JsNcXERnJwMT2cZIYb"
  content  = content + "6hP5vKln47iJWNgnTjkb"
  content  = content + "ujhPZjVQTu8wJPJf3P4X"
  content  = content + "OD66bzbyAoO27YF5SNFn"
  content  = content + "ui10WhQnMXyUKlRfwuCG"
  content  = content + "d2WXv38jAcoiB1MC0ogC"
  content  = content + "aQFIezQAlW9J2kqFGpcl"
  content  = content + "EanoejBgTcEt3mPj58Qs"
  content  = content + "12nMJtVopzoUmyh2rlvV"
  content  = content + "DBOknssTm5i40ts7HdZv"
  content  = content + "X22Eoa0o9dyk92S7uu49"
  content  = content + "7Zc2CxDyC1XD0bGpD6aE"
  content  = content + "rbVgirZ1rJAMEprc5uVF"
  content  = content + "5l6IrID0WJsl8TjqVRWK"
  content  = content + "opCU7knDY1qGgjON799M"
  content  = content + "X6BAPRKaxE8NVsbUswud"
  content  = content + "OagTNhboG5bVDFf1sDWF"
  content  = content + "KboKiv1ov661t6LOejb2"
  content  = content + "42S2HaD0QG5VCMi6sv2l"
  content  = content + "meAjW6gTRFQTQ4RM0lfP"
  content  = content + "teFzFcFWxyq9msk7jgDe"
  content  = content + "zxRhRifgm9S2cBGJZ1F1"
  content  = content + "s9s5VfMqnqQSVCJJtam2"
  content  = content + "ENtFYky7PjeG6ZTTorlf"
  content  = content + "H3lAoa6TO1CGot9Ut3cl"
  content  = content + "njDfX2i8HCTHKhpKt9zF"
  content  = content + "Gf4qpODCbkPwzG80J2ay"
  content  = content + "O2hbGmUFRrx4XQ4Crzym"
  content  = content + "w29MWBeURA1b8ycwOqPY"
  content  = content + "6VN5Xlzrc5NIitOzMNoF"
  content  = content + "iNEgWnfzTZ7xUaAnpvMq"
  content  = content + "wa4xd1JSaBi90nAfOtIY"
  content  = content + "WGlewUmcsq0Ph822Zfoz"
  content  = content + "Ib5n4OBokUAEhBTrN3fR"
  content  = content + "SsbjVGI10gfU4KZF6g8x"
  content  = content + "vzOesDMA3hhXMrVwn7Ca"
  content  = content + "38KPaPgnmgQIGxxUtMjX"
  content  = content + "IeZcRGZ2dW5ibhdE5T2Q"
  content  = content + "ENrznoL9CfHiBFePm3Bb"
  content  = content + "EFEkSQnP9R7cfPO0jVGT"
  content  = content + "rIFW7qAC9fW5EQzArnBq"
  content  = content + "sA4qdFK2KygCAA6kOvbN"
  content  = content + "HAynFLkN5FDL06LC1n4p"
  content  = content + "Wq9PaJoCAtgBL2e6qIig"
  content  = content + "aoYeFcPO7LzlHAr1H8XT"
  content  = content + "Eum2dG3Z1sOLf2DNFfwS"
  content  = content + "QemWtRylkLNwoIFJtOGz"
  content  = content + "9YPIRPVDiGkRF2vqUdUP"
  content  = content + "80BMrwmRYRFpX6zUqxST"
  content  = content + "4H58cflxJhUqUoMbTUGE"
  content  = content + "JvwOPDpvn5V4cbBfeZp6"
  content  = content + "H27vVpepS81ATmskFbmE"
  content  = content + "CJGbVXJwn7PjLzuNJMfa"
  content  = content + "ixsfS0f0pxXKXRvMyY2o"
  content  = content + "S6mL0lq0TnXDo1seCnuO"
  content  = content + "xUBwtbmVLlDMcZxLXuZT"
  content  = content + "MekdfYEaljAdTHnpX2a2"
  content  = content + "EbIpyKmdSz9cHc82suqQ"
  content  = content + "8pfbjeSsxiK3aOGQxG5A"
  content  = content + "U6qDZ4QGfvhDjvKljLJo"
  content  = content + "GJe4rL7Zy0gI5mDG3Ozg"
  content  = content + "XDqZj6HkrQucs2zwp9TN"
  content  = content + "LeeGnabLQ4IVitLNDGfF"
  content  = content + "5Sk7fmoH0FyeEks2hayA"
  content  = content + "s0M0BF04edXHM0O2wHzM"
  content  = content + "qecwBB78pbLtHIQaHpry"
  content  = content + "OBc83M5wst1NW5tRGOEw"
  content  = content + "3A02jJyhuiNOx8NDPU3g"
  content  = content + "QbZsDNj8Qupxqfgtxlb3"
  content  = content + "ynWJGhE61siihCIL4Jv1"
  content  = content + "dc9TSxgwvsTfzbS0qvuV"
  content  = content + "OxgYAvnkEKI2fTo6vOC3"
  content  = content + "13aPTfE3a8bkvZeJAAQX"
  content  = content + "wRaXO5fGaCsSSrzynOWP"
  content  = content + "Ib2crrIDMX73IN2tRacE"
  content  = content + "4NQG239lVLbOAinmKW5L"
  content  = content + "heLHG9tqc4NHVSIRdGjD"
  content  = content + "sATIBYLTZfIjxoK1IDxY"
  content  = content + "yfjiEpQhYwxESFT8KkAR"
  content  = content + "L9zlTcmmrqldEByHOepJ"
  content  = content + "POMjpdAOrnsrrm8EEPuG"
  content  = content + "OVpuW5gjzK5scouew7r8"
  content  = content + "iKGpTSJ7L2zGNhfBZjW4"
  content  = content + "EHdmLCQFVwurdQFOH5sr"
  content  = content + "6fEN4PmlAVPwvFBd5jVo"
  content  = content + "hY6IyUZPbUbU21Gk6QEe"
  content  = content + "9aLwsDJ88y8HkvnHZjMt"
  content  = content + "J8L615mGdLO7QZcsuvvx"
  content  = content + "PlKOO6GuIt3bfSG009Rg"
  content  = content + "IJGEhVYQLJLXemQIYVrL"
  content  = content + "4GGgpLSxDGcBH7S5LRd6"
  content  = content + "xNhnNmX1dOgSH4crtV7b"
  content  = content + "VQvAHJVFublQkEBsIihu"
  content  = content + "q2kHTooG6LzUA0eySYKE"
  content  = content + "C4yxgI7QTFQ2yxmv6PcC"
  content  = content + "wIZQPzfZnkr1HUwmbAvF"
  content  = content + "S6UOHeUk8clkkG1e5djC"
  content  = content + "swTwLm4miEiDzKFrgj1s"
  content  = content + "DSOfuXATse2d0tisrDMr"
  content  = content + "gGl81Yc7jhvDKAsGCQIb"
  content  = content + "U0JsNcXERnJwMT2cZIYb"
  content  = content + "6hP5vKln47iJWNgnTjkb"
  content  = content + "ujhPZjVQTu8wJPJf3P4X"
  content  = content + "OD66bzbyAoO27YF5SNFn"
  content  = content + "ui10WhQnMXyUKlRfwuCG"
  content  = content + "d2WXv38jAcoiB1MC0ogC"
  content  = content + "aQFIezQAlW9J2kqFGpcl"
  content  = content + "EanoejBgTcEt3mPj58Qs"
  content  = content + "12nMJtVopzoUmyh2rlvV"
  content  = content + "DBOknssTm5i40ts7HdZv"
  content  = content + "X22Eoa0o9dyk92S7uu49"
  content  = content + "7Zc2CxDyC1XD0bGpD6aE"
  content  = content + "rbVgirZ1rJAMEprc5uVF"
  content  = content + "5l6IrID0WJsl8TjqVRWK"
  content  = content + "opCU7knDY1qGgjON799M"
  content  = content + "X6BAPRKaxE8NVsbUswud"
  content  = content + "OagTNhboG5bVDFf1sDWF"
  content  = content + "KboKiv1ov661t6LOejb2"
  content  = content + "42S2HaD0QG5VCMi6sv2l"
  content  = content + "meAjW6gTRFQTQ4RM0lfP"
  content  = content + "teFzFcFWxyq9msk7jgDe"
  content  = content + "zxRhRifgm9S2cBGJZ1F1"
  content  = content + "s9s5VfMqnqQSVCJJtam2"
  content  = content + "ENtFYky7PjeG6ZTTorlf"
  content  = content + "H3lAoa6TO1CGot9Ut3cl"
  content  = content + "njDfX2i8HCTHKhpKt9zF"
  content  = content + "Gf4qpODCbkPwzG80J2ay"
  content  = content + "O2hbGmUFRrx4XQ4Crzym"
  content  = content + "w29MWBeURA1b8ycwOqPY"
  content  = content + "6VN5Xlzrc5NIitOzMNoF"
  content  = content + "iNEgWnfzTZ7xUaAnpvMq"
  content  = content + "wa4xd1JSaBi90nAfOtIY"
  content  = content + "WGlewUmcsq0Ph822Zfoz"
  content  = content + "Ib5n4OBokUAEhBTrN3fR"
  content  = content + "SsbjVGI10gfU4KZF6g8x"
  content  = content + "vzOesDMA3hhXMrVwn7Ca"
  content  = content + "38KPaPgnmgQIGxxUtMjX"
  content  = content + "IeZcRGZ2dW5ibhdE5T2Q"
  content  = content + "ENrznoL9CfHiBFePm3Bb"
  content  = content + "EFEkSQnP9R7cfPO0jVGT"
  content  = content + "rIFW7qAC9fW5EQzArnBq"
  content  = content + "sA4qdFK2KygCAA6kOvbN"
  content  = content + "HAynFLkN5FDL06LC1n4p"
  content  = content + "Wq9PaJoCAtgBL2e6qIig"
  content  = content + "aoYeFcPO7LzlHAr1H8XT"
  content  = content + "Eum2dG3Z1sOLf2DNFfwS"
  content  = content + "QemWtRylkLNwoIFJtOGz"
  content  = content + "9YPIRPVDiGkRF2vqUdUP"
  content  = content + "80BMrwmRYRFpX6zUqxST"
  content  = content + "4H58cflxJhUqUoMbTUGE"
  content  = content + "JvwOPDpvn5V4cbBfeZp6"
  content  = content + "H27vVpepS81ATmskFbmE"
  content  = content + "CJGbVXJwn7PjLzuNJMfa"
  content  = content + "ixsfS0f0pxXKXRvMyY2o"
  content  = content + "S6mL0lq0TnXDo1seCnuO"
  content  = content + "xUBwtbmVLlDMcZxLXuZT"
  content  = content + "MekdfYEaljAdTHnpX2a2"
  content  = content + "EbIpyKmdSz9cHc82suqQ"
  content  = content + "8pfbjeSsxiK3aOGQxG5A"
  content  = content + "U6qDZ4QGfvhDjvKljLJo"
  content  = content + "GJe4rL7Zy0gI5mDG3Ozg"
  content  = content + "XDqZj6HkrQucs2zwp9TN"
  content  = content + "LeeGnabLQ4IVitLNDGfF"
  content  = content + "5Sk7fmoH0FyeEks2hayA"
  content  = content + "s0M0BF04edXHM0O2wHzM"
  content  = content + "qecwBB78pbLtHIQaHpry"
  content  = content + "OBc83M5wst1NW5tRGOEw"
  content  = content + "3A02jJyhuiNOx8NDPU3g"
  content  = content + "QbZsDNj8Qupxqfgtxlb3"
  content  = content + "ynWJGhE61siihCIL4Jv1"
  content  = content + "dc9TSxgwvsTfzbS0qvuV"
  content  = content + "OxgYAvnkEKI2fTo6vOC3"
  content  = content + "13aPTfE3a8bkvZeJAAQX"
  content  = content + "wRaXO5fGaCsSSrzynOWP"
  content  = content + "Ib2crrIDMX73IN2tRacE"
  content  = content + "4NQG239lVLbOAinmKW5L"
  content  = content + "heLHG9tqc4NHVSIRdGjD"
  content  = content + "sATIBYLTZfIjxoK1IDxY"
  content  = content + "yfjiEpQhYwxESFT8KkAR"
  content  = content + "L9zlTcmmrqldEByHOepJ"
  content  = content + "POMjpdAOrnsrrm8EEPuG"
  content  = content + "OVpuW5gjzK5scouew7r8"
  content  = content + "iKGpTSJ7L2zGNhfBZjW4"
  content  = content + "EHdmLCQFVwurdQFOH5sr"

  for i := 1 ; i <= 5; i ++ {


    index_name = fmt.Sprintf("dpicollector%d", i)

    fmt.Printf("%s\n", index_name)

    createIndex, err := client.CreateIndex(index_name).Do(ctx)
    if err != nil {
      // Handle error
      panic(err)
    }
    if !createIndex.Acknowledged {
      // Not acknowledged
    }

    // create sample data
    num := randomdata.Number(30000, 50000)
    fmt.Println(num)


    for j := 1 ; j <= num ; j++ {

      dpicollector_name := fmt.Sprintf("%s_%d", index_name, j)
      //fmt.Println(dpicollector_name)

      sample_date1 := randomdata.FullDateInRange("2019-01-01", "2019-05-31")
      //fmt.Printf("%s\n", sample_date1)

      sample_date2, _ := time.Parse("Monday 2 Jan 2006", sample_date1)
      //fmt.Println(sample_date2)

      sample_date := sample_date2.Format("01/02/06")
      //fmt.Println(sample_date)

      s := strings.Split(sample_date, "/")
      dd, mm, yy := s[1], s[0], s[2]
      //fmt.Printf("%s %s %s\n", dd, mm, yy)

      dpicollector_yymmdd := index_name + "_" + yy + mm + dd
      //fmt.Println(dpicollector_yymmdd)

      data := dpicollector{int64(j), dpicollector_name, dpicollector_yymmdd, content}
      //fmt.Println(data)

      // Insert dpicollector
      _, err := client.Index().
  			Index(index_name).
  			Type("default").
  			BodyJson(data).
  			Do(ctx)

  		if err != nil {
  			panic(err)
  		}


    }

  }

  t := time.Now()
  elapsed := t.Sub(start)
  log.Printf("use time %s", elapsed)

}
