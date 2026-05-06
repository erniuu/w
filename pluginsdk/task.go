package pluginsdk

import PluModel "github.com/erniuu/w"

type TaskTS interface {
	TaskCreate(agentId string, cmdline string, client string, taskData PluModel.TaskData)
	TaskUpdate(agentId string, data PluModel.TaskData)
	TaskGetAvailableAll(agentId string, availableSize int) ([]PluModel.TaskData, error)
	TaskRunningExists(agentId string, taskId string) bool
}
