package test

import (
	"YadroTestWork/test/utils"
	"testing"
)

func TestEmpty(t *testing.T) {
	utils.Run(t, "empty")
}

func TestSample(t *testing.T) {
	utils.Run(t, "sample")
}

func TestLarge(t *testing.T) {
	utils.Run(t, "large")
}
