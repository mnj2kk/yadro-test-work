package test

import (
	"YadroTestWork/test/utils"
	"testing"
)

func TestEmpty(t *testing.T) {
	testCase := utils.WrapTest{Name: "empty"}
	testCase.Run(t)
}

func TestSample(t *testing.T) {
	testCase := utils.WrapTest{Name: "sample"}
	testCase.Run(t)
}

func TestLarge(t *testing.T) {
	testCase := utils.WrapTest{Name: "large"}
	testCase.Run(t)
}
