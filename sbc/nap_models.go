package sbc

type NapCallRateLimiting struct {
	ProcessingDelayHighThreshold     string `json:"processing_delay_high_threshold,omitempty"`
	MaximumCallsPerSecond            int    `json:"maximum_calls_per_second,omitempty"`
	MaximumCallBurst                 int    `json:"maximum_call_burst,omitempty"`
	MaximumSimultaneousOutgoingCalls int    `json:"maximum_simultaneous_outgoing_calls,omitempty"`
	MaximumSimultaneousIncomingCalls int    `json:"maximum_simultaneous_incoming_calls,omitempty"`
	MaximumOutgoingCallsPerSecond    int    `json:"maximum_outgoing_calls_per_second,omitempty"`
	MaximumIncomingCallsPerSecond    int    `json:"maximum_incoming_calls_per_second,omitempty"`
	MaximumSimultaneousTotalCalls    int    `json:"maximum_simultaneous_total_calls,omitempty"`
	ProcessingDelayLowThreshold      string `json:"processing_delay_low_threshold,omitempty"`
}

type NapRegistrationParams struct {
	AddressToRegister string `json:"address_to_register,omitempty"`
	RegisterToProxy   bool   `json:"register_to_proxy,omitempty"`
}

type NapSipiParams struct {
	IsupProtocolVariant    string `json:"isup_protocol_variant,omitempty"`
	AppendFToOutgoingCalls bool   `json:"append_f_to_outgoing_calls,omitempty"`
	Enable                 bool   `json:"enable,omitempty"`
	ContentType            string `json:"content_type,omitempty"`
	CallProgressMethod     string `json:"call_progress_method,omitempty"`
}

type NapAdvancedParams struct {
	MapAnyResponseToAvailableStatus bool   `json:"map_any_response_to_available_status,omitempty"`
	ResponseTimeout                 string `json:"response_timeout,omitempty"`
	PrivacyType                     string `json:"privacy_type,omitempty"`
	ProxyPollingMaxForwards         int    `json:"proxy_polling_max_forwards,omitempty"`
	TriggersCallProgress            bool   `json:"183_triggers_call_progress,omitempty"`
}

type NapFilterParams struct {
	FilterByProxyPort    bool `json:"filter_by_proxy_port,omitempty"`
	FilterByLocalPort    bool `json:"filter_by_local_port,omitempty"`
	FilterByProxyAddress bool `json:"filter_by_proxy_address,omitempty"`
}

type NapAuthParams struct {
	User           string `json:"user,omitempty"`
	Realm          string `json:"realm,omitempty"`
	ReuseChallenge bool   `json:"reuse_challenge,omitempty"`
	Password       string `json:"password,omitempty"`
	IgnoreRealm    bool   `json:"ignore_realm,omitempty"`
}

type NapNatParams struct {
	RemoteMethodSip string `json:"remote_method_sip,omitempty"`
	LocalMethodRtp  string `json:"local_method_rtp,omitempty"`
	RemoteMethodRtp string `json:"remote_method_rtp,omitempty"`
	LocalMethodSip  string `json:"local_method_sip,omitempty"`
}

type NapSipCfg struct {
	PollRemoteProxy           bool                  `json:"poll_remote_proxy,omitempty"`
	AcceptOnlyAuthorizedUsers bool                  `json:"accept_only_authorized_users,omitempty"`
	RegistrationParameters    NapRegistrationParams `json:"registration_parameters,omitempty"`
	SipiParameters            NapSipiParams         `json:"sipi_parameters,omitempty"`
	AdvancedParameters        NapAdvancedParams     `json:"advanced_parameters,omitempty"`
	ProxyPortType             string                `json:"proxy_port_type,omitempty"`
	NapSipAcls                []interface{}         `json:"nap_sip_acls,omitempty"`
	SipUseProxy               bool                  `json:"sip_use_proxy,omitempty"`
	ProxyPort                 int                   `json:"proxy_port,omitempty"`
	FilteringParameters       NapFilterParams       `json:"filtering_parameters,omitempty"`
	ProxyPollingInterval      string                `json:"proxy_polling_interval,omitempty"`
	ProxyAddress              string                `json:"proxy_address,omitempty"`
	AuthenticationParameters  NapAuthParams         `json:"authentication_parameters,omitempty"`
	NetworkAddressTranslation NapNatParams          `json:"network_address_translation,omitempty"`
}

type NapCongestionThreshold struct {
	PeriodDuration   string `json:"period_duration,omitempty"`
	NbCallsPerPeriod int    `json:"nb_calls_per_period,omitempty"`
}

type NapColumnValues struct {
	RoutesetsDefinition string `json:"routesets_definition"`
	RouteGroups         string `json:"route_groups"`
	RoutesetsDigitmap   string `json:"routesets_digitmap"`
	Weight              string `json:"weight"`
	BlackWhiteList      string `json:"black_white_list"`
	CalledPreRemap      string `json:"called_pre_remap"`
	Priority            string `json:"priority"`
}

// Nap GET /configurations/config_1/naps/pbx_dec0de/
type Nap struct {
	Name                string                 `json:"name,omitempty"`
	CallRateLimiting    NapCallRateLimiting    `json:"call_rate_limiting,omitempty"`
	Enabled             bool                   `json:"enabled,omitempty"`
	DefaultProfile      string                 `json:"default_profile,omitempty"`
	PortRanges          []string               `json:"port_ranges,omitempty"`
	SipTransportServers []string               `json:"sip_transport_servers,omitempty"`
	SipCfg              NapSipCfg              `json:"sip_cfg,omitempty"`
	CongestionThreshold NapCongestionThreshold `json:"congestion_threshold,omitempty"`
}

type NapStatus struct {
	AvailabilityDetectionStruct struct {
		PollRemoteProxy string `json:"poll_remote_proxy,omitempty"`
		IsAvailable     string `json:"is_available,omitempty"`
	} `json:"availability_detection_struct,omitempty"`
	PortRangeSharedUsagePercent    int    `json:"port_range_shared_usage_percent,omitempty"`
	AvailableCnt                   int    `json:"available_cnt,omitempty"`
	InstIncomingCallCntTerminating int    `json:"inst_incoming_call_cnt_terminating,omitempty"`
	InstIncomingCallCntAnswered    int    `json:"inst_incoming_call_cnt_answered,omitempty"`
	SignalingType                  string `json:"signaling_type,omitempty"`
	TotalIncomingFilePlaybacks     int    `json:"total_incoming_file_playbacks,omitempty"`
	InstOutgoingCallCnt            int    `json:"inst_outgoing_call_cnt,omitempty"`
	InstIncomingEmergencyCallCnt   int    `json:"inst_incoming_emergency_call_cnt,omitempty"`
	ResetAsrStats                  string `json:"reset_asr_stats,omitempty"`
	InstOutgoingCallRate           int    `json:"inst_outgoing_call_rate,omitempty"`
	InstIncomingCallRateAnswered   int    `json:"inst_incoming_call_rate_answered,omitempty"`
	InstIncomingCallRateAccepted   int    `json:"inst_incoming_call_rate_accepted,omitempty"`
	FirewallBlockedCnt             int    `json:"firewall_blocked_cnt,omitempty"`
	ResetStats                     string `json:"reset_stats,omitempty"`
	ResetNapDropStats              string `json:"reset_nap_drop_stats,omitempty"`
	AsrStatsIncomingStruct         struct {
		Last24HCallCnt        int `json:"last_24h_call_cnt,omitempty"`
		Last24HAsrPercent     int `json:"last_24h_asr_percent,omitempty"`
		TotalCallCnt          int `json:"total_call_cnt,omitempty"`
		GlobalAsrPercent      int `json:"global_asr_percent,omitempty"`
		LastHourCallCnt       int `json:"last_hour_call_cnt,omitempty"`
		CurrentHourCallCnt    int `json:"current_hour_call_cnt,omitempty"`
		TotalAnsweredCallCnt  int `json:"total_answered_call_cnt,omitempty"`
		TotalAcceptedCallCnt  int `json:"total_accepted_call_cnt,omitempty"`
		LastHourAsrPercent    int `json:"last_hour_asr_percent,omitempty"`
		CurrentHourAsrPercent int `json:"current_hour_asr_percent,omitempty"`
	} `json:"asr_stats_incoming_struct,omitempty"`
	UsagePercent                         int `json:"usage_percent,omitempty"`
	TotalIncomingInterceptions           int `json:"total_incoming_interceptions,omitempty"`
	InstIncomingFilePlaybacks            int `json:"inst_incoming_file_playbacks,omitempty"`
	InstOutgoingCallCntAnswered          int `json:"inst_outgoing_call_cnt_answered,omitempty"`
	InstIncomingEmergencyCallRateHighest int `json:"inst_incoming_emergency_call_rate_highest,omitempty"`
	UniqueId                             int `json:"unique_id,omitempty"`
	SystemDropStats                      struct {
	} `json:"system_drop_stats,omitempty"`
	LocalDropStats struct {
	} `json:"local_drop_stats,omitempty"`
	MosStruct struct {
		CurrentHourEgress  float64 `json:"current_hour_egress,omitempty"`
		LastHourEgress     float64 `json:"last_hour_egress,omitempty"`
		CurrentHourIngress float64 `json:"current_hour_ingress,omitempty"`
		LastHourIngress    float64 `json:"last_hour_ingress,omitempty"`
		Last24HIngress     float64 `json:"last_24h_ingress,omitempty"`
		Last24HEgress      float64 `json:"last_24h_egress,omitempty"`
	} `json:"mos_struct,omitempty"`
	SipSharedUsagePercent               int `json:"sip_shared_usage_percent,omitempty"`
	InstIncomingCallRateAnsweredHighest int `json:"inst_incoming_call_rate_answered_highest,omitempty"`
	InstIncomingCallCnt                 int `json:"inst_incoming_call_cnt,omitempty"`
	TotalOutgoingFileRecordings         int `json:"total_outgoing_file_recordings,omitempty"`
	InstOutgoingCallRateAnsweredHighest int `json:"inst_outgoing_call_rate_answered_highest,omitempty"`
	InstIncomingCallRate                int `json:"inst_incoming_call_rate,omitempty"`
	InstIncomingCallCntInProgress       int `json:"inst_incoming_call_cnt_in_progress,omitempty"`
	RemoteDropStats                     struct {
	} `json:"remote_drop_stats,omitempty"`
	AvailabilityPercent              int  `json:"availability_percent,omitempty"`
	InstIncomingFileRecordings       int  `json:"inst_incoming_file_recordings,omitempty"`
	InstOutgoingCallRateAccepted     int  `json:"inst_outgoing_call_rate_accepted,omitempty"`
	FirewallBlocked                  bool `json:"firewall_blocked,omitempty"`
	CallCongestionPeriodDroppedCalls int  `json:"call_congestion_period_dropped_calls,omitempty"`
	RegistrationStruct               struct {
		Registered      string `json:"registered,omitempty"`
		RegisterToProxy string `json:"register_to_proxy,omitempty"`
	} `json:"registration_struct,omitempty"`
	NetworkQualityStruct struct {
		CurrentHourEgress  int `json:"current_hour_egress,omitempty"`
		LastHourEgress     int `json:"last_hour_egress,omitempty"`
		CurrentHourIngress int `json:"current_hour_ingress,omitempty"`
		LastHourIngress    int `json:"last_hour_ingress,omitempty"`
		Last24HIngress     int `json:"last_24h_ingress,omitempty"`
		Last24HEgress      int `json:"last_24h_egress,omitempty"`
	} `json:"network_quality_struct,omitempty"`
	AsrStatsOutgoingStruct struct {
		Last24HCallCnt        int `json:"last_24h_call_cnt,omitempty"`
		Last24HAsrPercent     int `json:"last_24h_asr_percent,omitempty"`
		TotalCallCnt          int `json:"total_call_cnt,omitempty"`
		GlobalAsrPercent      int `json:"global_asr_percent,omitempty"`
		LastHourCallCnt       int `json:"last_hour_call_cnt,omitempty"`
		CurrentHourCallCnt    int `json:"current_hour_call_cnt,omitempty"`
		TotalAnsweredCallCnt  int `json:"total_answered_call_cnt,omitempty"`
		TotalAcceptedCallCnt  int `json:"total_accepted_call_cnt,omitempty"`
		LastHourAsrPercent    int `json:"last_hour_asr_percent,omitempty"`
		CurrentHourAsrPercent int `json:"current_hour_asr_percent,omitempty"`
	} `json:"asr_stats_outgoing_struct,omitempty"`
	InstOutgoingCallRateHighest          int  `json:"inst_outgoing_call_rate_highest,omitempty"`
	InstIncomingEmergencyCallRate        int  `json:"inst_incoming_emergency_call_rate,omitempty"`
	LowDelayRelaySharedUsagePercent      int  `json:"low_delay_relay_shared_usage_percent,omitempty"`
	TotalOutgoingInterceptions           int  `json:"total_outgoing_interceptions,omitempty"`
	InstOutgoingFilePlaybacks            int  `json:"inst_outgoing_file_playbacks,omitempty"`
	InstIncomingInterceptions            int  `json:"inst_incoming_interceptions,omitempty"`
	CallCongestion                       bool `json:"call_congestion,omitempty"`
	MipsSharedUsagePercent               int  `json:"mips_shared_usage_percent,omitempty"`
	SharedUsagePercent                   int  `json:"shared_usage_percent,omitempty"`
	UnavailableCnt                       int  `json:"unavailable_cnt,omitempty"`
	InstOutgoingFileRecordings           int  `json:"inst_outgoing_file_recordings,omitempty"`
	InstOutgoingCallRateAnswered         int  `json:"inst_outgoing_call_rate_answered,omitempty"`
	InstOutgoingCallCntTerminating       int  `json:"inst_outgoing_call_cnt_terminating,omitempty"`
	InstIncomingEmergencyCallCntAnswered int  `json:"inst_incoming_emergency_call_cnt_answered,omitempty"`
	RtpStatisticsStruct                  struct {
		FromNetNbOtherErrors        int `json:"from_net_nb_other_errors,omitempty"`
		FromNetNbLostPackets        int `json:"from_net_nb_lost_packets,omitempty"`
		T38NbPagesFromTdm           int `json:"t38_nb_pages_from_tdm,omitempty"`
		FromNetNbBadProtocolHeaders int `json:"from_net_nb_bad_protocol_headers,omitempty"`
		FromNetNbPackets            int `json:"from_net_nb_packets,omitempty"`
		ToNetNbPackets              int `json:"to_net_nb_packets,omitempty"`
		T38NbPagesToTdm             int `json:"t38_nb_pages_to_tdm,omitempty"`
		ToNetNbArpFailures          int `json:"to_net_nb_arp_failures,omitempty"`
		FromNetNbBufferOverflows    int `json:"from_net_nb_buffer_overflows,omitempty"`
		FromNetNbOutOfSeqPackets    int `json:"from_net_nb_out_of_seq_packets,omitempty"`
		FromNetNbEarlyLatePackets   int `json:"from_net_nb_early_late_packets,omitempty"`
		FromNetNbDuplicatePackets   int `json:"from_net_nb_duplicate_packets,omitempty"`
	} `json:"rtp_statistics_struct,omitempty"`
	ResetRtpStats                       string `json:"reset_rtp_stats,omitempty"`
	TotalOutgoingFilePlaybacks          int    `json:"total_outgoing_file_playbacks,omitempty"`
	InstOutgoingInterceptions           int    `json:"inst_outgoing_interceptions,omitempty"`
	TotalIncomingFileRecordings         int    `json:"total_incoming_file_recordings,omitempty"`
	InstOutgoingCallRateAcceptedHighest int    `json:"inst_outgoing_call_rate_accepted_highest,omitempty"`
	InstIncomingCallRateAcceptedHighest int    `json:"inst_incoming_call_rate_accepted_highest,omitempty"`
	InstIncomingCallRateHighest         int    `json:"inst_incoming_call_rate_highest,omitempty"`
}
