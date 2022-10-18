package sbc

import (
	"encoding/json"
)

type Nap struct {
	Name             string `json:"name"`
	CallRateLimiting struct {
		ProcessingDelayHighThreshold     string `json:"processing_delay_high_threshold"`
		MaximumCallsPerSecond            int    `json:"maximum_calls_per_second"`
		MaximumCallBurst                 int    `json:"maximum_call_burst"`
		MaximumSimultaneousOutgoingCalls int    `json:"maximum_simultaneous_outgoing_calls"`
		MaximumSimultaneousIncomingCalls int    `json:"maximum_simultaneous_incoming_calls"`
		MaximumOutgoingCallsPerSecond    int    `json:"maximum_outgoing_calls_per_second"`
		MaximumIncomingCallsPerSecond    int    `json:"maximum_incoming_calls_per_second"`
		Meta                             struct {
			ValidUrl bool `json:"valid_url"`
		} `json:"***meta***"`
		MaximumSimultaneousTotalCalls int    `json:"maximum_simultaneous_total_calls"`
		ProcessingDelayLowThreshold   string `json:"processing_delay_low_threshold"`
	} `json:"call_rate_limiting"`
	Enabled             bool     `json:"enabled"`
	DefaultProfile      string   `json:"default_profile"`
	PortRanges          []string `json:"port_ranges"`
	SipTransportServers []string `json:"sip_transport_servers"`
	Meta                struct {
		Version string `json:"version"`
		SrcPath string `json:"src_path"`
	} `json:"***meta***"`
	SipCfg struct {
		PollRemoteProxy           bool `json:"poll_remote_proxy"`
		AcceptOnlyAuthorizedUsers bool `json:"accept_only_authorized_users"`
		RegistrationParameters    struct {
			AddressToRegister string `json:"address_to_register"`
			RegisterToProxy   bool   `json:"register_to_proxy"`
			Meta              struct {
				ValidUrl bool `json:"valid_url"`
			} `json:"***meta***"`
		} `json:"registration_parameters"`
		SipiParameters struct {
			IsupProtocolVariant    string `json:"isup_protocol_variant"`
			AppendFToOutgoingCalls bool   `json:"append_f_to_outgoing_calls"`
			Meta                   struct {
				ValidUrl bool `json:"valid_url"`
			} `json:"***meta***"`
			Enable             bool   `json:"enable"`
			ContentType        string `json:"content_type"`
			CallProgressMethod string `json:"call_progress_method"`
		} `json:"sipi_parameters"`
		AdvancedParameters struct {
			MapAnyResponseToAvailableStatus bool   `json:"map_any_response_to_available_status"`
			ResponseTimeout                 string `json:"response_timeout"`
			PrivacyType                     string `json:"privacy_type"`
			ProxyPollingMaxForwards         int    `json:"proxy_polling_max_forwards"`
			Meta                            struct {
				ValidUrl bool `json:"valid_url"`
			} `json:"***meta***"`
			TriggersCallProgress bool `json:"183_triggers_call_progress"`
		} `json:"advanced_parameters"`
		ProxyPortType string        `json:"proxy_port_type"`
		NapSipAcls    []interface{} `json:"nap_sip_acls"`
		SipUseProxy   bool          `json:"sip_use_proxy"`
		Meta          struct {
			ValidUrl bool `json:"valid_url"`
		} `json:"***meta***"`
		ProxyPort           int `json:"proxy_port"`
		FilteringParameters struct {
			FilterByProxyPort bool `json:"filter_by_proxy_port"`
			Meta              struct {
				ValidUrl bool `json:"valid_url"`
			} `json:"***meta***"`
			FilterByLocalPort    bool `json:"filter_by_local_port"`
			FilterByProxyAddress bool `json:"filter_by_proxy_address"`
		} `json:"filtering_parameters"`
		ProxyPollingInterval     string `json:"proxy_polling_interval"`
		ProxyAddress             string `json:"proxy_address"`
		AuthenticationParameters struct {
			User           string `json:"user"`
			Realm          string `json:"realm"`
			ReuseChallenge bool   `json:"reuse_challenge"`
			Meta           struct {
				ValidUrl bool `json:"valid_url"`
			} `json:"***meta***"`
			Password    string `json:"password"`
			IgnoreRealm bool   `json:"ignore_realm"`
		} `json:"authentication_parameters"`
		NetworkAddressTranslation struct {
			RemoteMethodSip string `json:"remote_method_sip"`
			LocalMethodRtp  string `json:"local_method_rtp"`
			RemoteMethodRtp string `json:"remote_method_rtp"`
			Meta            struct {
				ValidUrl bool `json:"valid_url"`
			} `json:"***meta***"`
			LocalMethodSip string `json:"local_method_sip"`
		} `json:"network_address_translation"`
	} `json:"sip_cfg"`
	CongestionThreshold struct {
		PeriodDuration string `json:"period_duration"`
		Meta           struct {
			ValidUrl bool `json:"valid_url"`
		} `json:"***meta***"`
		NbCallsPerPeriod int `json:"nb_calls_per_period"`
	} `json:"congestion_threshold"`
}

func (n *Nap) Get(sbc Cfg) (*Nap, error) {
	// TODO include authentication information
	http, err := GetHttp(sbc.Host, "/configurations/"+sbc.Config+"/naps/"+n.Name)
	if err != nil {
		return nil, err
	}
	var nap *Nap
	err = json.Unmarshal(http, &nap)
	if err != nil {
		return nil, err
	}
	return nap, nil
	// TODO unmarshal body to api response modes
}
