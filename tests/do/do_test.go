package do

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Do(t *testing.T) {
	testTable := map[string]struct {
		someString     string
		someInt        int
		someBool       bool
		expectedResult string
		expectedError  string
	}{
		"valid string and int, false": {
			someString:     "a",
			someInt:        1,
			someBool:       false,
			expectedResult: "[1a]",
			expectedError:  "",
		},
		"valid string, ignore int, true": {
			someString:     "b",
			someInt:        13,
			someBool:       true,
			expectedResult: "B",
			expectedError:  "",
		},
		"invalid string, valid int, true": {
			someString:     "e",
			someInt:        1,
			someBool:       true,
			expectedResult: "",
			expectedError:  "invalid s",
		},
		"valid string, invalid int, true": {
			someString:     "a",
			someInt:        4,
			someBool:       true,
			expectedResult: "",
			expectedError:  "invalid s",
		},
	}

	for name, testCase := range testTable {
		t.Run(name, func(tt *testing.T) {
			result, err := Do(testCase.someString, testCase.someInt, testCase.someBool)
			t.Logf("Calling Do(%s, %d, %v), result: \"%s\", error: \"%s\"\n",
				testCase.someString, testCase.someInt, testCase.someBool, result, err)
			if err != nil {
				assert.EqualError(tt, err, testCase.expectedError, fmt.Sprintf("Incorrect error. Expected: %v, got %v", testCase.expectedError, err))
			}
			assert.Equal(tt, testCase.expectedResult, result, fmt.Sprintf("Incorrect result. Expected: %s, got %s", testCase.expectedResult, result))
		})
	}
}
