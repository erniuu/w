package pluginsdk

import "io"

type TunnelTS interface {
	TunnelChannelExists(channelId int) bool
	TunnelStart(tunnelId string) (string, error)
	
	TunnelCreateSocks4(agentId string, info string, lhost string, lport int) (string, error)
	TunnelCreateSocks5(agentId string, info string, lhost string, lport int, useAuth bool, username string, password string) (string, error)
	TunnelCreateLportfwd(agentId string, info string, lhost string, lport int, thost string, tport int) (string, error)
	TunnelCreateHttp(agentId string, info string, lhost string, lport int, useAuth bool, username string, password string) (string, error)

	TunnelStopSocks(agentId string, port int)
	TunnelStopLportfwd(agentId string, port int)

	TunnelConnectionClose(channelId int, writeOnly bool)
	TunnelConnectionHalt(channelId int, errorCode byte)
	TunnelConnectionResume(agentId string, channelId int, ioDirect bool)
	TunnelConnectionData(channelId int, data []byte)
	TunnelPause(channelId int)
	TunnelResume(channelId int)

	TunnelGetPipe(agentId string, channelId int) (*io.PipeReader, *io.PipeWriter, error)
}
