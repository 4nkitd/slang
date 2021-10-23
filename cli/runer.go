package cli

import (
	"flag"
	"fmt"
	"os"
	"slang/evaluator"
	"slang/lexer"
	"slang/object"
	"slang/parser"
	"time"
)

var engine = flag.String("file", "{FILE_NAME}", "use 'FILE_NAME' to run a file.")

func Run() bool {

	flag.Parse()

	if *engine == "{FILE_NAME}" {
		return false
	}

	dat, _ := os.ReadFile(*engine)

	input := string(dat)

	fmt.Println(input)

	var duration time.Duration
	var result object.Object

	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	env := object.NewEnvironment()
	start := time.Now()
	result = evaluator.Eval(program, env)
	duration = time.Since(start)

	fmt.Printf(
		"duration=%s\n",
		duration)

	fmt.Println("-== [ Result ] +++ \n" + result.Inspect())

	return true

}
