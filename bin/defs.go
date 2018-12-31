package bin

var taskFunctions = []string{}
var targetFuncList = []string{}

type KindType int

const (
	KBool  KindType = 1
	KError KindType = 2
)
