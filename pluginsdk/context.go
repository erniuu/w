package pluginsdk

type PluginContext interface {
	Agent() AgentTS
	Browser() BrowserTS
	Encoding() EncodingTS
	DownLoad() DownLoadTS
	Listener() ListenerTS
	Pivot() PivotTS
	Plugin() PluginTS
	ScreenShot() ScreenShotTS
	Screen() ScreenTS
	Tunnel() TunnelTS
	Task() TaskTS
	Terminal() TerminalTS
}
