package examples

import (
	"github.com/gookit/ini/parser"
	"fmt"
)

func main()  {
	iniStr := `
# comments
name = inhere
age = 28
debug = true
hasQuota1 = 'this is val'
hasQuota2 = "this is val1"
shell = ${SHELL}
noEnv = ${NotExist|defValue}

; array in def section
tags[] = a
tags[] = b
tags[] = c

; comments
[sec1]
key = val0
some = value
stuff = things
; array in section
types[] = x
types[] = y
`
	// p, err := parser.Parse(iniStr, parser.FullMode)
	p, err := parser.Parse(iniStr, parser.FullMode, parser.NoDefSection)
	if err != nil {
		panic(err)
	}

	fmt.Printf("full parse:\n%#v\n", p.FullData())
}
