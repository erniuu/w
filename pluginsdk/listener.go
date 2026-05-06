package pluginsdk

type ListenerTS interface {
	ListenerInternalHandler(watermark string, data []byte) (string, error)
}
