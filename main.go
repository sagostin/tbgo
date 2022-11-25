package main

import (
	"bytes"
	"encoding/json"
	"flag"
	log "github.com/sirupsen/logrus"
	"tbgo/sbc"
)

/** INFO */

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

	/*err := client.Request("GET", "/configurations/config_1?recursive=yes", nil, nil)
	if err != nil {
		log.Error(err)
	}*/

	// File_DB is default?
	// get digitmap
	/*def, err := client.TBFileDBs("File_DB").GetDigitMap("config_1", "digitmap_new.csv")
	if err != nil {
		log.Error(err)
	}*/

	var digitMap []*sbc.TBDigitMap

	// create new item
	newDigitMapping := &sbc.TBDigitMap{
		Called:       "2508591501",
		Calling:      "",
		RouteSetName: "dec0de",
	}

	// append item
	digitMap = append(digitMap, newDigitMapping)

	// update digit map
	err := client.TBFileDBs("File_DB").UpdateDigitMap("config_1", "digitmap_new.csv", digitMap)
	if err != nil {
		log.Error(err)
	}

	// get digitmap
	getAgain, err := client.TBFileDBs("File_DB").GetDigitMap("config_1", "digitmap_new.csv")
	if err != nil {
		log.Error(err)
	}

	marshal1, err := json.Marshal(&digitMap)
	if err != nil {
		log.Error(err)
	}
	marshal2, err := json.Marshal(&getAgain)
	if err != nil {
		log.Error(err)
	}

	pretty1, err := prettyJson(marshal1)
	if err != nil {
		log.Error(err)
	}

	pretty2, err := prettyJson(marshal2)
	if err != nil {
		log.Error(err)
	}

	log.Printf("\n" + pretty1)
	log.Printf("\n" + pretty2)

	/*names, err := client.TBNaps().GetNap("config_1", "pbx_dec0de")
	if err != nil {
		return
	}

	marshal, err := json.Marshal(names)
	if err != nil {
		return
	}*/

	/*log.Infof("%s", marshal)*/

	/*	log.Warnf("%s", "Creating NAP")

		nap := sbc.Nap{
			Name: "pbx_tops1",
			CallRateLimiting: sbc.NapCallRateLimiting{
				ProcessingDelayHighThreshold: "6 seconds",
				ProcessingDelayLowThreshold:  "3 seconds",
			},
			Enabled:             true,
			DefaultProfile:      "default",
			PortRanges:          []string{"Host.pr_voice_vlan"},
			SipTransportServers: []string{"voice_net"},
			SipCfg: sbc.NapSipCfg{
				PollRemoteProxy: true,
				SipiParameters: sbc.NapSipiParams{
					IsupProtocolVariant: "ITU",
					ContentType:         "itu-t",
					CallProgressMethod:  "183 Call Progress",
				},
				AdvancedParameters: sbc.NapAdvancedParams{
					MapAnyResponseToAvailableStatus: true,
					ResponseTimeout:                 "12 seconds",
					PrivacyType:                     "P-Asserted-Identity",
					ProxyPollingMaxForwards:         1,
				},
				ProxyPortType: "UDP",
				SipUseProxy:   true,
				ProxyPort:     5060,
				FilteringParameters: sbc.NapFilterParams{
					FilterByLocalPort:    true,
					FilterByProxyPort:    true,
					FilterByProxyAddress: true,
				},
				ProxyPollingInterval: "1 minute",
				ProxyAddress:         "10.0.40.12",
				NetworkAddressTranslation: sbc.NapNatParams{
					RemoteMethodSip: "None",
					RemoteMethodRtp: "None",
				},
			},
			CongestionThreshold: sbc.NapCongestionThreshold{
				PeriodDuration:   "1 minute",
				NbCallsPerPeriod: 1,
			},
		}

		err = client.TBNaps().CreateNap("config_1", nap)
		if err != nil {
			log.Fatal(err)
		}
	*/
	/*

		{
		  "name": "pbx_dec0de",
		  "call_rate_limiting": {
		    "processing_delay_high_threshold": "6 seconds",
		    "processing_delay_low_threshold": "3 seconds"
		  },
		  "enabled": true,
		  "default_profile": "default",
		  "port_ranges": [
		    "Host.pr_voice_vlan"
		  ],
		  "sip_transport_servers": [
		    "voice_net"
		  ],
		  "sip_cfg": {
		    "poll_remote_proxy": true,
		    "registration_parameters": {},
		    "sipi_parameters": {
		      "isup_protocol_variant": "ITU",
		      "content_type": "itu-t",
		      "call_progress_method": "183 Call Progress"
		    },
		    "advanced_parameters": {
		      "map_any_response_to_available_status": true,
		      "response_timeout": "12 seconds",
		      "privacy_type": "P-Asserted-Identity",
		      "proxy_polling_max_forwards": 1
		    },
		    "proxy_port_type": "UDP",
		    "sip_use_proxy": true,
		    "proxy_port": 5060,
		    "filtering_parameters": {
		      "filter_by_proxy_port": true,
		      "filter_by_local_port": true,
		      "filter_by_proxy_address": true
		    },
		    "proxy_polling_interval": "1 minute",
		    "proxy_address": "10.0.40.11",
		    "authentication_parameters": {},
		    "network_address_translation": {
		      "remote_method_sip": "None",
		      "remote_method_rtp": "None"
		    }
		  },
		  "congestion_threshold": {
		    "period_duration": "1 minute",
		    "nb_calls_per_period": 1
		  }
		}

	*/

	/*// get configs
	getConfigs, err := client.TBConfigs().GetNames()
	if err != nil {
		log.Error(err)
	}
	log.Printf("%s",getConfigs)

	// get list of naps
	client.TBNaps().GetNames()*/

	// todo implement ability to download the routedef and such.
	/*names, err := client.TBFileDBs("File_DB").GetRouteDef("config_1", "dec0de_routedef.csv")
	if err != nil {
		return
	}
	log.Printf("%s", names)*/
}

func prettyJson(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, data, "", "\t")
	if error != nil {
		return "", error
	}
	return string(prettyJSON.Bytes()), nil
}
