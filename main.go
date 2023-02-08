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
	var fNapCreate = flag.Bool("create", false, "specify to use the creation flag")
	var flagUpdate = flag.Bool("update", false, "specify to use the update flag")
	var flagNap = flag.Bool("nap", true, "defines what type of change to make, eg. nap, digitmap, etc")
	var flagPbx = flag.Bool("pbx", true, "defines if a nap is a pbx")

	var fNapName = flag.String("customer", "", "customer name used in nap names and routedefs, etc")

	var fConfigName = flag.String("config", "config_1", "config name to use (default: config_1)")
	var fDigitMap = flag.String("digitmap", "", "digit map to be modified/updated/etc")

	var fNapProxyHost = flag.String("proxyhost", "", "proxy host in ip/host:port format")
	var fPhoneNumbers = flag.String("numbers", "", "phone numbers seperated by "+
		"commands used for various functions (default: empty)")
	var fPortRange = flag.String("portrange", "", "port range for rtp??")
	var fSipTransport = flag.String("siptransport", "", "port range for rtp??")
	var fRdefRouteGroups = flag.String("rdefroutegroups", "", "routegroups to be modified/updated/etc in rdef")
	var fNapcRouteGroups = flag.String("napcroutegroups", "", "routegroups to be modified/updated/etc in napc")
	var fNAPProfile = flag.String("napprofile", "default", "nap profile to use when creating naps, default is default")

	var fMaxTotalCalls = flag.Int("maxtotalcalls", 6, "maximum total of calls allowed on nap (default is 2)")

	// port ranges, interface,sip transport servers,
	flag.Parse()

	fNapCreateBool := *fNapCreate
	fNap := *flagNap
	fUpdate := *flagUpdate
	fPbx := *flagPbx
	fConfigNameStr := *fConfigName
	fDigitMapFile := *fDigitMap

	fRdefRouteGroupsCSV := *fRdefRouteGroups
	fNapcRouteGroupsCSV := *fNapcRouteGroups
	fSipTransportStr := *fSipTransport
	fPortRangeStr := *fPortRange
	fPhoneNumbersStr := *fPhoneNumbers
	fNapProxyHostStr := *fNapProxyHost
	fNapNameStr := *fNapName
	napProfile := *fNAPProfile
	napTotalCalls := *fMaxTotalCalls

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

	napName := fNapNameStr
	if fPbx {
		napName = "pbx_" + napName
	}

	if fUpdate {
		// if the update flag is set, proceed to check if they are wanting to update a nap, pbx, etc
		if fNap && fNapNameStr != "" {
			// todo update max call limit
			nap, err := client.TBNaps().GetNap(fConfigNameStr, napName)
			if err != nil {
				log.Error(err)
				return
			}

			nap.CallRateLimiting.MaximumSimultaneousTotalCalls = napTotalCalls

			newNap := *nap

			err = client.TBNaps().UpdateNap(fConfigNameStr, newNap)
			if err != nil {
				log.Error(err)
				return
			}
		}
	}

	// checks for flags
	if fNapCreateBool && fPbx &&
		fNap &&
		fNapNameStr != "" &&
		fPhoneNumbersStr != "" &&
		fNapProxyHostStr != "" &&
		fConfigNameStr != "" &&
		fPortRangeStr != "" &&
		fSipTransportStr != "" &&
		fDigitMapFile != "" && fPbx && fRdefRouteGroupsCSV != "" && fNapcRouteGroupsCSV != "" {

		// File_DB is default?
		// get digitmap
		// todo have option to provide digitmap routeDefFile name
		digitMap, err := client.TBFileDBs("File_DB").GetDigitMap(fConfigNameStr, fDigitMapFile)
		if err != nil {
			log.Error(err)
			return
		}

		// copy to digitmaporig incase something fails
		var digitMapOrig []sbc.TBDigitMap
		copy(digitMapOrig, digitMap)

		// split numbers fdrom flag and append to digitmap
		phoneNumbers := strings.Split(fPhoneNumbersStr, ",")

		for pN1, _ := range phoneNumbers {
			for pN2, _ := range phoneNumbers {
				if pN1 != pN2 {
					if phoneNumbers[pN1] == phoneNumbers[pN2] {
						log.Error("Duplicate numbers trying to be sent.")
						return
					}
				}
			}
		}

		var duplicateNumber = false
		for _, i := range digitMap {
			for _, i2 := range phoneNumbers {
				if strings.Contains(i.Called, i2) {
					duplicateNumber = true
					log.Error("DigitMap already contains ", i2)
				}
			}
		}

		if duplicateNumber {
			log.Fatalf("Duplicate numbers are being requested to be added to DigitMap")
		}

		for _, i := range phoneNumbers {
			// create new item
			newDigitMapping := sbc.TBDigitMap{
				Called:  i,
				Calling: "",
				//todo friendlyify name to match scheme
				RouteSetName: fNapNameStr,
			}

			// append item
			digitMap = append(digitMap, newDigitMapping)
		}

		// todo create routeset routeDefFile
		routeDefFile := sbc.TBFile{
			Name: fNapNameStr + "_routedef.csv",
		}

		// makes the napname start nwith pbx

		// define a new route for the routedef routeDefFile
		var routeDef []*sbc.TBRouteDef
		route := &sbc.TBRouteDef{
			RouteSetName: fNapNameStr,
			Priority:     10,
			Weight:       50,
			// todo includem routegroups in flags (eg. 55,11,12,32??)
			RouteGroup: fRdefRouteGroupsCSV,
		}
		routeDef = append(routeDef, route)

		// convert the array of routes to a csv
		marsh, err := gocsv.MarshalString(routeDef)
		if err != nil {
			log.Error(err)
			return
		}

		// include correct new line formatting for api
		formatted := strings.ReplaceAll(marsh, "\n", "\r\n")

		// set tbfile contents
		routeDefFile.Content = formatted

		// push data to api
		//todo make it remove files and such if the creation failed
		err = client.TBFileDBs("File_DB").CreateRouteDef(fConfigNameStr, routeDefFile)
		if err != nil {
			log.Error(err)
			return
		}
		//done creating routedef

		// update digit map
		err = client.TBFileDBs("File_DB").UpdateDigitMap(fConfigNameStr, fDigitMapFile, digitMap)
		if err != nil {
			log.Error(err)
			return
		}

		// create nap
		sipHostInfo := strings.Split(fNapProxyHostStr, ":")
		sipProxyPort, err := strconv.ParseInt(sipHostInfo[1], 10, 0)

		nap := sbc.Nap{
			Name: napName,
			CallRateLimiting: sbc.NapCallRateLimiting{
				ProcessingDelayHighThreshold:  "6 seconds",
				ProcessingDelayLowThreshold:   "3 seconds",
				MaximumSimultaneousTotalCalls: napTotalCalls,
			},
			Enabled:        true,
			DefaultProfile: napProfile,
			// Host.pr_voice_vlan
			// todo make a command seperated list for these instead of a single value
			SipTransportServers: []string{fSipTransportStr}, // WAN0_5060
			PortRanges:          []string{fPortRangeStr},    // Host.pr_WAN0
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

		err = client.TBNaps().CreateNap(fConfigNameStr, nap)
		if err != nil {
			// update digit map
			log.Warn("NAP creation failed. Reverting digit map.")
			err = client.TBFileDBs("File_DB").UpdateDigitMap(fConfigNameStr, fDigitMapFile, digitMapOrig)
			if err != nil {
				log.Error(err)
				return
			}

			log.Warn("NAP creation failed. Reverting route def. Deleting file...")
			err = client.TBFileDBs("File_DB").DeleteRouteDef(fConfigNameStr, routeDefFile.Name)
			if err != nil {
				log.Error(err)
				return
			}

			// todo remove created digitmap & routedef if nap creation fails
			log.Error(err)
			return
		}

		// todo update nap columns

		napColumn := sbc.NapColumnValues{
			RoutesetsDefinition: routeDefFile.Name,
			RouteGroups:         fNapcRouteGroupsCSV,
			RoutesetsDigitmap:   fDigitMapFile,
			Weight:              "50",
			//todo flag options to specify these??
			BlackWhiteList: "default_blacklist.csv",
			CalledPreRemap: "/^\\+?1?([2-9]\\d{9})$/\\1", // /^\+?1?([2-9]\d{9})$/\1
			Priority:       "10",
		}

		err = client.TBNapColumnsValues().UpdateNapColumnValues(fConfigNameStr, napName, napColumn)
		if err != nil {
			log.Error(err)
			return
		}

		log.Info("Successfully created NAP: " + nap.Name)
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
