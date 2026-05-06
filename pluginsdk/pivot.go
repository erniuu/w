package pluginsdk

type PivotTS interface {
	PivotCreate(pivotId string, pAgentId string, chAgentId string, pivotName string, isRestore bool) error
	GetPivotInfoByName(pivotName string) (string, string, string)
	GetPivotInfoById(pivotId string) (string, string, string)
	PivotDelete(pivotId string) error
}
