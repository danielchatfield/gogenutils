package gogenutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldName(t *testing.T) {
	var testcases = []struct {
		In  string
		Out string
	}{
		{
			In:  "this is a test",
			Out: "ThisIsATest",
		},
		{
			In:  "this is a/test",
			Out: "ThisIsATest",
		},
	}

	for _, testcase := range testcases {
		name := testcase.In

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testcase.Out, FieldName(testcase.In))
		})
	}
}

func TestJSONFieldName(t *testing.T) {
	var testcases = []struct {
		In  string
		Out string
	}{
		{
			In:  "this is a test",
			Out: "this_is_a_test",
		},
		{
			In:  "this is a/test",
			Out: "this_is_a_test",
		},
		{
			In:  "this is  a/test",
			Out: "this_is_a_test",
		},
		{
			In:  "this is a test ",
			Out: "this_is_a_test",
		},
	}

	for _, testcase := range testcases {
		name := testcase.In

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testcase.Out, JSONFieldName(testcase.In))
		})
	}
}

func TestPascalCase(t *testing.T) {
	var testcases = []struct {
		In  string
		Out string
	}{
		{
			In:  "this is a test",
			Out: "ThisIsATest",
		},
	}

	for _, testcase := range testcases {
		name := testcase.In

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testcase.Out, PascalCase(testcase.In))
		})
	}
}

func TestSnakeCase(t *testing.T) {
	var testcases = []struct {
		In  string
		Out string
	}{
		{
			In:  "this is a test",
			Out: "this_is_a_test",
		},
		{
			In:  "This is a test",
			Out: "this_is_a_test",
		},
	}

	for _, testcase := range testcases {
		name := testcase.In

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testcase.Out, SnakeCase(testcase.In))
		})
	}
}

func TestPascalCaseToSnakeCase(t *testing.T) {
	var testcases = []struct {
		In  string
		Out string
	}{
		{
			In:  "ThisIsATest",
			Out: "this_is_a_test",
		},
		{
			In:  "HTTPBuffer",
			Out: "http_buffer",
		},
		{
			In:  "DE002PrimaryAccountNumber",
			Out: "de_002_primary_account_number",
		},
		{
			In:  "FABScott",
			Out: "fab_scott",
		},
	}

	for _, testcase := range testcases {
		name := testcase.In

		t.Run(name, func(t *testing.T) {
			assert.Equal(t, testcase.Out, PascalCaseToSnakeCase(testcase.In))
		})
	}
}
