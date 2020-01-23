package main

// A Var identifies a variable, e.g., x.
type Var string

// A literal is a numeric constant, e.g., 3.141.
type literal float64

type Expr interface{
	// Eval returns the value of this Expr in the environment env.
	Eval(Env Env) float64
}

type Env map[Var]float64

type unary struct {
	op rune // one of '+', '-'
	x Expr
}

type binary struct {
	op rune // one of '+', '-', '*', '/'
	x, y Expr
}

type call struct {
	fn string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (v Var) Eval(env Env) float64 {
	return env[v]
}