package evaluator

import (
	"github.com/ath0m/monkey/ast"
	"github.com/ath0m/monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}
