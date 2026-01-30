package model

import "time"

type Agent struct {
	Crc          string `json:"crc" gorm:"comment:Agent唯一校验码/指纹"`
	Id           string `json:"id" gorm:"comment:Agent ID"`
	Name         string `json:"name" gorm:"comment:Agent名称/别名"`
	SessionKey   []byte `json:"session_key" gorm:"comment:会话密钥(Session Key)"`
	Listener     string `json:"listener" gorm:"comment:对应监听器名称"`
	Async        bool   `json:"async" gorm:"comment:是否异步通信"`
	ExternalIP   string `json:"external_ip" gorm:"comment:外网IP"`
	InternalIP   string `json:"internal_ip" gorm:"comment:内网IP"`
	GmtOffset    int    `json:"gmt_offset" gorm:"comment:与GMT的时区偏移"`
	Sleep        uint   `json:"sleep" gorm:"comment:Sleep时间(秒)"`
	Jitter       uint   `json:"jitter" gorm:"comment:Jitter范围百分比"`
	Pid          string `json:"pid" gorm:"comment:进程PID"`
	Tid          string `json:"tid" gorm:"comment:线程TID"`
	Arch         string `json:"arch" gorm:"comment:系统架构(x86/x64/arm等)"`
	Elevated     bool   `json:"elevated" gorm:"comment:是否高权限/Elevated"`
	Process      string `json:"process" gorm:"comment:当前进程名"`
	Os           int    `json:"os" gorm:"comment:操作系统类型编号"`
	OsDesc       string `json:"os_desc" gorm:"comment:操作系统描述"`
	Domain       string `json:"domain" gorm:"comment:域名/工作组"`
	Computer     string `json:"computer" gorm:"comment:计算机名"`
	Username     string `json:"username" gorm:"comment:当前用户名"`
	Impersonated string `json:"impersonated" gorm:"comment:是否伪装令牌/Impersonated用户"`
	OemCP        int    `json:"oemcp" gorm:"comment:OEM代码页"`
	ACP          int    `json:"acp" gorm:"comment:ANSI代码页"`
	LastTick     int    `json:"last_tick" gorm:"comment:最后心跳时间戳"`
	KillDate     int    `json:"killdate" gorm:"comment:自动失效时间(Kill Date)"`
	WorkingTime  int    `json:"workingtime" gorm:"comment:运行时长秒数"`
	Tags         string `json:"tags" gorm:"comment:标签(管理员自定义)"`
	Mark         string `json:"mark" gorm:"comment:备注信息"`
	Color        string `json:"color" gorm:"comment:颜色标记(界面展示用)"`

	CreatedAt time.Time `gorm:"comment:创建时间"`
	UpdatedAt time.Time `gorm:"comment:更新时间"`
}
