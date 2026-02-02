package model

type TaskData struct {
	Type        int       `json:"type"`
	TaskId      string    `json:"task_id"`
	AgentId     string    `json:"agent_id"`
	Client      string    `json:"client"`
	HookId      string    `json:"hook_id"`
	User        string    `json:"user"`
	Computer    string    `json:"computer"`
	StartDate   int64     `json:"start_date"`
	FinishDate  int64     `json:"finish_date"`
	Data        []byte    `json:"data"`
	CommandLine string    `json:"command_line"`
	MessageType int       `json:"message_type"`
	Message     string    `json:"message"`
	ClearText   string    `json:"clear_text"`
	Completed   bool      `json:"completed"`
	Sync        bool      `json:"sync"`
}
