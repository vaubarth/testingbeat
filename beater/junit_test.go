package beater

import (
	"testing"
)

var expectedJunitSuite = JunitTestSuite{
	Name:     "org.tests.TestCase",
	Duration: 88.386,
	StdErr:   "sys err",
	StdOut:   "sys out",
	TestCases: []JunitTestcase{
		{
			Classname: "org.tests.TestCase",
			Name:      "This test PASSED in 2.067 seconds",
			Duration:  2.067,
			Skipped:   nil,
			Failure:   JunitFailure{},
		},
		{
			Classname: "org.tests.TestCase",
			Name:      "This test FAILED in 4.974 seconds",
			Duration:  4.974,
			Skipped:   nil,
			Failure: JunitFailure{
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
			Classname: "org.tests.TestCase",
			Name:      "This test was SKIPPED in 1.115 seconds",
			Duration:  1.115,
			Skipped:   &struct{}{},
			Failure:   JunitFailure{},
		},
	},
}

var junitSuite = readJunitFile("../tests/ExampleTests.xml")

func TestReadJunitFileSuite(t *testing.T) {
	if junitSuite.Duration != expectedJunitSuite.Duration {
		t.Errorf("Duration not correct, got '%f' - expected '%f'", junitSuite.Duration, expectedJunitSuite.Duration)
	}
	if junitSuite.Name != expectedJunitSuite.Name {
		t.Errorf("Name not correct, got '%s' - expected '%s'", junitSuite.Name, expectedJunitSuite.Name)
	}
	if junitSuite.StdErr != expectedJunitSuite.StdErr {
		t.Errorf("StdErr not correct, got '%s' - expected '%s'", junitSuite.StdErr, expectedJunitSuite.StdErr)
	}
	if junitSuite.StdOut != expectedJunitSuite.StdOut {
		t.Errorf("StdOut not correct, got '%s' - expected '%s'", junitSuite.StdOut, expectedJunitSuite.StdOut)
	}
}

func TestReadJunitFileTests(t *testing.T) {
	for i, test := range junitSuite.TestCases {
		expectedTest := expectedJunitSuite.TestCases[i]
		if test.Name != expectedTest.Name {
			t.Errorf("Name not correct, got '%s' - expected '%s'", test.Name, expectedTest.Name)
		}
		if test.Duration != expectedTest.Duration {
			t.Errorf("Duration not correct, got '%f' - expected '%f'", test.Duration, expectedTest.Duration)
		}
		if test.Skipped != nil && expectedTest.Skipped == nil {
			t.Errorf("Was not skipped bit should be")
		}
		if test.Skipped == nil && expectedTest.Skipped != nil {
			t.Errorf("Was sipped but shouldn't be")
		}
		if test.Classname != expectedTest.Classname {
			t.Errorf("Classname not correct, got '%s' - expected '%s'", test.Classname, expectedTest.Classname)
		}
	}
}

func TestReadJunitFileFailure(t *testing.T) {
	for i, test := range junitSuite.TestCases {
		expectedTest := expectedJunitSuite.TestCases[i]
		if test.Failure.Type != expectedTest.Failure.Type {
			t.Errorf("Failure.Type not correct, got '%s' - expected '%s'", test.Failure.Type, expectedTest.Failure.Type)
		}
		if test.Failure.Message != expectedTest.Failure.Message {
			t.Errorf("Failure.Message not correct, got '%s' - expected '%s'", test.Failure.Message, expectedTest.Failure.Message)
		}
		if test.Failure.Body != expectedTest.Failure.Body {
			t.Errorf("Failure.Body not correct, got '%s' - expected '%s'", test.Failure.Body, expectedTest.Failure.Body)
		}
	}
}

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

func TestMakeJunitReportSuite(t *testing.T) {
	testResults := junitSuite.makeJunitReport()
	for i, test := range testResults {
		expectedTest := expectedTestResults[i]
		if test.Suite.Duration != expectedTest.Suite.Duration {
			t.Errorf("Suite.Duration not correct, got '%f' - expected '%f'", test.Suite.Duration, expectedTest.Suite.Duration)
		}
		if test.Suite.Name != expectedTest.Suite.Name {
			t.Errorf("Suite.Name not correct, got '%s' - expected '%s'", test.Suite.Name, expectedTest.Suite.Name)
		}
		if test.Suite.StdErr != expectedTest.Suite.StdErr {
			t.Errorf("Suite.StdErr not correct, got '%s' - expected '%s'", test.Suite.StdErr, expectedTest.Suite.StdErr)
		}
		if test.Suite.StdOut != expectedTest.Suite.StdOut {
			t.Errorf("Suite.StdOut not correct, got '%s' - expected '%s'", test.Suite.StdOut, expectedTest.Suite.StdOut)
		}
	}
}

func TestMakeJunitReportTests(t *testing.T) {
	testResults := junitSuite.makeJunitReport()
	for i, test := range testResults {
		expectedTest := expectedTestResults[i]
		if test.Name != expectedTest.Name {
			t.Errorf("Name not correct, got '%s' - expected '%s'", test.Name, expectedTest.Name)
		}
		if test.Duration != expectedTest.Duration {
			t.Errorf("Duration not correct, got '%f' - expected '%f'", test.Duration, expectedTest.Duration)
		}
		if test.Skipped != expectedTest.Skipped {
			t.Errorf("Skipped not correct, got '%t' - expected '%t'", test.Skipped, expectedTest.Skipped)
		}
		if test.Success != expectedTest.Success {
			t.Errorf("Success not correct, got '%t' - expected '%t'", test.Success, expectedTest.Success)
		}
		if test.Failed != expectedTest.Failed {
			t.Errorf("Failed not correct, got '%t' - expected '%t'", test.Failed, expectedTest.Failed)
		}
	}
}

func TestMakeJunitReportFailure(t *testing.T) {
	testResults := junitSuite.makeJunitReport()

	for i, test := range testResults {
		expectedTest := expectedTestResults[i]
		if test.Failure.Type != expectedTest.Failure.Type {
			t.Errorf("Failure.Type not correct, got '%s' - expected '%s'", test.Failure.Type, expectedTest.Failure.Type)
		}
		if test.Failure.Title != expectedTest.Failure.Title {
			t.Errorf("Failure.Message not correct, got '%s' - expected '%s'", test.Failure.Title, expectedTest.Failure.Title)
		}
		if test.Failure.Body != expectedTest.Failure.Body {
			t.Errorf("Failure.Body not correct, got '%s' - expected '%s'", test.Failure.Body, expectedTest.Failure.Body)
		}
	}
}
