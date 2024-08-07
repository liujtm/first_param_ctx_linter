package main

import (
	"github.com/liujtm/first_param_ctx_linter/firstparamcontext"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(firstparamcontext.Analyzer)
}
