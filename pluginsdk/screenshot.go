package pluginsdk

type ScreenShotTS interface {
	ScreenshotAdd(agentId string, note string, content []byte) error
}
