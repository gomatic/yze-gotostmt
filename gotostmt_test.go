package gotostmt_test

import (
	"testing"

	gotostmt "github.com/gomatic/yze-go-gotostmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestGotoStatementIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), gotostmt.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, gotostmt.Registration.Validate())
	assert.Equal(t, "yze/go/gotostmt", gotostmt.Registration.RuleID())
	assert.Same(t, gotostmt.Analyzer, gotostmt.Registration.Analyzer)
}
