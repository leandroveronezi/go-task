package bin

import (
	"go/ast"
	"go/token"
)

var taskFunctions = []string{}
var targetFuncList = []string{}

var flagFileName *string
var flagSilent *bool
var flagKeepFile *bool
var flagViewGenSource *bool
var flagSortFunctions *bool
var flagContinueOnErrors *bool
var flagTargetFunc *string
var flagGroupFunc *string

var fset *token.FileSet
var nodeFile *ast.File

// Kind of task return
type KindType int

const (
	KBool  KindType = 1
	KError KindType = 2
)
