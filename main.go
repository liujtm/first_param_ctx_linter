package main

import (
	"firstparamcontextlinter/firstparamcontext"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(firstparamcontext.Analyzer)
}
