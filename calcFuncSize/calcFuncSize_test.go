package calcFuncSize_test

import (
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/ysaito8015/calcFuncSize"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, calcFuncSize.Analyzer, "a")
}
