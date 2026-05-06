package pluginsdk

type EncodingTS interface {
	ConvertCpToUTF8(input string, codePage int) string
	ConvertUTF8toCp(input string, codePage int) string
	Win32Error(errorCode uint) string
}
