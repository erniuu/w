package pluginsdk

import "io"

type ScreenTS interface {
	ScreenConnExists(screenId string) bool
	ScreenGetPipe(agentId string, screenId string) (*io.PipeReader, *io.PipeWriter, error)
	ScreenConnResume(agentId string, screenId string, ioDirect bool)
	ScreenConnData(screenId string, data []byte)
	AgentScreenCloseChannel(screenId string, status string) error
	ScreenConnClose(screenId string, status string) error
}
