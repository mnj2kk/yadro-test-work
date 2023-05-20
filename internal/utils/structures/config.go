package structures

import (
	"YadroTestWork/internal/utils/functions/handler"
	"encoding/json"
	"strings"
)

type Config struct {
	Type   string  `json:"type"`
	Value  int     `json:"value"`
	Return bool    `json:"return"`
	Else   *Config `json:"else"`
}

func compare(a, b int) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func Parse(c Config) func(a, b Pair[string, int]) bool {
	if c.Else == nil {
		return func(a, b Pair[string, int]) bool {
			return c.Return
		}
	}
	next := Parse(*c.Else)
	if c.Type == "first" {
		return func(a, b Pair[string, int]) bool {
			if strings.Compare(a.First, b.First) == c.Value {
				return c.Return
			}
			return next(a, b)
		}
	}
	return func(a, b Pair[string, int]) bool {
		if compare(a.Second, b.Second) == c.Value {
			return c.Return
		}
		return next(a, b)
	}
}

func Build(data []byte) (config Config) {
	handler.Handle(json.Unmarshal(data, &config))
	return
}
