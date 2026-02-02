package model

type ListenerData struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	BindHost  string `json:"bind_host"`
	BindPort  string `json:"bind_port"`
	AgentAddr string `json:"agent_addr"`
	Status    string `json:"status"`
	Data      string `json:"data"`
	Watermark string `json:"watermark"`
}
