package types

type NodeModal struct {
	Id                int64   `json:"id"`
	Name              string  `json:"name"`
	Type              bool    `json:"type"`
	Server            string  `json:"server"`
	Info              string  `json:"info"`
	Status            string  `json:"status"`
	Sort              int     `json:"sort"`
	TrafficRate       float64 `json:"traffic_rate"`
	Class             int64   `json:"node_class"`
	SpeedLimit        int64   `json:"node_speedlimit"`
	Connector         int64   `json:"node_connector"`
	Bandwidth         int64   `json:"node_bandwidth"`
	BandwidthLimit    int64   `json:"node_bandwidth_limit"`
	BandwidthResetDay int     `json:"bandwidthlimit_resetday"`
	Heartbeat         int     `json:"node_heartbeat"`
	IP                string  `json:"node_ip"`
	Group             int64   `json:"node_group"`
	MuType            int     `json:"mu_only"`
	Online            int     `json:"online"`
	GFWBlock          int     `json:"gfw_block"`

	// String stored JSON
	CustomConfig string `json:"custom_config"`
}
