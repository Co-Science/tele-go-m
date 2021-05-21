package main

import (
	"encoding/json"
	// "errors"
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

type k struct {
	// this must always start in capital case to export as json
	Ok bool `json:"ok"`
}


func main() {

	url := "https://raw.githubusercontent.com/Co-Science/tele-go-m/Joel-Nickson-patch-1/test.json"

	resp, err := http.Get(url)
	if err != nil {
	   log.Fatalln(err)
	}

	 body, err := ioutil.ReadAll(resp.Body)
   if err != nil {
      log.Fatalln(err)
   }
	//Convert the body to type string
   sb := "["+string(body)+"]"

   fmt.Println(sb)
   bytes := []byte(sb)
 	
   var json_in_go []k
   json.Unmarshal(bytes, &json_in_go)


   fmt.Println(json_in_go)
   for l := range json_in_go {
        fmt.Printf("Id = %v ", json_in_go[l].Ok)
        fmt.Println()
    }

fmt.Println("\n\n:)")
}
