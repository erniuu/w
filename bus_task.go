package business

import "time"

type Task struct {
	Type        int       `json:"t_type"         gorm:"comment:任务类型"`
	TaskId      string    `json:"t_task_id"      gorm:"primaryKey;size:64;comment:任务ID"`
	AgentId     string    `json:"t_agent_id"     gorm:"size:64;index;comment:关联Agent ID"`
	Client      string    `json:"t_client"       gorm:"size:64;comment:客户端来源"`
	HookId      string    `json:"t_hook_id"      gorm:"size:64;comment:Hook ID"`
	User        string    `json:"t_user"         gorm:"size:64;comment:执行用户"`
	Computer    string    `json:"t_computer"     gorm:"size:128;index;comment:主机名"`
	StartDate   int64     `json:"t_start_date"   gorm:"index;comment:任务开始时间"`
	FinishDate  int64     `json:"t_finish_date"  gorm:"index;comment:任务完成时间"`
	Data        []byte    `json:"t_data"         gorm:"type:longblob;comment:任务数据"`
	CommandLine string    `json:"t_command_line" gorm:"size:255;comment:命令行"`
	MessageType int       `json:"t_message_type" gorm:"comment:消息类型"`
	Message     string    `json:"t_message"      gorm:"size:500;comment:消息说明"`
	ClearText   string    `json:"t_clear_text"   gorm:"size:1000;comment:明文内容"`
	Completed   bool      `json:"t_completed"    gorm:"comment:是否完成"`
	Sync        bool      `json:"t_sync"         gorm:"comment:是否同步执行"`
	CreatedAt   time.Time `gorm:"comment:创建时间"`
	UpdatedAt   time.Time `gorm:"comment:更新时间"`
}

func (Task) TableName() string {
	return "bus_task"
}
