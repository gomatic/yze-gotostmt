// Package gotostmt provides a go/analysis analyzer that forbids the goto
// statement, per the gomatic Go standard that control flow uses early returns and
// extracted helpers rather than goto.
package gotostmt

import (
	"go/ast"
	"go/token"

	goyze "github.com/gomatic/go-yze"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const message = "goto is not permitted; restructure with early returns or helper functions"

// Analyzer reports every use of the goto statement.
var Analyzer = &analysis.Analyzer{
	Name:     "gotostmt",
	Doc:      "reports use of the goto statement, which the gomatic Go standard forbids",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

// Registration declares this analyzer to the yze framework.
var Registration = goyze.Registration{
	Name:       "gotostmt",
	Categories: []goyze.Category{"patterns"},
	URL:        "https://docs.gomatic.dev/yze/gotostmt",
	Analyzer:   Analyzer,
}

// run reports every goto statement in the analyzed package.
func run(pass *analysis.Pass) (any, error) {
	insp := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	insp.Preorder([]ast.Node{(*ast.BranchStmt)(nil)}, func(n ast.Node) {
		if stmt := n.(*ast.BranchStmt); stmt.Tok == token.GOTO {
			pass.Reportf(stmt.Pos(), message)
		}
	})
	return nil, nil
}
