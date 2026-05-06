package pluginsdk

import "io"

type TerminalTS interface {
	TerminalConnExists(terminalId string) bool
	TerminalGetPipe(agentId string, terminalId string) (*io.PipeReader, *io.PipeWriter, error)
	TerminalConnResume(agentId string, terminalId string, ioDirect bool)
	TerminalConnData(terminalId string, data []byte)
	TerminalConnClose(terminalId string, status string) error

	AgentTerminalCloseChannel(terminalId string, status string) error
}
