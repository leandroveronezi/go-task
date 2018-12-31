package bin

var taskFunctions = []string{}
var targetFuncList = []string{}

// Kind of task return
type KindType int

const (
	KBool  KindType = 1
	KError KindType = 2
)
