package main

import (
	"github.com/ysaito8015/calcFuncSize"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(calcFuncSize.Analyzer) }
