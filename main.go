package main

import (
	"github.com/liujtm/firstparamcontextlinter/firstparamcontext"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(firstparamcontext.Analyzer)
}
