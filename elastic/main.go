package main

import (
       "encoding/json"
       "fmt"

	elastic "gopkg.in/olivere/elastic.v2"

)

func main() {

     // Create a client
     client, err := elastic.NewClient()
     if err != nil {
        // Handle error
     }

     res, err := client.Get().Index("falcor").Type("beach").Id("1").Do()
     if err != nil {
        fmt.Println(err)
        return
     }

     fmt.Println(res.Source)
     j, err := json.Marshal(&res.Source)
     if err != nil {
        fmt.Println(err)
     }

     fmt.Println(string(j))

}
