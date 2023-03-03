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
	var flagDelete = flag.Bool("delete", false, "specify to use the delete flag")
	var flagNap = flag.Bool("nap", false, "defines what type of change to make, eg. nap, digitmap, etc")
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
	var fNAPProxyPoll = flag.Bool("napproxypoll", false, "specify wether to update this section of the cfg")
	var fNAPProxyPollEnable = flag.Bool("napproxypollenable", false, "enable nap proxy polling (default is false, and is to disable)")
	var fNapColumn = flag.Bool("napcolumns", false, "specify whether to update the nap columns of the provided info")

	var fRemoteMethodSIP = flag.String("remote_method_sip", "None", "nat remote sip method")
	var fRemoteMethodRTP = flag.String("remote_method_rtp", "None", "nat remote rtp method")
	var fLocalMethodSIP = flag.String("local_method_sip", "", "nat local sip method")
	var fLocalMethodRTP = flag.String("local_method_rtp", "", "nat local rtp method")

	var fMaxTotalCalls = flag.Int("maxtotalcalls", 6, "maximum total of calls allowed on nap (default is 6)")

	// port ranges, interface,sip transport servers,
	flag.Parse()

	fNapCreateBool := *fNapCreate
	fNap := *flagNap
	fUpdate := *flagUpdate
	fDelete := *flagDelete
	fPbx := *flagPbx
	fConfigNameStr := *fConfigName
	fDigitMapFile := *fDigitMap

	fRdefRouteGroupsCSV := *fRdefRouteGroups
	fNapcRouteGroupsCSV := *fNapcRouteGroups
	fSipTransportStr := *fSipTransport
	fPortRangeStr := *fPortRange
	fPhoneNumbersStr := *fPhoneNumbers
	fNapProxyHostStr := *fNapProxyHost
	fCustomerName := *fNapName
	napProfile := *fNAPProfile
	napTotalCalls := *fMaxTotalCalls
	napProxyPollEnable := *fNAPProxyPollEnable
	napProxyPoll := *fNAPProxyPoll
	napColumn := *fNapColumn

	// nap NAT
	natRemoteSIP := *fRemoteMethodSIP
	natRemoteRTP := *fRemoteMethodRTP
	natLocalSIP := *fLocalMethodSIP
	natLocalRTP := *fLocalMethodRTP

	phoneNumbers := strings.Split(fPhoneNumbersStr, ",")

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

	napName := fCustomerName
	if fPbx {
		napName = "pbx_" + napName
	}

	/*client.TBSystem().GetSystem("system_1")*/

	// delete
	if fDelete {
		// delete single numbers or batch
		// delete full naps (remove numbers, remove routedefs, etc.)

		if len(phoneNumbers) > 0 && fDigitMapFile != "" {
			dMap, err := client.TBFileDBs("File_DB").GetDigitMap(fConfigNameStr, fDigitMapFile)
			if err != nil {
				log.Error(err)
				return
			}

			var newDmap []sbc.TBDigitMap
			// loop through digits in the digitmap
			for _, d := range dMap {
				// loop through numbers that are to be removed
				match := false
				for _, number := range phoneNumbers {
					// check if customer name and called number match one of the numbers to be "deleted" from digitmap
					if d.Called == number && d.RouteSetName == fCustomerName {
						// set match as true for further processing
						match = true
						continue
					}
				}

				// if there wasn't a match on that number, add it to the new temp digitmap
				if !match {
					newDmap = append(newDmap, d)
				} else {
					log.Info("Found match for " + d.Called + ". Not adding to temporary digitmap. Deleting number...")
				}
				continue
			}

			err = client.TBFileDBs("File_DB").UpdateDigitMap(fConfigNameStr, fDigitMapFile, newDmap)
			if err != nil {
				log.Error(err)
				return
			}
		}
	}

	// process updating for config

	if fUpdate && fConfigNameStr != "" {

		// if create flag is true and phone numbers is more than 0
		// add numbers to the digitmap for the specified nap, but validate that it already exists
		// we don't want people accidentally typing the wrong shit
		if fNapCreateBool && len(phoneNumbers) > 0 {
			// todo validate the nap already exists
			digitMap, err := client.TBFileDBs("File_DB").GetDigitMap(fConfigNameStr, fDigitMapFile)
			if err != nil {
				log.Error(err)
				return
			}

			updatedDigitmap, err := sbc.Append2Digitmap(fCustomerName, phoneNumbers, digitMap)
			if err != nil {
				log.Error(err)
				return
			}

			err = client.TBFileDBs("File_DB").UpdateDigitMap(fConfigNameStr, fDigitMapFile, updatedDigitmap)
			if err != nil {
				log.Error(err)
				return
			}
		}

		// process the max call limit update request
		if fNap && fCustomerName != "" {
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
			return
		}

		// update route group values in nap columns
		if napColumn && fCustomerName != "" {
			v, err := client.TBNaps().GetColumnValues(fConfigNameStr, napName)
			if err != nil {
				log.Error(err)
				return
			}

			v.RouteGroups = fNapcRouteGroupsCSV
			valuesCopy := *v

			err = client.TBNapColumnsValues().UpdateNapColumnValues(fConfigNameStr, napName, valuesCopy)
			if err != nil {
				log.Error(err)
				return
			}

			log.Info(valuesCopy)
			return
		}

		// mass enable / disable proxy polling for all naps
		if napProxyPoll {
			getNaps, err := client.TBNaps().GetNames(fConfigNameStr)

			if err != nil {
				log.Error(err)
				return
			}

			for _, nap := range getNaps {
				if !strings.Contains(nap, "pbx_") {
					log.Error("nap does not start with pbx_")
					continue
				}

				getNap, err := client.TBNaps().GetNap(fConfigNameStr, nap)
				if err != nil {
					log.Error(err)
					return
				}
				getNap.SipCfg.PollRemoteProxy = napProxyPollEnable
				newNap := *getNap

				err = client.TBNaps().UpdateNap(fConfigNameStr, newNap)
				if err != nil {
					log.Error(err)
					return
				}
				log.Info("Updated NAP for " + nap)
			}

			return
		}
	}

	// checks for flags
	if fNapCreateBool && fPbx &&
		fNap &&
		fCustomerName != "" &&
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

		updatedDigitmap, err := sbc.Append2Digitmap(fCustomerName, phoneNumbers, digitMap)
		if err != nil {
			log.Error(err)
			return
		}

		// todo append items to digitmap function :)

		// todo create routeset routeDefFile
		routeDefFile := sbc.TBFile{
			Name: fCustomerName + "_routedef.csv",
		}

		// makes the napname start nwith pbx

		// define a new route for the routedef routeDefFile
		var routeDef []*sbc.TBRouteDef
		route := &sbc.TBRouteDef{
			RouteSetName: fCustomerName,
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
		err = client.TBFileDBs("File_DB").UpdateDigitMap(fConfigNameStr, fDigitMapFile, updatedDigitmap)
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
				PollRemoteProxy: napProxyPollEnable,
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
					RemoteMethodSip: natRemoteSIP,
					RemoteMethodRtp: natRemoteRTP,
					LocalMethodSip:  natLocalSIP,
					LocalMethodRtp:  natLocalRTP,
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
			err = client.TBFileDBs("File_DB").UpdateDigitMap(fConfigNameStr, fDigitMapFile, digitMap)
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
