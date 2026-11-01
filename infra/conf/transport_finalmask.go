package conf

// FinalMask is stubbed out in minimal build (feature removed)
type FinalMask struct {
	Tcp  []interface{} `json:"tcp"`
	Udp  []interface{} `json:"udp"`
	QuicParams *struct {
		BbrProfile                      string      `json:"bbrProfile"`
		Congestion                      string      `json:"congestion"`
		BrutalUp                        interface{} `json:"brutalUp"`
		BrutalDown                      interface{} `json:"brutalDown"`
		BrutalDisableLossCompensation   bool        `json:"brutalDisableLossCompensation"`
		InitStreamReceiveWindow         uint64      `json:"initStreamReceiveWindow"`
		MaxStreamReceiveWindow          uint64      `json:"maxStreamReceiveWindow"`
		InitConnectionReceiveWindow     uint64      `json:"initConnReceiveWindow"`
		MaxConnectionReceiveWindow      uint64      `json:"maxConnReceiveWindow"`
		MaxIdleTimeout                  int64       `json:"maxIdleTimeout"`
		KeepAlivePeriod                 int64       `json:"keepAlivePeriod"`
		DisablePathMTUDiscovery         bool        `json:"disablePathMTUDiscovery"`
		DisableChromeParrot             bool        `json:"disableChromeParrot"`
		DisableGSO                      bool        `json:"disableGSO"`
		MaxIncomingStreams              int64       `json:"maxIncomingStreams"`
		DisableStatelessReset           bool        `json:"disableStatelessReset"`
		Debug                           bool        `json:"debug"`
	} `json:"quicParams"`
}
