package pluginsdk

type PluginTS interface {
	ListenerReg(listenerInfo ListenerInfo) error
	ListenerTypeByName(listenerName string) (string, error)
	AgentReg(agentInfo AgentInfo) error

	ListenerUnreg(listenerName string) error
	AgentUnreg(agentName string) error
}
