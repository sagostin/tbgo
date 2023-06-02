package sbc

// TBStatus client
type TBStatus struct {
	Client
}

// TBStatus constructor (from Client)
func (c Client) TBStatus() TBStatus {
	return TBStatus{
		Client: c,
	}
}

func (c TBStatus) GetStatus() (*Status, error) {
	var nap *Status

	err := c.Client.Request("GET", "/status", nil, &nap)
	if err != nil {
		return nil, err
	}

	return nap, nil
}

type Status struct {
	Sips struct {
		WarningThreshold  int `json:"warning_threshold"`
		HealthySipCfgList struct {
			Name string `json:"name"`
		} `json:"healthy_sip_cfg_list"`
		NotConfigured  bool `json:"not_configured"`
		CfgCnt         int  `json:"cfg_cnt"`
		AlarmCnt       int  `json:"alarm_cnt"`
		ErrorThreshold int  `json:"error_threshold"`
		HealthyCnt     int  `json:"healthy_cnt"`
		FailingCnt     int  `json:"failing_cnt"`
	} `json:"sips"`
	SipSimulators struct {
		FullyAvailableSipSimulatorCnt     int `json:"fully_available_sip_simulator_cnt"`
		WarningThreshold                  int `json:"warning_threshold"`
		CumulNbSipMsg                     int `json:"cumul_nb_sip_msg"`
		ErrorThreshold                    int `json:"error_threshold"`
		CumulNbCall                       int `json:"cumul_nb_call"`
		UnavailableSipSimulatorCnt        int `json:"unavailable_sip_simulator_cnt"`
		PartiallyAvailableSipSimulatorCnt int `json:"partially_available_sip_simulator_cnt"`
	} `json:"sip_simulators"`
	CircuitGroups struct {
		WarningThreshold          int  `json:"warning_threshold"`
		DownCnt                   int  `json:"down_cnt"`
		NotConfigured             bool `json:"not_configured"`
		IncomingCnt               int  `json:"incoming_cnt"`
		LocallyRemotelyBlockedCnt int  `json:"locally_remotely_blocked_cnt"`
		OutgoingCnt               int  `json:"outgoing_cnt"`
		RemotelyBlockedCnt        int  `json:"remotely_blocked_cnt"`
		ErrorThreshold            int  `json:"error_threshold"`
		SuspendedCnt              int  `json:"suspended_cnt"`
		ResetCnt                  int  `json:"reset_cnt"`
		LocallyBlockedCnt         int  `json:"locally_blocked_cnt"`
		IdleCnt                   int  `json:"idle_cnt"`
	} `json:"circuit_groups"`
	Calls struct {
		MatchFilterCallCnt int `json:"match_filter_call_cnt"`
		ActiveCallCnt      int `json:"active_call_cnt"`
	} `json:"calls"`
	Hosts struct {
		WarningThreshold        int `json:"warning_threshold"`
		UpdatesAvailableCnt     int `json:"updates_available_cnt"`
		AuthorizationExpiredCnt int `json:"authorization_expired_cnt"`
		MissingPrerequisitesCnt int `json:"missing_prerequisites_cnt"`
		NotReadyHostList        struct {
		} `json:"not_ready_host_list"`
		AuthorizationAboutToExpireHostList struct {
		} `json:"authorization_about_to_expire_host_list"`
		RebootRequiredCnt            int `json:"reboot_required_cnt"`
		MissingPrerequisitesHostList struct {
		} `json:"missing_prerequisites_host_list"`
		InvalidLicenseCnt        int `json:"invalid_license_cnt"`
		UpdatesAvailableHostList struct {
		} `json:"updates_available_host_list"`
		ReadyCnt            int `json:"ready_cnt"`
		FeatureMismatchList struct {
		} `json:"feature_mismatch_list"`
		ReadyHostList struct {
			Name string `json:"name"`
		} `json:"ready_host_list"`
		InvalidLicenseHostList struct {
		} `json:"invalid_license_host_list"`
		ErrorThreshold                int `json:"error_threshold"`
		AuthorizationAboutToExpireCnt int `json:"authorization_about_to_expire_cnt"`
		SchedulingProblemHostList     struct {
		} `json:"scheduling_problem_host_list"`
		RebootRequiredHostList struct {
		} `json:"reboot_required_host_list"`
		FeatureMismatchCnt             int `json:"feature_mismatch_cnt"`
		LicenseExpiredCnt              int `json:"license_expired_cnt"`
		InvalidWebPortalConfigHostList struct {
		} `json:"invalid_web_portal_config_host_list"`
		InvalidWebPortalConfigCnt    int `json:"invalid web_portal_config_cnt"`
		SchedulingProblemCnt         int `json:"scheduling_problem_cnt"`
		LicenseAboutToExpireCnt      int `json:"license_about_to_expire_cnt"`
		NotReadyCnt                  int `json:"not_ready_cnt"`
		LicenseAboutToExpireHostList struct {
		} `json:"license_about_to_expire_host_list"`
		AuthorizationExpiredHostList struct {
		} `json:"authorization_expired_host_list"`
		LicenseExpiredHostList struct {
		} `json:"license_expired_host_list"`
	} `json:"hosts"`
	HostEthernetBondings struct {
		WarningThreshold int `json:"warning_threshold"`
		DownCnt          int `json:"down_cnt"`
		ActiveCnt        int `json:"active_cnt"`
		BackupCnt        int `json:"backup_cnt"`
		ErrorThreshold   int `json:"error_threshold"`
	} `json:"host_ethernet_bondings"`
	IsdnStacks struct {
		WarningThreshold int  `json:"warning_threshold"`
		DownCnt          int  `json:"down_cnt"`
		PartiallyUpCnt   int  `json:"partially_up_cnt"`
		NotConfigured    bool `json:"not_configured"`
		UpCnt            int  `json:"up_cnt"`
		ErrorThreshold   int  `json:"error_threshold"`
	} `json:"isdn_stacks"`
	Intercept struct {
		WarningThreshold         int `json:"warning_threshold"`
		IriBlockedQueueErrorList struct {
		} `json:"iri_blocked_queue_error_list"`
		IriServerConnectionFailureErrorList struct {
		} `json:"iri_server_connection_failure_error_list"`
		ErrorThreshold     int `json:"error_threshold"`
		IriUploadErrorList struct {
		} `json:"iri_upload_error_list"`
		InterceptionLegErrorList struct {
		} `json:"interception_leg_error_list"`
	} `json:"intercept"`
	Dns struct {
		WarningThreshold               int `json:"warning_threshold"`
		AvailableDnsGroupCnt           int `json:"available_dns_group_cnt"`
		PartiallyAvailableDnsGroupList struct {
			Name string `json:"name"`
		} `json:"partially_available_dns_group_list"`
		PartiallyAvailableDnsGroupCnt int `json:"partially_available_dns_group_cnt"`
		ErrorThreshold                int `json:"error_threshold"`
		UnavailableDnsGroupCnt        int `json:"unavailable_dns_group_cnt"`
	} `json:"dns"`
	SipSaps struct {
		WarningThreshold int `json:"warning_threshold"`
		SapCnt           int `json:"sap_cnt"`
		SipMap           struct {
			Name string `json:"name"`
		} `json:"sip_map"`
		ErrorThreshold int `json:"error_threshold"`
		SIPSTACK       struct {
			SipSapList struct {
				Name []string `json:"name"`
			} `json:"sip_sap_list"`
		} `json:"SIP_STACK"`
	} `json:"sip_saps"`
	HostFilesystems struct {
		WarningThreshold int `json:"warning_threshold"`
		DiskFullCnt      int `json:"disk_full_cnt"`
		ErrorThreshold   int `json:"error_threshold"`
		DiskFullHostList struct {
		} `json:"disk_full_host_list"`
	} `json:"host_filesystems"`
	HostEthernetPorts struct {
		HostMap struct {
			Name string `json:"name"`
		} `json:"host_map"`
		DownCnt         int `json:"down_cnt"`
		UpCnt           int `json:"up_cnt"`
		Telcobridgespro struct {
			UpPortList struct {
				Name []string `json:"name"`
			} `json:"up_port_list"`
		} `json:"telcobridgespro"`
		ErrorThreshold int `json:"error_threshold"`
	} `json:"host_ethernet_ports"`
	Adapters struct {
		WarningThreshold  int `json:"warning_threshold"`
		UpdatingCnt       int `json:"updating_cnt"`
		DownCnt           int `json:"down_cnt"`
		ConfiguringCnt    int `json:"configuring_cnt"`
		RebootRequiredCnt int `json:"reboot_required_cnt"`
		FaultCnt          int `json:"fault_cnt"`
		UpCnt             int `json:"up_cnt"`
		PowerFaultCnt     int `json:"power_fault_cnt"`
		Np1UnusableCnt    int `json:"np1_unusable_cnt"`
		Meta              struct {
			Version string `json:"version"`
			SrcPath string `json:"src_path"`
		} `json:"***meta***"`
		ErrorThreshold          int `json:"error_threshold"`
		DisabledCnt             int `json:"disabled_cnt"`
		LicenseExpiredCnt       int `json:"license_expired_cnt"`
		LicenseAboutToExpireCnt int `json:"license_about_to_expire_cnt"`
		InvalidBoardModeCnt     int `json:"invalid_board_mode_cnt"`
		CorruptedFileSystemCnt  int `json:"corrupted_file_system_cnt"`
	} `json:"adapters"`
	Naps struct {
		WarningThreshold      int `json:"warning_threshold"`
		FirewallBlockedNapCnt int `json:"firewall_blocked_nap_cnt"`
		CallCongestionNapCnt  int `json:"call_congestion_nap_cnt"`
		AvailableNapList      struct {
			Name []string `json:"name"`
		} `json:"available_nap_list"`
		AvailableNapCnt          int  `json:"available_nap_cnt"`
		NotConfigured            bool `json:"not_configured"`
		ErrorThreshold           int  `json:"error_threshold"`
		UnavailableNapCnt        int  `json:"unavailable_nap_cnt"`
		PartiallyAvailableNapCnt int  `json:"partially_available_nap_cnt"`
	} `json:"naps"`
	IsdnAnalyzers struct {
		UnavailableIsdnAnalyzerCnt        int `json:"unavailable_isdn_analyzer_cnt"`
		WarningThreshold                  int `json:"warning_threshold"`
		CurrentNbCallTerminating          int `json:"current_nb_call_terminating"`
		PartiallyAvailableIsdnAnalyzerCnt int `json:"partially_available_isdn_analyzer_cnt"`
		CurrentNbCallAnswered             int `json:"current_nb_call_answered"`
		CumulNbCallTerminatedTimeout      int `json:"cumul_nb_call_terminated_timeout"`
		CumulNbCallAnswered               int `json:"cumul_nb_call_answered"`
		ErrorThreshold                    int `json:"error_threshold"`
		CurrentNbCallSetup                int `json:"current_nb_call_setup"`
		CurrentNbCall                     int `json:"current_nb_call"`
		CumulNbCallTerminatedCollision    int `json:"cumul_nb_call_terminated_collision"`
		CumulNbCall                       int `json:"cumul_nb_call"`
		CumulNbCallTerminatedNetwork      int `json:"cumul_nb_call_terminated_network"`
		FullyAvailableIsdnAnalyzerCnt     int `json:"fully_available_isdn_analyzer_cnt"`
		CumulNbCallSetup                  int `json:"cumul_nb_call_setup"`
	} `json:"isdn_analyzers"`
	HostIpInterfaces struct {
		UnavailableCnt int `json:"unavailable_cnt"`
		HostMap        struct {
			Name string `json:"name"`
		} `json:"host_map"`
		Telcobridgespro struct {
			AvailableIpInterfaceList struct {
				Name []string `json:"name"`
			} `json:"available_ip_interface_list"`
		} `json:"telcobridgespro"`
		ErrorThreshold int `json:"error_threshold"`
		AvailableCnt   int `json:"available_cnt"`
	} `json:"host_ip_interfaces"`
	Database struct {
		SecondaryReady string `json:"secondary_ready"`
		PrimaryReady   string `json:"primary_ready"`
		ErrorThreshold int    `json:"error_threshold"`
	} `json:"database"`
	SipRegistrations struct {
		WarningThreshold int `json:"warning_threshold"`
		ErrorThreshold   int `json:"error_threshold"`
	} `json:"sip_registrations"`
	Firewalls struct {
		WarningThreshold  int `json:"warning_threshold"`
		RebootRequiredCnt int `json:"reboot_required_cnt"`
		ListFullSbcList   struct {
		} `json:"list_full_sbc_list"`
		ReadySbcList struct {
			Name string `json:"name"`
		} `json:"ready_sbc_list"`
		SubOptimalConfigCnt int `json:"sub_optimal_config_cnt"`
		ListFullCnt         int `json:"list_full_cnt"`
		ReadyCnt            int `json:"ready_cnt"`
		HighCpuSbcList      struct {
		} `json:"high_cpu_sbc_list"`
		DdosSbcList struct {
		} `json:"ddos_sbc_list"`
		ErrorThreshold   int `json:"error_threshold"`
		DdosCnt          int `json:"ddos_cnt"`
		ThresholdSbcList struct {
		} `json:"threshold_sbc_list"`
		WarningSbcList struct {
		} `json:"warning_sbc_list"`
		RebootRequiredSbcList struct {
		} `json:"reboot_required_sbc_list"`
		WarningCnt              int `json:"warning_cnt"`
		HighCpuCnt              int `json:"high_cpu_cnt"`
		ThresholdCnt            int `json:"threshold_cnt"`
		SchedulingProblemCnt    int `json:"scheduling_problem_cnt"`
		NotReadyCnt             int `json:"not_ready_cnt"`
		SubOptimalConfigSbcList struct {
		} `json:"sub_optimal_config_sbc_list"`
		SchedulingProblemSbcList struct {
		} `json:"scheduling_problem_sbc_list"`
		NotReadySbcList struct {
		} `json:"not_ready_sbc_list"`
	} `json:"firewalls"`
	Certificate struct {
		WarningThreshold int `json:"warning_threshold"`
		ValidCnt         int `json:"valid_cnt"`
		AboutToExpireCnt int `json:"about_to_expire_cnt"`
		ErrorThreshold   int `json:"error_threshold"`
		InvalidCnt       int `json:"invalid_cnt"`
		ValidList        struct {
			Name []string `json:"name"`
		} `json:"valid_list"`
	} `json:"certificate"`
	HostApplications struct {
		WarningThreshold int `json:"warning_threshold"`
		UpdatingCnt      int `json:"updating_cnt"`
		HostMap          struct {
			Name string `json:"name"`
		} `json:"host_map"`
		FaultCnt        int `json:"fault_cnt"`
		Telcobridgespro struct {
			ReadyList struct {
				Name []string `json:"name"`
			} `json:"ready_list"`
		} `json:"telcobridgespro"`
		ReadyCnt              int `json:"ready_cnt"`
		ErrorThreshold        int `json:"error_threshold"`
		NotRunningHostDownCnt int `json:"not_running_host_down_cnt"`
		NotRunningCnt         int `json:"not_running_cnt"`
	} `json:"host_applications"`
	Ha struct {
		HaDbReplicationState      string `json:"ha_db_replication_state"`
		HaRepositoryH248WarmStart string `json:"ha_repository_h248_warm_start"`
		HaStandbyApplications     string `json:"ha_standby_applications"`
		HaStandbyHost             string `json:"ha_standby_host"`
		ErrorThreshold            int    `json:"error_threshold"`
		HaSystemSwitchover        string `json:"ha_system_switchover"`
	} `json:"ha"`
	Clock struct {
		WarningThreshold          int `json:"warning_threshold"`
		ActiveClockReferencesList struct {
		} `json:"active_clock_references_list"`
		UnqualifiedClockReferencesList struct {
		} `json:"unqualified_clock_references_list"`
		NotConfigured                     bool `json:"not_configured"`
		BeingQualifiedClockReferencesList struct {
		} `json:"being_qualified_clock_references_list"`
		QualifiedClockReferencesList struct {
		} `json:"qualified_clock_references_list"`
		ErrorThreshold int `json:"error_threshold"`
	} `json:"clock"`
	Bits struct {
		WarningThreshold int  `json:"warning_threshold"`
		DownCnt          int  `json:"down_cnt"`
		NotConfigured    bool `json:"not_configured"`
		UpCnt            int  `json:"up_cnt"`
		ErrorThreshold   int  `json:"error_threshold"`
	} `json:"bits"`
	CallLegs struct {
		IncomingLegsHighest                int    `json:"incoming_legs_highest"`
		ResetStats                         string `json:"reset_stats"`
		OutgoingLegsAnsweredRate           int    `json:"outgoing_legs_answered_rate"`
		TotalOutgoingLegsAccepted          int    `json:"total_outgoing_legs_accepted"`
		FilePlaybacks                      int    `json:"file_playbacks"`
		OutgoingLegsAnsweredRateHighest    int    `json:"outgoing_legs_answered_rate_highest"`
		IncomingLegs                       int    `json:"incoming_legs"`
		CallLegsAnsweredRate               int    `json:"call_legs_answered_rate"`
		OutgoingLegsAnsweredHighest        int    `json:"outgoing_legs_answered_highest"`
		TotalMediaOnlyLegs                 int    `json:"total_media_only_legs"`
		OutgoingLegsRate                   int    `json:"outgoing_legs_rate"`
		TotalOutgoingLegsFilePlaybacks     int    `json:"total_outgoing_legs_file_playbacks"`
		IncomingLegsAnsweredRate           int    `json:"incoming_legs_answered_rate"`
		TotalIncomingEmergencyLegs         int    `json:"total_incoming_emergency_legs"`
		OutgoingLegs                       int    `json:"outgoing_legs"`
		MediaOnlyLegs                      int    `json:"media_only_legs"`
		FileRecordings                     int    `json:"file_recordings"`
		MediaOnlyLegsHighest               int    `json:"media_only_legs_highest"`
		IncomingLegsAnsweredRateHighest    int    `json:"incoming_legs_answered_rate_highest"`
		TotalIncomingLegsAnswered          int    `json:"total_incoming_legs_answered"`
		IncomingLegsAcceptedRateHighest    int    `json:"incoming_legs_accepted_rate_highest"`
		CallLegsSetup                      int    `json:"call_legs_setup"`
		IncomingLegsRate                   int    `json:"incoming_legs_rate"`
		TotalOutgoingLegsFileRecordings    int    `json:"total_outgoing_legs_file_recordings"`
		TotalCallLegsAccepted              int    `json:"total_call_legs_accepted"`
		OutgoingLegsSetup                  int    `json:"outgoing_legs_setup"`
		TotalIncomingLegsFileRecordings    int    `json:"total_incoming_legs_file_recordings"`
		IncomingLegsAcceptedRate           int    `json:"incoming_legs_accepted_rate"`
		OutgoingLegsTerminating            int    `json:"outgoing_legs_terminating"`
		OutgoingLegsAcceptedRateHighest    int    `json:"outgoing_legs_accepted_rate_highest"`
		CallLegsAcceptedRateHighest        int    `json:"call_legs_accepted_rate_highest"`
		TotalIncomingLegsFilePlaybacks     int    `json:"total_incoming_legs_file_playbacks"`
		CallLegsAnsweredHighest            int    `json:"call_legs_answered_highest"`
		IncomingLegsAnsweredHighest        int    `json:"incoming_legs_answered_highest"`
		IncomingEmergencyLegsRateHighest   int    `json:"incoming_emergency_legs_rate_highest"`
		CallLegsRate                       int    `json:"call_legs_rate"`
		CallLegsAnsweredRateHighest        int    `json:"call_legs_answered_rate_highest"`
		CallLegsRateHighest                int    `json:"call_legs_rate_highest"`
		CallLegsTerminating                int    `json:"call_legs_terminating"`
		OutgoingLegsAcceptedRate           int    `json:"outgoing_legs_accepted_rate"`
		IncomingEmergencyLegsRate          int    `json:"incoming_emergency_legs_rate"`
		TotalOutgoingLegsInterceptions     int    `json:"total_outgoing_legs_interceptions"`
		TotalIncomingEmergencyLegsAnswered int    `json:"total_incoming_emergency_legs_answered"`
		TotalMediaOnlyLegsFileRecordings   int    `json:"total_media_only_legs_file_recordings"`
		TotalIncomingLegsInterceptions     int    `json:"total_incoming_legs_interceptions"`
		TotalCallLegsInterceptions         int    `json:"total_call_legs_interceptions"`
		TotalCallLegsFilePlaybacks         int    `json:"total_call_legs_file_playbacks"`
		TotalOutgoingLegsAnswered          int    `json:"total_outgoing_legs_answered"`
		IncomingEmergencyLegsHighest       int    `json:"incoming_emergency_legs_highest"`
		TotalMediaOnlyLegsInterceptions    int    `json:"total_media_only_legs_interceptions"`
		TotalCallLegs                      int    `json:"total_call_legs"`
		TotalMediaOnlyLegsFilePlaybacks    int    `json:"total_media_only_legs_file_playbacks"`
		TotalIncomingLegs                  int    `json:"total_incoming_legs"`
		TotalIncomingLegsAccepted          int    `json:"total_incoming_legs_accepted"`
		IncomingLegsAnswered               int    `json:"incoming_legs_answered"`
		TotalOutgoingLegs                  int    `json:"total_outgoing_legs"`
		CallLegsAcceptedRate               int    `json:"call_legs_accepted_rate"`
		OutgoingLegsHighest                int    `json:"outgoing_legs_highest"`
		Interceptions                      int    `json:"interceptions"`
		MediaOnlyLegsRateHighest           int    `json:"media_only_legs_rate_highest"`
		IncomingEmergencyLegs              int    `json:"incoming_emergency_legs"`
		TotalCallLegsAnswered              int    `json:"total_call_legs_answered"`
		MediaOnlyLegsRate                  int    `json:"media_only_legs_rate"`
		OutgoingLegsAnswered               int    `json:"outgoing_legs_answered"`
		CallLegs                           int    `json:"call_legs"`
		IncomingLegsRateHighest            int    `json:"incoming_legs_rate_highest"`
		IncomingLegsSetup                  int    `json:"incoming_legs_setup"`
		TotalCallLegsFileRecordings        int    `json:"total_call_legs_file_recordings"`
		CallLegsAnswered                   int    `json:"call_legs_answered"`
		IncomingLegsTerminating            int    `json:"incoming_legs_terminating"`
		CallLegsHighest                    int    `json:"call_legs_highest"`
		OutgoingLegsRateHighest            int    `json:"outgoing_legs_rate_highest"`
	} `json:"call_legs"`
	TmsIp struct {
		WarningThreshold         int `json:"warning_threshold"`
		FromNetNbOutOfSeqPackets int `json:"from_net_nb_out_of_seq_packets"`
		FromNetNbPackets         int `json:"from_net_nb_packets"`
		TotalNbTmsIpTxConn       int `json:"total_nb_tms_ip_tx_conn"`
		TotalConnFailures        int `json:"total_conn_failures"`
		FromNetNbLostPackets     int `json:"from_net_nb_lost_packets"`
		TotalNbTmsIpRxConn       int `json:"total_nb_tms_ip_rx_conn"`
		HighestNbTmsIpRxConn     int `json:"highest_nb_tms_ip_rx_conn"`
		FromNetNbOtherErrors     int `json:"from_net_nb_other_errors"`
		FromNetDuration          int `json:"from_net_duration"`
		FromNetNbOctets          int `json:"from_net_nb_octets"`
		CurrentNbTmsIpTxConn     int `json:"current_nb_tms_ip_tx_conn"`
		CurrentNbTmsIpRxConn     int `json:"current_nb_tms_ip_rx_conn"`
		ErrorThreshold           int `json:"error_threshold"`
		FromNetNbBufferOverflows int `json:"from_net_nb_buffer_overflows"`
		FromNetNbLatePackets     int `json:"from_net_nb_late_packets"`
		HighestNbTmsIpTxConn     int `json:"highest_nb_tms_ip_tx_conn"`
		MaxNbTmsIpConn           int `json:"max_nb_tms_ip_conn"`
	} `json:"tms_ip"`
	LineServicesLeaf struct {
		WarningThreshold     int `json:"warning_threshold"`
		AvailableTrunkCnt    int `json:"available_trunk_cnt"`
		ErrorThreshold       int `json:"error_threshold"`
		UnavailableTrunkList struct {
		} `json:"unavailable_trunk_list"`
		UnavailableTrunkCnt int `json:"unavailable_trunk_cnt"`
	} `json:"line_services_leaf"`
	AdaptersUsage struct {
		WarningThreshold           int `json:"warning_threshold"`
		HighCpuUsageCnt            int `json:"high_cpu_usage_cnt"`
		HighVoipUsageCnt           int `json:"high_voip_usage_cnt"`
		HighLowDelayRelayUsageList struct {
		} `json:"high_low_delay_relay_usage_list"`
		HighDspUsageCnt    int `json:"high_dsp_usage_cnt"`
		HighTmsIpUsageList struct {
		} `json:"high_tms_ip_usage_list"`
		CallCongestionCnt  int `json:"call_congestion_cnt"`
		CallCongestionList struct {
		} `json:"call_congestion_list"`
		ErrorThreshold    int `json:"error_threshold"`
		HighTmsIpUsageCnt int `json:"high_tms_ip_usage_cnt"`
		HighDspUsageList  struct {
		} `json:"high_dsp_usage_list"`
		HighCpuUsageList struct {
		} `json:"high_cpu_usage_list"`
		HighLowDelayRelayUsageCnt int `json:"high_low_delay_relay_usage_cnt"`
		HighVoipUsageList         struct {
		} `json:"high_voip_usage_list"`
	} `json:"adapters_usage"`
}
