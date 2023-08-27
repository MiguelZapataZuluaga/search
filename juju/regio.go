package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	/*query := `{
	    "search_type": "match",
	    "query":
	    {

	        "term": "DUJMOVITS"

	    },
	    "from": 0,
	    "max_results": 1,
	    "_source": []
	}`*/
	var keyCharacters string
	query := fmt.Sprintf(`{"query": {"match": {"_all": "%s" }}, "size":10}`, keyCharacters)

	req, err := http.NewRequest("POST", "http://localhost:4080/api/olympics/_search", strings.NewReader(query))
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("admin", "Complexpass#123")
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println(resp.StatusCode)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))
}
