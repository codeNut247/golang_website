package main

import (
	"encoding/json"

	"github.com/codenut247/golang_website/02_jsonapi/app/shared/database"
	"github.com/verifiedninja/webapp/shared/jsonconfig"
	"github.com/verifiedninja/webapp/shared/server"
)

type configuration struct {
	Database database.Info `json:"Database"`
	Server   server.Server `json:"Server"`
}

var config = &configuration{}

func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

func main() {
	jsonconfig.Load("config.json", config)
	database.Connect(config.Database)
	server.Run(httpHandlers, httpsHandlers, s)
}
