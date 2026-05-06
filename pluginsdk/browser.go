package pluginsdk

import PluModel "github.com/erniuu/w"

type BrowserTS interface {
	ClientGuiDisksWindows(taskData PluModel.TaskData, drives []PluModel.ListingDrivesDataWin)
	ClientGuiFilesStatus(taskData PluModel.TaskData)
	ClientGuiFilesWindows(taskData PluModel.TaskData, path string, files []PluModel.ListingFileDataWin)
	ClientGuiFilesUnix(taskData PluModel.TaskData, path string, files []PluModel.ListingFileDataUnix)
	ClientGuiProcessWindows(taskData PluModel.TaskData, process []PluModel.ListingProcessDataWin)
	ClientGuiProcessUnix(taskData PluModel.TaskData, process []PluModel.ListingProcessDataUnix)
}
