package pluginsdk

import PluModel "github.com/erniuu/w"

type AgentTS interface {
	AgentIsExists(agentId string) bool
	AgentSetTick(agentId string, listenerName string) error
	AgentGetHostedAll(agentId string, maxDataSize int) ([]byte, error)
	AgentGetHostedTasks(agentId string, maxDataSize int) ([]byte, error)
	AgentCreate(agentCrc string, agentId string, beat []byte, listenerName string, externalIP string, async bool) (PluModel.AgentData, error)

	AgentProcessData(agentId string, bodyData []byte) error
	AgentUpdateData(newAgentData PluModel.AgentData) error
	AgentTerminate(agentId string, terminateTaskId string) error
	AgentUpdateDataPartial(agentId string, updateData interface{}) error

	AgentBuildExecute(builderId string, workingDir string, program string, args ...string) error
	AgentBuildLog(builderId string, status int, message string) error

	AgentConsoleOutput(agentId string, messageType int, message string, clearText string, store bool)
}
