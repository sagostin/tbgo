package sbc

type CallRateLimiting struct {
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
}

type Nap struct {
	Name                string           `json:"name"`
	CallRateLimiting    CallRateLimiting `json:"call_rate_limiting"`
	Enabled             bool             `json:"enabled"`
	DefaultProfile      string           `json:"default_profile"`
	PortRanges          []string         `json:"port_ranges"`
	SipTransportServers []string         `json:"sip_transport_servers"`
	SipCfg              struct {
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
		PeriodDuration   string `json:"period_duration"`
		NbCallsPerPeriod int    `json:"nb_calls_per_period"`
	} `json:"congestion_threshold"`
}

/*func (n *Nap) Get(sbc Cfg) (*Nap, error) {
	http, err := GetHttp(sbc.Host, "/configurations/"+sbc.Config+"/naps/"+n.Name, sbc.User, sbc.Password)
	if err != nil {
		return nil, err
	}
	log.Info(string(http))

	var nap *Nap
	err = json.Unmarshal(http, &nap)
	if err != nil {
		return nil, err
	}
	return nap, nil
}*/

type NapStatus struct {
	AvailabilityDetectionStruct struct {
		PollRemoteProxy string `json:"poll_remote_proxy"`
		IsAvailable     string `json:"is_available"`
	} `json:"availability_detection_struct"`
	PortRangeSharedUsagePercent    int    `json:"port_range_shared_usage_percent"`
	AvailableCnt                   int    `json:"available_cnt"`
	InstIncomingCallCntTerminating int    `json:"inst_incoming_call_cnt_terminating"`
	InstIncomingCallCntAnswered    int    `json:"inst_incoming_call_cnt_answered"`
	SignalingType                  string `json:"signaling_type"`
	TotalIncomingFilePlaybacks     int    `json:"total_incoming_file_playbacks"`
	InstOutgoingCallCnt            int    `json:"inst_outgoing_call_cnt"`
	InstIncomingEmergencyCallCnt   int    `json:"inst_incoming_emergency_call_cnt"`
	ResetAsrStats                  string `json:"reset_asr_stats"`
	InstOutgoingCallRate           int    `json:"inst_outgoing_call_rate"`
	InstIncomingCallRateAnswered   int    `json:"inst_incoming_call_rate_answered"`
	InstIncomingCallRateAccepted   int    `json:"inst_incoming_call_rate_accepted"`
	FirewallBlockedCnt             int    `json:"firewall_blocked_cnt"`
	ResetStats                     string `json:"reset_stats"`
	ResetNapDropStats              string `json:"reset_nap_drop_stats"`
	AsrStatsIncomingStruct         struct {
		Last24HCallCnt        int `json:"last_24h_call_cnt"`
		Last24HAsrPercent     int `json:"last_24h_asr_percent"`
		TotalCallCnt          int `json:"total_call_cnt"`
		GlobalAsrPercent      int `json:"global_asr_percent"`
		LastHourCallCnt       int `json:"last_hour_call_cnt"`
		CurrentHourCallCnt    int `json:"current_hour_call_cnt"`
		TotalAnsweredCallCnt  int `json:"total_answered_call_cnt"`
		TotalAcceptedCallCnt  int `json:"total_accepted_call_cnt"`
		LastHourAsrPercent    int `json:"last_hour_asr_percent"`
		CurrentHourAsrPercent int `json:"current_hour_asr_percent"`
	} `json:"asr_stats_incoming_struct"`
	UsagePercent                         int `json:"usage_percent"`
	TotalIncomingInterceptions           int `json:"total_incoming_interceptions"`
	InstIncomingFilePlaybacks            int `json:"inst_incoming_file_playbacks"`
	InstOutgoingCallCntAnswered          int `json:"inst_outgoing_call_cnt_answered"`
	InstIncomingEmergencyCallRateHighest int `json:"inst_incoming_emergency_call_rate_highest"`
	UniqueId                             int `json:"unique_id"`
	SystemDropStats                      struct {
	} `json:"system_drop_stats"`
	LocalDropStats struct {
	} `json:"local_drop_stats"`
	MosStruct struct {
		CurrentHourEgress  float64 `json:"current_hour_egress"`
		LastHourEgress     float64 `json:"last_hour_egress"`
		CurrentHourIngress float64 `json:"current_hour_ingress"`
		LastHourIngress    float64 `json:"last_hour_ingress"`
		Last24HIngress     float64 `json:"last_24h_ingress"`
		Last24HEgress      float64 `json:"last_24h_egress"`
	} `json:"mos_struct"`
	SipSharedUsagePercent               int `json:"sip_shared_usage_percent"`
	InstIncomingCallRateAnsweredHighest int `json:"inst_incoming_call_rate_answered_highest"`
	InstIncomingCallCnt                 int `json:"inst_incoming_call_cnt"`
	TotalOutgoingFileRecordings         int `json:"total_outgoing_file_recordings"`
	InstOutgoingCallRateAnsweredHighest int `json:"inst_outgoing_call_rate_answered_highest"`
	InstIncomingCallRate                int `json:"inst_incoming_call_rate"`
	InstIncomingCallCntInProgress       int `json:"inst_incoming_call_cnt_in_progress"`
	RemoteDropStats                     struct {
	} `json:"remote_drop_stats"`
	AvailabilityPercent              int  `json:"availability_percent"`
	InstIncomingFileRecordings       int  `json:"inst_incoming_file_recordings"`
	InstOutgoingCallRateAccepted     int  `json:"inst_outgoing_call_rate_accepted"`
	FirewallBlocked                  bool `json:"firewall_blocked"`
	CallCongestionPeriodDroppedCalls int  `json:"call_congestion_period_dropped_calls"`
	Meta                             struct {
		Version string `json:"version"`
		SrcPath string `json:"src_path"`
	} `json:"***meta***"`
	RegistrationStruct struct {
		Registered      string `json:"registered"`
		RegisterToProxy string `json:"register_to_proxy"`
	} `json:"registration_struct"`
	NetworkQualityStruct struct {
		CurrentHourEgress  int `json:"current_hour_egress"`
		LastHourEgress     int `json:"last_hour_egress"`
		CurrentHourIngress int `json:"current_hour_ingress"`
		LastHourIngress    int `json:"last_hour_ingress"`
		Last24HIngress     int `json:"last_24h_ingress"`
		Last24HEgress      int `json:"last_24h_egress"`
	} `json:"network_quality_struct"`
	AsrStatsOutgoingStruct struct {
		Last24HCallCnt        int `json:"last_24h_call_cnt"`
		Last24HAsrPercent     int `json:"last_24h_asr_percent"`
		TotalCallCnt          int `json:"total_call_cnt"`
		GlobalAsrPercent      int `json:"global_asr_percent"`
		LastHourCallCnt       int `json:"last_hour_call_cnt"`
		CurrentHourCallCnt    int `json:"current_hour_call_cnt"`
		TotalAnsweredCallCnt  int `json:"total_answered_call_cnt"`
		TotalAcceptedCallCnt  int `json:"total_accepted_call_cnt"`
		LastHourAsrPercent    int `json:"last_hour_asr_percent"`
		CurrentHourAsrPercent int `json:"current_hour_asr_percent"`
	} `json:"asr_stats_outgoing_struct"`
	InstOutgoingCallRateHighest          int  `json:"inst_outgoing_call_rate_highest"`
	InstIncomingEmergencyCallRate        int  `json:"inst_incoming_emergency_call_rate"`
	LowDelayRelaySharedUsagePercent      int  `json:"low_delay_relay_shared_usage_percent"`
	TotalOutgoingInterceptions           int  `json:"total_outgoing_interceptions"`
	InstOutgoingFilePlaybacks            int  `json:"inst_outgoing_file_playbacks"`
	InstIncomingInterceptions            int  `json:"inst_incoming_interceptions"`
	CallCongestion                       bool `json:"call_congestion"`
	MipsSharedUsagePercent               int  `json:"mips_shared_usage_percent"`
	SharedUsagePercent                   int  `json:"shared_usage_percent"`
	UnavailableCnt                       int  `json:"unavailable_cnt"`
	InstOutgoingFileRecordings           int  `json:"inst_outgoing_file_recordings"`
	InstOutgoingCallRateAnswered         int  `json:"inst_outgoing_call_rate_answered"`
	InstOutgoingCallCntTerminating       int  `json:"inst_outgoing_call_cnt_terminating"`
	InstIncomingEmergencyCallCntAnswered int  `json:"inst_incoming_emergency_call_cnt_answered"`
	RtpStatisticsStruct                  struct {
		FromNetNbOtherErrors        int `json:"from_net_nb_other_errors"`
		FromNetNbLostPackets        int `json:"from_net_nb_lost_packets"`
		T38NbPagesFromTdm           int `json:"t38_nb_pages_from_tdm"`
		FromNetNbBadProtocolHeaders int `json:"from_net_nb_bad_protocol_headers"`
		FromNetNbPackets            int `json:"from_net_nb_packets"`
		ToNetNbPackets              int `json:"to_net_nb_packets"`
		T38NbPagesToTdm             int `json:"t38_nb_pages_to_tdm"`
		ToNetNbArpFailures          int `json:"to_net_nb_arp_failures"`
		FromNetNbBufferOverflows    int `json:"from_net_nb_buffer_overflows"`
		FromNetNbOutOfSeqPackets    int `json:"from_net_nb_out_of_seq_packets"`
		FromNetNbEarlyLatePackets   int `json:"from_net_nb_early_late_packets"`
		FromNetNbDuplicatePackets   int `json:"from_net_nb_duplicate_packets"`
	} `json:"rtp_statistics_struct"`
	ResetRtpStats                       string `json:"reset_rtp_stats"`
	TotalOutgoingFilePlaybacks          int    `json:"total_outgoing_file_playbacks"`
	InstOutgoingInterceptions           int    `json:"inst_outgoing_interceptions"`
	TotalIncomingFileRecordings         int    `json:"total_incoming_file_recordings"`
	InstOutgoingCallRateAcceptedHighest int    `json:"inst_outgoing_call_rate_accepted_highest"`
	InstIncomingCallRateAcceptedHighest int    `json:"inst_incoming_call_rate_accepted_highest"`
	InstIncomingCallRateHighest         int    `json:"inst_incoming_call_rate_highest"`
}

/*func (n *NapStatus) Get(sbc Cfg, nap Nap) (*NapStatus, error) {
	http, err := GetHttp(sbc.Host, "/configurations/"+sbc.Config+"/naps/"+nap.Name, sbc.User, sbc.Password)
	if err != nil {
		return nil, err
	}
	log.Info(string(http))

	var napStatus *NapStatus
	err = json.Unmarshal(http, &nap)
	if err != nil {
		return nil, err
	}
	return napStatus, nil
}*/
