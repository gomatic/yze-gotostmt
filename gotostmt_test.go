package gotostmt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/tools/go/analysis/analysistest"

	gotostmt "github.com/gomatic/yze-gotostmt"
)

func TestGotoStatementIsReported(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), gotostmt.Analyzer, "a")
}

func TestRegistrationIsWellFormed(t *testing.T) {
	assert.NoError(t, gotostmt.Registration.Validate())
	assert.Equal(t, "yze/gotostmt", gotostmt.Registration.RuleID())
	assert.Same(t, gotostmt.Analyzer, gotostmt.Registration.Analyzer)
}
