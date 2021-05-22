package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/buger/jsonparser"
)

func parseIncomingRequest(httpResp *http.Response) (string, error) {

	bodyBytes, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		panic(err)
	}

	body, err := jsonparser.GetString(bodyBytes, "result", "[0]", "message", "text")
	if err != nil {
		log.Fatal("Error in parsing data")
	}

	return body, nil
}

func main() {

	httpreq, err := http.Get("https://api.telegram.org/bot1815331593:AAGM_U2Dw5KxQo3rjTIajSesZvfcj9r_iYw/getUpdates?limit=1")
	if err != nil {
		log.Printf("Error in rerieving request")
	}
	
	parsedData, err := parseIncomingRequest(httpreq)
	if err != nil {
		fmt.Println("Error in parsing retreived data!")
	}

	fmt.Println(parsedData)
}