package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type ApiUser struct {
	AuthToken string
	Email     string
	Pwd       string
}

var (
	requestMethod int
)

func init() {
	flag.IntVar(&requestMethod, "requmeth", 1, "an integer")
}

func main() {
	u := ApiUser{AuthToken: "oa4i§&2P1OJN%§5I%P", Email: "hello@chillster.com", Pwd: "blodedSchlumSchlumSchlipediDum"}
	switch requestMethod {
	case 1:
		PostMeth("http://localhost:8080/api/dog", u)
	case 2:
		PostMeth("http://localhost:8080/api/dog", u)
	}

}

func PostMeth(URL string, SendStruct interface{}) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(SendStruct)
	res, err := http.Post(URL, "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Println(err)
	}

	io.Copy(os.Stdout, res.Body)
}

// func RequestMeth(URL string, SendStruct interface{}) {
// 	client := &http.Client{}
// 	v := url.Values{}
// 	v.Set("name", name)
// 	//pass the values to the request's body
// 	req, err := http.NewRequest("POST", URL, strings.NewReader(v.Encode()))
// 	req.SetBasicAuth(EMAIL, PASSWORD)
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	bodyText, err := ioutil.ReadAll(resp.Body)
// 	s := string(bodyText)
// }
