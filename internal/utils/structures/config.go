package structures

import "encoding/json"

type Config struct {
	Type   string  `json:"type"`
	Value  int     `json:"value"`
	Return bool    `json:"return"`
	Else   *Config `json:"else"`
}

func stringCompare(a, b string) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

func intCompare(a, b int) int {
	if a < b {
		return -1
	}
	if a == b {
		return 0
	}
	return 1
}

func (c Config) Parse() func(a, b Pair[string, int]) bool {
	if c.Else == nil {
		return func(a, b Pair[string, int]) bool {
			return c.Return
		}
	}
	next := c.Else.Parse()
	if c.Type == "second" {
		return func(a, b Pair[string, int]) bool {
			if intCompare(a.Second, b.Second) == c.Value {
				return c.Return
			}
			return next(a, b)
		}
	}
	return func(a, b Pair[string, int]) bool {
		if stringCompare(a.First, b.First) == c.Value {
			return c.Return
		}
		return next(a, b)
	}
}

func Build(data []byte) (config Config) {
	err := json.Unmarshal(data, &config)
	if err != nil {
		return Config{}
	}
	return
}
