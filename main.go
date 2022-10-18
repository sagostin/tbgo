package main

import (
	log "github.com/sirupsen/logrus"
	"tbgo/sbc"
)

/** INFO
TelcoBridges GoLang Wrapper/API??

TODO:
General Utilities for the Web Requests (POST, GET, PUT)
Models for SBC data types such as NAPs, etc.

FUTURE:
Other TelcoBridges products
*/

func main() {
	cfg := sbc.Cfg{Host: "https://sbc01.dec0de.xyz", Config: "config_1"}

	nap := sbc.Nap{Name: "pbx_dec0de"}
	get, err := nap.Get(cfg)
	if err != nil {
		return
	}
	log.Infof("%s", get)
}
