package ngrok

type Ngrok struct {
	Tunnels []Tunnel `json:"tunnels"`
}

type TunnelConfig struct {
	Address string `json:"addr"`
}

type Tunnel struct {
	PublicUrl string       `json:"public_url"`
	Protocol  string       `json:"proto"`
	Config    TunnelConfig `json:"config"`
}
