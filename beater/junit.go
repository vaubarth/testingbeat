package beater

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type JunitTestSuite struct {
	XMLName   xml.Name        `xml:"testsuite"`
	Name      string          `xml:"name,attr"`
	Duration  float64         `xml:"time,attr"`
	StdErr    string          `xml:"system-err"`
	StdOut    string          `xml:"system-out"`
	TestCases []JunitTestcase `xml:"testcase"`
}

type JunitTestcase struct {
	XMLName   xml.Name     `xml:"testcase"`
	Classname string       `xml:"classname,attr"`
	Name      string       `xml:"name,attr"`
	Duration  float64      `xml:"time,attr"`
	Skipped   *struct{}    `xml:"skipped"`
	Failure   JunitFailure `xml:"failure"`
}

type JunitFailure struct {
	XMLName xml.Name `xml:"failure"`
	Message string   `xml:"message,attr"`
	Type    string   `xml:"type,attr"`
	Body    string   `xml:",chardata"`
}

// Reads a Junit xml result file and serializes it into a JunitTestSuite
func readJunitFile(fileName string) JunitTestSuite {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)
	var testSuite JunitTestSuite
	xmlerr := xml.Unmarshal(byteValue, &testSuite)
	if xmlerr != nil {
		fmt.Println(xmlerr)
	}

	return testSuite
}

// Creates a testingbeat TestResult from a JunitTestSuite
func (junitResult JunitTestSuite) makeJunitReport() []TestResult {
	var results []TestResult

	for i, test := range junitResult.TestCases {
		results = append(results, TestResult{
			Suite: Suite{
				Name:     junitResult.Name,
				Duration: junitResult.Duration,
			},
			Name:      test.Name,
			Classname: test.Classname,
			Duration:  test.Duration,
			Skipped:   false,
			Failed:    false,
			Success:   false,
			StdOut:    junitResult.StdOut,
			StdErr:    junitResult.StdErr,
			Failure: Failure{
				Title: test.Failure.Message,
				Type:  test.Failure.Type,
				Body:  test.Failure.Body,
			},
		})
		// Set test state boolean values
		if test.Skipped != nil {
			results[i].Skipped = true
		} else if test.Failure.Type != "" {
			results[i].Failed = true
		} else {
			results[i].Success = true
		}
	}

	return results

}
