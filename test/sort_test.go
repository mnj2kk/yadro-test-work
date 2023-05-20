package test

import (
	"YadroTestWork/test/utils"
	"testing"
)

func TestEmpty(t *testing.T) {
	utils.Run(t, "empty")
}

func TestLarge(t *testing.T) {
	utils.Run(t, "large")
}

func TestSample(t *testing.T) {
	utils.Run(t, "sample")
}
