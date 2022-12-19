package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"github.com/gocarina/gocsv"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
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

	// flags to do different things

	// nap create variables
	var fNapCreate = flag.Bool("napcreate", false, "create a nap w/ routes")
	fNapCreateBool := *fNapCreate

	var flagPbx = flag.Bool("pbx", true, "defines if a nap is a pbx")
	fPbx := *flagPbx

	var fNapName = flag.String("customer", "", "customer name used in nap names and routedefs, etc")
	fNapNameStr := *fNapName

	var fNapProxyHost = flag.String("napproxyhost", "", "proxy host in ip/host:port format")
	fNapProxyHostStr := *fNapProxyHost

	var fPhoneNumbers = flag.String("numbers", "", "phone numbers seperated by "+
		"commands used for various functions (default: empty)")
	fPhoneNumbersStr := *fPhoneNumbers

	var fConfigName = flag.String("config", "config_1", "config name to use (default: config_1)")
	fConfigNameStr := *fConfigName

	var fPortRange = flag.String("portrange", "", "port range for rtp??")
	fPortRangeStr := *fPortRange

	var fSipTransport = flag.String("portrange", "", "port range for rtp??")
	fSipTransportStr := *fSipTransport

	var fDigitMap = flag.String("digitmap", "", "digit map to be modified/updated/etc")
	fDigitMapFile := *fDigitMap

	// port ranges, interface,sip transport servers,

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

	// checks for flags
	if fNapCreateBool && fNapNameStr != "" &&
		fPhoneNumbersStr != "" &&
		fNapProxyHostStr != "" &&
		fConfigNameStr != "" &&
		fPortRangeStr != "" &&
		fSipTransportStr != "" &&
		fDigitMapFile != "" && fPbx {

		// todo create routeset routeDefFile
		routeDefFile := sbc.TBFile{
			Name: fNapNameStr + "_routedef.csv",
		}

		// makes the napname start nwith pbx

		napName := fNapNameStr
		if fPbx {
			napName = "pbx_" + napName
		}

		// define a new route for the routedef routeDefFile
		var routeDef []*sbc.TBRouteDef
		route := &sbc.TBRouteDef{
			RouteSetName: fNapNameStr,
			Priority:     10,
			Weight:       50,
			// todo includem routegroups in flags (eg. 55,11,12,32??)
			RouteGroup: "",
		}
		routeDef = append(routeDef, route)

		// convert the array of routes to a csv
		marsh, err := gocsv.MarshalString(routeDef)
		if err != nil {
			return
		}

		// include correct new line formatting for api
		formatted := strings.ReplaceAll(marsh, "\n", "\r\n")

		// set tbfile contents
		routeDefFile.Content = formatted

		// push data to api
		err = client.TBFileDBs("File_DB").CreateRouteDef(fConfigNameStr, routeDefFile.Name, routeDefFile)
		if err != nil {
			return
		}
		//done creating routedef

		// File_DB is default?
		// get digitmap
		// todo have option to provide digitmap routeDefFile name
		digitMap, err := client.TBFileDBs("File_DB").GetDigitMap(fConfigNameStr, "digitmap_new.csv")
		if err != nil {
			log.Error(err)
		}

		// split numbers fdrom flag and append to digitmap
		phoneNumbers := strings.Split(fPhoneNumbersStr, ",")
		for _, i := range phoneNumbers {
			// create new item
			newDigitMapping := &sbc.TBDigitMap{
				Called:  i,
				Calling: "",
				//todo friendlyify name to match scheme
				RouteSetName: fNapNameStr,
			}

			// append item
			digitMap = append(digitMap, newDigitMapping)
		}

		// update digit map
		err = client.TBFileDBs("File_DB").UpdateDigitMap("config_1", "digitmap_new.csv", digitMap)
		if err != nil {
			log.Error(err)
		}

		// create nap
		sipHostInfo := strings.Split(fNapProxyHostStr, ":")
		sipProxyPort, err := strconv.ParseInt(sipHostInfo[1], 10, 0)

		nap := sbc.Nap{
			Name: napName,
			CallRateLimiting: sbc.NapCallRateLimiting{
				ProcessingDelayHighThreshold: "6 seconds",
				ProcessingDelayLowThreshold:  "3 seconds",
			},
			Enabled:        true,
			DefaultProfile: "default",
			// Host.pr_voice_vlan
			// todo make a command seperated list for these instead of a single value
			PortRanges:          []string{fPortRangeStr},
			SipTransportServers: []string{fSipTransportStr},
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
				ProxyPort:     int(sipProxyPort),
				FilteringParameters: sbc.NapFilterParams{
					FilterByLocalPort:    true,
					FilterByProxyPort:    true,
					FilterByProxyAddress: true,
				},
				ProxyPollingInterval: "1 minute",
				ProxyAddress:         sipHostInfo[0],
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
			return
		}

		// todo update nap columns

		// todo

		/*
			TODO
			1. Create Routedef
			2. Modify Digitmap to route numbers to Routedef
			3. Update the Nap COlumns
			4. Generate routes

		*/

		return
	}
}

func prettyJson(data []byte) (string, error) {
	var prettyJSON bytes.Buffer
	error := json.Indent(&prettyJSON, data, "", "\t")
	if error != nil {
		return "", error
	}
	return string(prettyJSON.Bytes()), nil
}
