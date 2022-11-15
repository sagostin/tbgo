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

	client.TBFileDBs().Get("config_1")

	/*	_, err := client.TBConfigs().GetConfig("config_1")
		if err != nil {
			return
		}

		naps, err := client.TBNaps().GetNap("config_1", "pbx_dec0de")
		if err != nil {
			log.Error(err)
		}

		marshal, err := json.Marshal(naps)
		if err != nil {
			return
		}

		pretty, err := prettyJson(marshal)
		if err != nil {
			return
		}

		log.Printf("%s", pretty)*/
}

func prettyJson(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, data, "", "\t")
	if error != nil {
		return "", error
	}
	return string(prettyJSON.Bytes()), nil
}
