package sbc

type T2 struct {
	Applications struct {
	} `json:"applications"`
	Naps struct {
	} `json:"naps"`
	Notes        string `json:"notes"`
	RouteColumns struct {
	} `json:"route_columns"`
	HostPortRanges struct {
	} `json:"host_port_ranges"`
	Np1Groups struct {
	} `json:"np1_groups"`
	IcmpServices struct {
	} `json:"icmp_services"`
	Name    string `json:"name"`
	FileDbs struct {
	} `json:"file_dbs"`
	UserToneDefinitionProfiles struct {
	} `json:"user_tone_definition_profiles"`
	NapColumnsValues struct {
	} `json:"nap_columns_values"`
	HttpServices struct {
	} `json:"http_services"`
	Routes struct {
	} `json:"routes"`
	RadiusServers struct {
	} `json:"radius_servers"`
	DnsServices struct {
	} `json:"dns_services"`
	UnappliedChanges bool `json:"unapplied_changes"`
	RoutingScripts   struct {
	} `json:"routing_scripts"`
	Certificates struct {
	} `json:"certificates"`
	PrivilegeLevel      int `json:"privilege_level"`
	LawfulInterceptions struct {
	} `json:"lawful_interceptions"`
	Hosts struct {
	} `json:"hosts"`
	NapColumns struct {
	} `json:"nap_columns"`
	HostVirtualPorts struct {
	} `json:"host_virtual_ports"`
	SipRegistrationDomains struct {
	} `json:"sip_registration_domains"`
	RadiusClients struct {
	} `json:"radius_clients"`
	SnmpServices struct {
	} `json:"snmp_services"`
	SigMemorySchemes []interface{} `json:"sig_memory_schemes"`
	InternalNats     struct {
	} `json:"internal_nats"`
	Gateways struct {
	} `json:"gateways"`
	TlsProfiles struct {
	} `json:"tls_profiles"`
	DnsGroups struct {
	} `json:"dns_groups"`
	TelephonyServices struct {
	} `json:"telephony_services"`
	Meta struct {
		Version string `json:"version"`
		SrcPath string `json:"src_path"`
	} `json:"***meta***"`
	HardwareUnits struct {
	} `json:"hardware_units"`
	SipEmergencyCfgs struct {
	} `json:"sip_emergency_cfgs"`
	Firewalls struct {
	} `json:"firewalls"`
	DateTimeServices struct {
	} `json:"date_time_services"`
	SshServices struct {
	} `json:"ssh_services"`
	ValidationStatus string `json:"validation_status"`
	Profiles         struct {
	} `json:"profiles"`
	HostIpInterfaces struct {
	} `json:"host_ip_interfaces"`
	SipStacks struct {
	} `json:"sip_stacks"`
	SipHeaderManipulationCfgs struct {
	} `json:"sip_header_manipulation_cfgs"`
	NatTraversals struct {
	} `json:"nat_traversals"`
	WebProfile struct {
		Cas         bool `json:"cas"`
		TmsIp       bool `json:"tms_ip"`
		Isup        bool `json:"isup"`
		Sip         bool `json:"sip"`
		Firewall    bool `json:"firewall"`
		SigtranM2Pa bool `json:"sigtran_m2pa"`
		SccpTcap    bool `json:"sccp_tcap"`
		Meta        struct {
			ValidUrl bool `json:"valid_url"`
		} `json:"***meta***"`
		H248              bool `json:"h248"`
		HardwareUnit      bool `json:"hardware_unit"`
		SigtranM2Ua       bool `json:"sigtran_m2ua"`
		SigtranM3Ua       bool `json:"sigtran_m3ua"`
		SigtranIua        bool `json:"sigtran_iua"`
		Gateway           bool `json:"gateway"`
		Isdn              bool `json:"isdn"`
		Monitoring        bool `json:"monitoring"`
		StatisticsHistory bool `json:"statistics_history"`
	} `json:"web_profile"`
}
