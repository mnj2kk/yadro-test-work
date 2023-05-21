package test

import (
	"YadroTestWork/internal/utils/functions/handler"
	"YadroTestWork/internal/utils/functions/sorting"
	"YadroTestWork/internal/utils/structures"
	"YadroTestWork/test/resources"
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/seehuhn/mt19937"
	"github.com/stretchr/testify/suite"
	"math/rand"
	"os"
	"testing"
	"time"
)

var (
	rng     = rand.New(mt19937.New())
	configs = map[string]func(a, b structures.Pair[string, int]) bool{
		"desc-count-desc-names": func(a, b structures.Pair[string, int]) bool {
			if a.Second < b.Second {
				return true
			}
			if a.Second > b.Second {
				return false
			}
			return a.First < b.First
		},
		"asc-count-desc-names": func(a, b structures.Pair[string, int]) bool {
			if a.Second > b.Second {
				return true
			}
			if a.Second < b.Second {
				return false
			}
			return a.First < b.First
		},
	}
)

const (
	RandomTestNumber = 17
	Small            = 31
	Medium           = 1337
	Large            = 6661337
)

type TestSuite struct {
	suite.Suite
	names []string
}

func getMapResult(names []string) map[string]int {
	result := make(map[string]int, 0)
	for _, v := range names {
		_, ok := result[v]
		if !ok {
			result[v] = 0
		}
		result[v]++
	}
	return result
}

func dataRun(s *TestSuite, names []string) {
	actual := sorting.Sort(names, func(a, b structures.Pair[string, int]) bool {
		return a.First < b.First
	})
	expected := getMapResult(names)

	s.Equalf(len(expected), len(actual), "Expected %d unique names, but found %d", len(expected), len(actual))

	for _, kv := range actual {
		val, ok := expected[kv.First]

		s.Truef(ok, "Incorrect name %s in result.", kv.First)
		s.Equalf(val, kv.Second, "The count of name %s is not equal.", kv.First)
	}
}

func getJson(s *TestSuite, name string) (data resources.Data, expected resources.Expected) {
	dataFile, dataErr := os.ReadFile(fmt.Sprintf("resources/data/%s.json", name))
	s.NoErrorf(dataErr, "Can't read file resources/data/%s.json", name)
	s.NoErrorf(json.Unmarshal(dataFile, &data), "Invalid json of resources/data/%s.json", name)

	expectedFile, expectedErr := os.ReadFile(fmt.Sprintf("resources/expected/%s.json", name))
	s.NoErrorf(expectedErr, "Can't read file resources/expected/%s.json", name)
	s.NoErrorf(json.Unmarshal(expectedFile, &expected), "Invalid json of resources/expected/%s.json", name)
	return
}

func fileRun(s *TestSuite, name string) {
	data, expected := getJson(s, name)
	result := sorting.Sort(data.Names, func(a, b structures.Pair[string, int]) bool {
		return a.First < b.First
	})
	s.Equal(len(result), expected.Count, "Invalid count of unique names.")

	for _, kv := range result {
		val, ok := expected.Valid[kv.First]

		s.Truef(ok, "Incorrect name %s in result.", kv.First)
		s.Equalf(val, kv.Second, "The count of name %s is not equal.", kv.First)
	}
}

func getNames(suite *TestSuite, size int) (result []string) {
	for i := 0; i < size; i++ {
		result = append(result, suite.names[rng.Intn(len(suite.names))])
	}
	return
}

func randomCheck(suite *TestSuite, MAX int) {
	for i := 0; i < RandomTestNumber; i++ {
		names := getNames(suite, rng.Intn(MAX-1)+1)
		dataRun(suite, names)
	}
}

func configCheck(suite *TestSuite, MAX int, name string) {
	dataFile, jsonErr := os.ReadFile(fmt.Sprintf("resources/configs/%s.json", name))
	handler.Handle(jsonErr)

	compare := structures.Parse(structures.Build(dataFile))

	for i := 0; i < RandomTestNumber; i++ {
		names := getNames(suite, rng.Intn(MAX-1)+1)
		actual := sorting.Sort(names, configs[name])
		expected := sorting.Sort(names, compare)

		suite.Equalf(actual, expected, "Wrong config.")
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (suite *TestSuite) SetupSuite() {
	file, err := os.Open("resources/names.txt")
	handler.Handle(err)
	defer func(file *os.File) {
		handler.Handle(file.Close())
	}(file)

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		suite.names = append(suite.names, scanner.Text())
	}
}

func (suite *TestSuite) SetupTest() {
	rng.Seed(time.Now().UnixNano())
}

func (suite *TestSuite) TestEmpty() {
	fileRun(suite, "empty")
}

func (suite *TestSuite) TestSample() {
	fileRun(suite, "sample")
}

func (suite *TestSuite) TestLarge() {
	fileRun(suite, "large")
}

func (suite *TestSuite) TestSmallRandom() {
	randomCheck(suite, Small)
}

func (suite *TestSuite) TestMediumRandom() {
	randomCheck(suite, Medium)
}

func (suite *TestSuite) TestLargeRandom() {
	randomCheck(suite, Large)
}

func (suite *TestSuite) TestSmallFirstSortConfig() {
	configCheck(suite, Small, "desc-count-desc-names")
}

func (suite *TestSuite) TestSmallSecondSortConfig() {
	configCheck(suite, Small, "asc-count-desc-names")
}

func (suite *TestSuite) TestMediumFirstSortConfig() {
	configCheck(suite, Medium, "desc-count-desc-names")
}

func (suite *TestSuite) TestMediumSecondSortConfig() {
	configCheck(suite, Medium, "asc-count-desc-names")
}

func (suite *TestSuite) TestLargeFirstSortConfig() {
	configCheck(suite, Large, "desc-count-desc-names")
}

func (suite *TestSuite) TestLargeSecondSortConfig() {
	configCheck(suite, Large, "asc-count-desc-names")
}
