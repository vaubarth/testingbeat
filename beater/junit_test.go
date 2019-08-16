package beater

import (
	"encoding/xml"
	"testing"

	td "github.com/maxatome/go-testdeep"
)

var testCaseXmlName = xml.Name{
	Space: "",
	Local: "testcase",
}

var expectedJunitSuite = JunitTestSuite{
	Name:     "org.tests.TestCase",
	Duration: 88.386,
	StdErr:   "sys err",
	StdOut:   "sys out",
	XMLName: xml.Name{
		Space: "",
		Local: "testsuite",
	},
	TestCases: []JunitTestcase{
		{
			XMLName:   testCaseXmlName,
			Classname: "org.tests.TestCase",
			Name:      "This test PASSED in 2.067 seconds",
			Duration:  2.067,
			Skipped:   nil,
			Failure:   JunitFailure{},
		},
		{
			XMLName:   testCaseXmlName,
			Classname: "org.tests.TestCase",
			Name:      "This test FAILED in 4.974 seconds",
			Duration:  4.974,
			Skipped:   nil,
			Failure: JunitFailure{
				XMLName: xml.Name{
					Space: "",
					Local: "failure",
				},
				Message: "1 is not 2",
				Type:    "org.scalatest.exceptions.TestFailedException",
				Body: `
Body text
Probably stacktrace
...
...
...`,
			},
		},
		{
			XMLName:   testCaseXmlName,
			Classname: "org.tests.TestCase",
			Name:      "This test was SKIPPED in 1.115 seconds",
			Duration:  1.115,
			Skipped:   &struct{}{},
			Failure:   JunitFailure{},
		},
	},
}

var junitSuite = readJunitFile("../tests/ExampleTests.xml")

var expectedSuite = Suite{
	Name:     "org.tests.TestCase",
	Duration: 88.386,
	StdOut:   "sys out",
	StdErr:   "sys err",
}

var expectedTestResults = []TestResult{
	{
		Suite:     expectedSuite,
		Name:      "This test PASSED in 2.067 seconds",
		Classname: "org.tests.TestCase",
		Duration:  2.067,
		Skipped:   false,
		Failed:    false,
		Success:   true,
		Failure: Failure{
			Title: "",
			Type:  "",
			Body:  "",
		},
	},
	{
		Suite:     expectedSuite,
		Name:      "This test FAILED in 4.974 seconds",
		Classname: "org.tests.TestCase",
		Duration:  4.974,
		Skipped:   false,
		Failed:    true,
		Success:   false,
		Failure: Failure{
			Title: "1 is not 2",
			Type:  "org.scalatest.exceptions.TestFailedException",
			Body: `
Body text
Probably stacktrace
...
...
...`,
		},
	},
	{
		Suite:     expectedSuite,
		Name:      "This test was SKIPPED in 1.115 seconds",
		Classname: "org.tests.TestCase",
		Duration:  1.115,
		Skipped:   true,
		Failed:    false,
		Success:   false,
		Failure: Failure{
			Title: "",
			Type:  "",
			Body:  "",
		},
	},
}

func TestMakeJunitReport(t *testing.T) {
	td.Cmp(t, junitSuite.makeJunitReport(), expectedTestResults)
}

func TestReadJunitFile(t *testing.T) {
	td.Cmp(t, junitSuite, expectedJunitSuite)
}
