package model

import (
	"os"
)

const (
	OS_UNKNOWN = 0
	OS_WINDOWS = 1
	OS_LINUX   = 2
	OS_MAC     = 3

	TASK_TYPE_LOCAL      = 0
	TASK_TYPE_TASK       = 1
	TASK_TYPE_BROWSER    = 2
	TASK_TYPE_JOB        = 3
	TASK_TYPE_TUNNEL     = 4
	TASK_TYPE_PROXY_DATA = 5

	MESSAGE_INFO    = 5
	MESSAGE_ERROR   = 6
	MESSAGE_SUCCESS = 7

	BUILD_LOG_NONE    = 0
	BUILD_LOG_INFO    = 1
	BUILD_LOG_ERROR   = 2
	BUILD_LOG_SUCCESS = 3

	DOWNLOAD_STATE_RUNNING  = 1
	DOWNLOAD_STATE_STOPPED  = 2
	DOWNLOAD_STATE_FINISHED = 3
	DOWNLOAD_STATE_CANCELED = 4

	TUNNEL_TYPE_SOCKS4      = 1
	TUNNEL_TYPE_SOCKS5      = 2
	TUNNEL_TYPE_SOCKS5_AUTH = 3
	TUNNEL_TYPE_LOCAL_PORT  = 4
	TUNNEL_TYPE_REVERSE     = 5

	ADDRESS_TYPE_IPV4   = 1
	ADDRESS_TYPE_DOMAIN = 3
	ADDRESS_TYPE_IPV6   = 4

	SOCKS5_SUCCESS                 byte = 0
	SOCKS5_SERVER_FAILURE          byte = 1
	SOCKS5_NOT_ALLOWED_RULESET     byte = 2
	SOCKS5_NETWORK_UNREACHABLE     byte = 3
	SOCKS5_HOST_UNREACHABLE        byte = 4
	SOCKS5_CONNECTION_REFUSED      byte = 5
	SOCKS5_TTL_EXPIRED             byte = 6
	SOCKS5_COMMAND_NOT_SUPPORTED   byte = 7
	SOCKS5_ADDR_TYPE_NOT_SUPPORTED byte = 8
)

type PluginListener interface {
	Create(name, config string, customData []byte) (ExtenderListener, ListenerData, []byte, error)
}

type ExtenderListener interface {
	Start() error
	Edit(config string) (ListenerData, []byte, error)
	Stop() error
	GetProfile() ([]byte, error)
	InternalHandler(data []byte) (string, error)
}

type PluginAgent interface {
	GenerateProfiles(profile BuildProfile) ([][]byte, error)
	BuildPayload(profile BuildProfile, agentProfiles [][]byte) ([]byte, string, error)

	GetExtender() ExtenderAgent
	CreateAgent(beat []byte) (AgentData, ExtenderAgent, error)
}

type ExtenderAgent interface {
	Encrypt(data []byte, key []byte) ([]byte, error)
	Decrypt(data []byte, key []byte) ([]byte, error)

	PackTasks(agentData AgentData, tasks []TaskData) ([]byte, error)
	PivotPackData(pivotId string, data []byte) (TaskData, error)

	CreateCommand(agentData AgentData, args map[string]any) (TaskData, ConsoleMessageData, error)
	ProcessData(agentData AgentData, decryptedData []byte) error

	TunnelCallbacks() TunnelCallbacks
	TerminalCallbacks() TerminalCallbacks
}

type TunnelCallbacks struct {
	ConnectTCP func(channelId, tunnelType, addressType int, address string, port int) TaskData
	ConnectUDP func(channelId, tunnelType, addressType int, address string, port int) TaskData
	WriteTCP   func(channelId int, data []byte) TaskData
	WriteUDP   func(channelId int, data []byte) TaskData
	Pause      func(channelId int) TaskData
	Resume     func(channelId int) TaskData
	Close      func(channelId int) TaskData
	Reverse    func(tunnelId, port int) TaskData
}

type TerminalCallbacks struct {
	Start func(terminalId int, program string, sizeH, sizeW, oemCP int) TaskData
	Write func(terminalId, oemCP int, data []byte) TaskData
	Close func(terminalId int) TaskData
}

type ListenerData struct {
	Name       string `json:"name"`
	RegName    string `json:"reg_name"`
	Protocol   string `json:"protocol"`
	Type       string `json:"type"`
	BindHost   string `json:"bind_host"`
	BindPort   string `json:"bind_port"`
	AgentAddr  string `json:"agent_addr"`
	CreateTime int64  `json:"create_time"`
	Status     string `json:"status"`
	Data       string `json:"data"`
	Watermark  string `json:"watermark"`
}

type AgentData struct {
	Crc          string `json:"crc"`
	Id           string `json:"id"`
	Name         string `json:"name"`
	SessionKey   []byte `json:"session_key"`
	Listener     string `json:"listener"`
	Async        bool   `json:"async"`
	ExternalIP   string `json:"external_ip"`
	InternalIP   string `json:"internal_ip"`
	GmtOffset    int    `json:"gmt_offset"`
	Sleep        uint   `json:"sleep"`
	Jitter       uint   `json:"jitter"`
	Pid          string `json:"pid"`
	Tid          string `json:"tid"`
	Arch         string `json:"arch"`
	Elevated     bool   `json:"elevated"`
	Process      string `json:"process"`
	Os           int    `json:"os"`
	OsDesc       string `json:"os_desc"`
	Domain       string `json:"domain"`
	Computer     string `json:"computer"`
	Username     string `json:"username"`
	Impersonated string `json:"impersonated"`
	OemCP        int    `json:"oemcp"`
	ACP          int    `json:"acp"`
	CreateTime   int64  `json:"create_time"`
	LastTick     int    `json:"last_tick"`
	KillDate     int    `json:"killdate"`
	WorkingTime  int    `json:"workingtime"`
	Tags         string `json:"tags"`
	Mark         string `json:"mark"`
	Color        string `json:"color"`
	TargetId     string `json:"target"`
	CustomData   []byte `json:"custom_data"`
}

type TaskDataTunnel struct {
	ChannelId int
	Data      TaskData
}

type TaskData struct {
	Type        int    `json:"type"`
	TaskId      string `json:"task_id"`
	AgentId     string `json:"agent_id"`
	Client      string `json:"client"`
	HookId      string `json:"hook_id"`
	HandlerId   string `json:"handler_id"`
	User        string `json:"user"`
	Computer    string `json:"computer"`
	StartDate   int64  `json:"start_date"`
	FinishDate  int64  `json:"finish_date"`
	Data        []byte `json:"data"`
	CommandLine string `json:"command_line"`
	MessageType int    `json:"message_type"`
	Message     string `json:"message"`
	ClearText   string `json:"clear_text"`
	Completed   bool   `json:"completed"`
	Sync        bool   `json:"sync"`
}

type ConsoleMessageData struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Text    string `json:"text"`
}

type ListingFileDataWin struct {
	IsDir    bool   `json:"is_dir"`
	Size     int64  `json:"size"`
	Date     int64  `json:"date"`
	Filename string `json:"filename"`
}

type ListingFileDataUnix struct {
	IsDir    bool   `json:"is_dir"`
	Mode     string `json:"mode"`
	User     string `json:"user"`
	Group    string `json:"group"`
	Size     int64  `json:"size"`
	Date     string `json:"date"`
	Filename string `json:"filename"`
}

type ListingProcessDataWin struct {
	Pid         uint   `json:"pid"`
	Ppid        uint   `json:"ppid"`
	SessionId   uint   `json:"session_id"`
	Arch        string `json:"arch"`
	Context     string `json:"context"`
	ProcessName string `json:"process_name"`
}

type ListingProcessDataUnix struct {
	Pid         uint   `json:"pid"`
	Ppid        uint   `json:"ppid"`
	TTY         string `json:"tty"`
	Context     string `json:"context"`
	ProcessName string `json:"process_name"`
}

type ListingDrivesDataWin struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ChatData struct {
	Username string `json:"username"`
	Message  string `json:"message"`
	Date     int64  `json:"date"`
}

type DownloadData struct {
	FileId     string `json:"file_id"`
	AgentId    string `json:"agent_id"`
	AgentName  string `json:"agent_name"`
	User       string `json:"user"`
	Computer   string `json:"computer"`
	RemotePath string `json:"remote_path"`
	LocalPath  string `json:"local_path"`
	TotalSize  int64  `json:"total_size"`
	RecvSize   int64  `json:"recv_size"`
	Date       int64  `json:"date"`
	State      int    `json:"state"`
	File       *os.File
}

type ScreenData struct {
	ScreenId  string `json:"screen_id"`
	AgentId     string `json:"agent_id"`
	User      string `json:"user"`
	Computer  string `json:"computer"`
	LocalPath string `json:"local_path"`
	Note      string `json:"note"`
	Date      int64  `json:"date"`
	Content   []byte `json:"content"`
}

type TunnelData struct {
	TunnelId  string `json:"tunnel_id"`
	AgentId   string `json:"agent_id"`
	Computer  string `json:"computer"`
	Username  string `json:"username"`
	Process   string `json:"process"`
	Type      string `json:"type"`
	Info      string `json:"info"`
	Interface string `json:"interface"`
	Port      string `json:"port"`
	Client    string `json:"client"`
	Fhost     string `json:"fhost"`
	Fport     string `json:"fport"`
	AuthUser  string `json:"auth_user"`
	AuthPass  string `json:"auth_pass"`
}

type PivotData struct {
	PivotId       string `json:"pivot_id"`
	PivotName     string `json:"pivot_name"`
	ParentAgentId string `json:"parent_agent_id"`
	ChildAgentId  string `json:"child_agent_id"`
}

type CredsData struct {
	CredId   string `json:"creds_id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Realm    string `json:"realm"`
	Type     string `json:"type"`
	Tag      string `json:"tag"`
	Date     int64  `json:"date"`
	Storage  string `json:"storage"`
	AgentId  string `json:"agent_id"`
	Host     string `json:"host"`
}

type TargetData struct {
	TargetId string   `json:"target_id"`
	Computer string   `json:"computer"`
	Domain   string   `json:"domain"`
	Address  string   `json:"address"`
	Os       int      `json:"os"`
	OsDesk   string   `json:"os_desk"`
	Tag      string   `json:"tag"`
	Info     string   `json:"info"`
	Date     int64    `json:"date"`
	Alive    bool     `json:"alive"`
	Agents   []string `json:"agents"`
}

type TransportProfile struct {
	Watermark string `json:"watermark"`
	Profile   []byte `json:"profile"`
}

type BuildProfile struct {
	BuilderId        string             `json:"build_id"`
	AgentConfig      string             `json:"agent_params"`
	ListenerProfiles []TransportProfile `json:"listener_profiles"`
}
