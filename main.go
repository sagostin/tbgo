package main

import (
	"bytes"
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"tbgo/sbc"
)

/** INFO
 */

func main() {

	// init cli flags
	var fHost = flag.String("host", "", "specify the telcobridges host api")
	if fHost == nil {
		log.Fatalf("no api host provided")
	}

	var fUsername = flag.String("username", "", "telcobridges api username")
	if fUsername == nil {
		log.Fatalf("no username provided")
	}
	var fPassword = flag.String("password", "", "telcobridges api password")
	if fPassword == nil {
		log.Fatalf("no password provided")
	}

	flag.Parse()

	// change pointer to non to be able to compare
	apiUsername := *fUsername
	apiPassword := *fPassword
	apiHost := *fHost

	cfg := sbc.NewClientConfig()

	if apiUsername != "" {
		cfg.APIUsername = apiUsername
	}

	if apiPassword != "" {
		cfg.APIPassword = apiPassword
	}

	if apiHost != "" {
		cfg.APIHost = apiHost
	}

	// init the http client constructor thingy 🤪
	client := sbc.NewClient(cfg)

	var nap sbc.NapStatus

	err := client.Request("GET", "/configurations/config_1/naps/pbx_dec0de/status", nil, nil)
	if err != nil {
		log.Error(err)
	}

	marshal, err := json.Marshal(nap)
	if err != nil {
		log.Error(err)
	}

	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, marshal, "", "\t")
	if error != nil {
		log.Error(err)
	}
	log.Printf("\n" + string(prettyJSON.Bytes()))
}
