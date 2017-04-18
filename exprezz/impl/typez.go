package exprezz

import (
	"fmt"
	"math"
)

//-------------------------------------------------------------
// A Var identifies a variable, e.g., z
type Var string
func (v Var) Eval(env Env) float64 {
	return env[v]
}

//-------------------------------------------------------------
// A literal is a numeric constant, like 3.141
type literal float64
func (l literal) Eval(_ Env) float64{
	return float64(l)
}

//-------------------------------------------------------------
// A unary represents a unary operator expression, like -x
type unary struct{
	op rune // one of '+', '-'
	x Expr
}

func (u unary) Eval(env Env) float64{
	switch u.op{
		case '+':
			return +u.x.Eval(env)
		case '-':
			return -u.x.Eval(env)
	}
	//otherwise pannic
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

//-------------------------------------------------------------
// a binary represents a binary operator expression, like x+y
type binary struct{
	op rune // + - / * 
	x,y Expr
}

func (b binary) Eval (env Env) float64{
	switch(b.op){
		case '+':
			return b.x.Eval(env) + b.y.Eval(env)
		case '-':
			return b.x.Eval(env) - b.y.Eval(env)
		case '*':
			return b.x.Eval(env) * b.y.Eval(env)
		case '/':
			return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unknown binary operator: %q", b.op))
}

//-------------------------------------------------------------
// A call represents a function call expression
type call struct{
	fn string // pow, sin, sqrt
	args []Expr
}

func (c call) Eval(env Env) float64{
	switch(c.fn){
		case "pow":
			return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
		case "sin":
			return math.Sin(c.args[0].Eval(env))
		case "sqrt":
			return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unknown fn: %q", c.fn))
}

//Env ...
type Env map[Var]float64 //to store variables

// Expr is an arithmetic expression
type Expr interface{
	Eval(env Env) float64
}


