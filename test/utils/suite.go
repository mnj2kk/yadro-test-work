package utils

import (
	"YadroTestWork/internal/utils/functions/sorting"
	"YadroTestWork/test/resources"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func getJson(t *testing.T, name string) (data resources.Data, expected resources.Expected) {
	dataFile, dataErr := os.ReadFile(fmt.Sprintf("resources/data/%s.json", name))
	assert.NoErrorf(t, dataErr, "Can't read file resources/data/%s.json", name)
	assert.NoErrorf(t, json.Unmarshal(dataFile, &data), "Invalid json of resources/data/%s.json", name)

	expectedFile, expectedErr := os.ReadFile(fmt.Sprintf("resources/expected/%s.json", name))
	assert.NoErrorf(t, expectedErr, "Can't read file resources/expected/%s.json", name)
	assert.NoErrorf(t, json.Unmarshal(expectedFile, &expected), "Invalid json of resources/expected/%s.json", name)
	return
}

func Run(t *testing.T, name string) {
	data, expected := getJson(t, name)
	result := sorting.Sort(data.Names)
	assert.Equal(t, len(result), expected.Count, "Invalid count of unique names.")

	for _, kv := range result {
		val, ok := expected.Valid[kv.First]

		assert.Truef(t, ok, "Incorrect name %s in result.", kv.First)
		assert.Equalf(t, val, kv.Second, "The count of name %s is not equal.", kv.First)
	}
}
