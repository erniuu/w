package pluginsdk

type ListenerInfo struct {
	Name     string
	Protocol string
	Type     string
	UI       string
}

type AgentInfo struct {
	Name           string
	Watermark      string
	ListenersJson  string
	MenusJson      string
	CommandsJson   string
	MultiListeners bool
}
