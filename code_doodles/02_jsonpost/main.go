package main

import (
	"bytes"
	"encoding/json"
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

func main() {
	u := ApiUser{AuthToken: "oa4i§&2P1OJN%§5I%P", Email: "hello@chillster.com", Pwd: "blodedSchlumSchlumSchlipediDum"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	res, err := http.Post("http://localhost:8080/api/dog", "application/json; charset=utf-8", b)
	if err != nil {
		fmt.Println(err)
	}
	io.Copy(os.Stdout, res.Body)
}
