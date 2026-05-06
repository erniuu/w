package pluginsdk

type DownLoadTS interface {
	DownloadAdd(agentId string, fileId string, fileName string, fileSize int) error
	DownloadUpdate(fileId string, state int, data []byte) error
	DownloadClose(fileId string, reason int) error
	DownloadSave(agentId string, fileId string, filename string, content []byte) error
}
