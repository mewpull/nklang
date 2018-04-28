package ast

import "fmt"

type Object interface {
	Expression
	IsTrue() bool
}
type Function struct {
	Parameters []string
	Statements []Statement
}

func (n *Function) String() string {
	return fmt.Sprintf("Function{Parameters: %s, Statements: %s}", n.Parameters, n.Statements)
}

func (n *Function) Evaluate() (Object, error) {
	return n, nil
}

func (n *Function) IsTrue() bool {
	return true
}

type Integer struct {
	Value int64
}

func (n *Integer) String() string {
	return fmt.Sprintf("Integer{Value: %d}", n.Value)
}

func (n *Integer) Evaluate() (Object, error) {
	return n, nil
}

func (n *Integer) IsTrue() bool {
	return n.Value != 0
}

type String struct {
	Value string
}

func (n *String) String() string {
	return fmt.Sprintf("String{Value: %s}", n.Value)
}

func (n *String) Evaluate() (Object, error) {
	return n, nil
}

func (n *String) IsTrue() bool {
	return n.Value != ""
}

type Boolean struct {
	Value bool
}

func (n *Boolean) String() string {
	return fmt.Sprintf("Boolean{Value: %t}", n.Value)
}

func (n *Boolean) Evaluate() (Object, error) {
	return n, nil
}

func (n *Boolean) IsTrue() bool {
	return n.Value
}
