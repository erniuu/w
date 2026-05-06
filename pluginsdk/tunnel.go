package pluginsdk

import "io"

type TunnelTS interface {
	TunnelChannelExists(channelId int) bool
	TunnelStart(tunnelId string) (string, error)
	TunnelCreateSocks4(agentId string, info string, lhost string, lport int) (string, error)
	TunnelCreateSocks5(agentId string, info string, lhost string, lport int, useAuth bool, username string, password string) (string, error)

	TunnelCreateRportfwd(agentId string, info string, lport int, thost string, tport int) (string, error)
	TunnelUpdateRportfwd(tunnelId int, result bool) (string, string, error)

	TunnelStopSocks(agentId string, port int)
	TunnelStopRportfwd(agentId string, port int)

	TunnelConnectionClose(channelId int, writeOnly bool)
	TunnelConnectionHalt(channelId int, errorCode byte)
	TunnelConnectionResume(agentId string, channelId int, ioDirect bool)
	TunnelConnectionData(channelId int, data []byte)
	TunnelConnectionAccept(tunnelId int, channelId int)
	TunnelPause(channelId int)
	TunnelResume(channelId int)

	TunnelGetPipe(agentId string, channelId int) (*io.PipeReader, *io.PipeWriter, error)
}
