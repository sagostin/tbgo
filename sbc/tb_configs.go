package sbc

import (
	"encoding/json"
	"strings"
)

/*
potential paths

"date_time_services": {},
        "dns_groups": {},
        "notes": "",
        "dns_services": {},
        "hosts": {},
        "host_ip_interfaces": {},
        "certificates": {},
        "profiles": {},
        "routing_scripts": {},
        "gateways": {},
        "tls_profiles": {},
        "telephony_services": {},
        "radius_clients": {},
        "http_services": {},
        "validation_status": "Successful",
        "file_dbs": {},
        "sip_stacks": {},
        "firewalls": {},
        "host_virtual_ports": {},
        "np1_groups": {},
        "radius_servers": {},
        "user_tone_definition_profiles": {},
        "nap_columns": {},
        "privilege_level": 0,
        "applications": {},
        "sip_emergency_cfgs": {},
        "nat_traversals": {},
        "snmp_services": {},
        "nap_columns_values": {},
        "host_port_ranges": {},
        "sip_registration_domains": {},
        "ssh_services": {},
        "sig_memory_schemes": [],
        "unapplied_changes": false,
        "naps": {},
        "icmp_services": {},
        "hardware_units": {},
        "sip_header_manipulation_cfgs": {},
        "routes": {},
        "route_columns": {},
        "internal_nats": {},
        "lawful_interceptions": {},
*/

/*
TODO
- Ability to copy and load new template based on current newest configuration
- After applying the configuration, activate it, but don't delete the old configuration
- After x amount of old configs, prune?
-
*/

// TBConfigs client
type TBConfigs struct {
	Client
}

// TBConfigs constructor (from Client)
func (c Client) TBConfigs() TBConfigs {
	return TBConfigs{
		Client: c,
	}
}

func (c TBConfigs) GetNames() ([]string, error) {
	d := make(map[string]json.RawMessage)

	err := c.Client.Request("GET", "/configurations/", nil, &d)
	if err != nil {
		return nil, err
	}

	// thank u stackoverflow: https://stackoverflow.com/questions/17452722/how-to-get-the-key-value-from-a-json-string-in-go
	// a string slice to hold the keys
	k := make([]string, len(d)-1)
	// iteration counter
	i := 0
	// copy c's keys into k
	for s := range d {
		if strings.Contains(s, "meta") {
			continue
		}
		k[i] = s
		i++
	}

	return k, nil
}

type Configuration struct {
	WebProfile struct {
		HardwareUnit      bool `json:"hardware_unit"`
		SigtranIua        bool `json:"sigtran_iua"`
		Sip               bool `json:"sip"`
		Isdn              bool `json:"isdn"`
		Isup              bool `json:"isup"`
		Monitoring        bool `json:"monitoring"`
		H248              bool `json:"h248"`
		StatisticsHistory bool `json:"statistics_history"`
		Cas               bool `json:"cas"`
		Firewall          bool `json:"firewall"`
		Gateway           bool `json:"gateway"`
		SigtranM2Pa       bool `json:"sigtran_m2pa"`
		SccpTcap          bool `json:"sccp_tcap"`
		TmsIp             bool `json:"tms_ip"`
		SigtranM2Ua       bool `json:"sigtran_m2ua"`
		SigtranM3Ua       bool `json:"sigtran_m3ua"`
	} `json:"web_profile"`
	Name string `json:"name"`
}

func (c TBConfigs) GetConfig(name string) (*Configuration, error) {
	var config *Configuration

	err := c.Client.Request("GET", "/configurations/"+name, nil, &config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
