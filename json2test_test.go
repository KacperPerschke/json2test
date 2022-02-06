package main

import (
	"fmt"
	"testing"
)

type testCaseT struct {
	in  string
	out string
	err string
}

var testCases = []testCaseT{
	{
		in:  ``,
		out: ``,
		err: `unexpected end of JSON input`,
	},
	{
		in:  `[{"Time":"…","Action":"…","Package":"…","Test":"…"}]`,
		out: ``,
		err: `json: cannot unmarshal array into Go value of type struct { Time *time.Time "json:\",omitempty\""; Action string; Package string "json:\",omitempty\""; Test string "json:\",omitempty\""; Elapsed *float64 "json:\",omitempty\""; Output string "json:\",omitempty\"" }`,
	},
	{
		in:  `{"Time":"1989-01-04T12:00:00.575597712+01:00","Action":"run","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases"}`,
		out: ``,
	},
	{
		in:  `{"Time":"1989-01-04T12:00:00.575855378+01:00","Action":"output","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases/La_configuration.","Output":"=== RUN   TestAllCases/La_configuration.\n"}`,
		out: "=== RUN   TestAllCases/La_configuration.\n",
	},
	{
		in:  `{"Time":"1989-01-04T12:00:00.818412481+01:00","Action":"output","Package":"inteligo.com.pl/srv-template-go","Test":"TestAllCases/La_configuration.","Output":"\t--- PASS: TestAllCases/La_configuration. (0.17s)\n"}`,
		out: "\t--- PASS: TestAllCases/La_configuration. (0.17s)\n",
	},
}

func TestJson2Test(t *testing.T) {
	for _, testCase := range testCases {
		res, err := json2test([]byte(testCase.in))
		if err != nil {
			errWant := testCase.err
			errGot := fmt.Sprintf("%s", err)
			if errWant != errGot {
				t.Errorf("Expected err message `%s`, but got `%s`", errWant, errGot)
			}
		}
		out := testCase.out
		if res != out {
			t.Errorf("Expected `%s` but got `%s`", out, res)
		}
	}
}
