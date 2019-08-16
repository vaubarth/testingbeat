package beater

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/vaubarth/testingbeat/config"
	"testing"

	td "github.com/maxatome/go-testdeep"
)

var testResults = []TestResult{
	{
		Name:      "testname",
		Classname: "classname",
		Duration:  3.5,
		Skipped:   true,
		Failed:    false,
		Success:   false,
		Suite: Suite{
			Name:     "suitename",
			Duration: 2.2,
			StdOut:   "stdout",
			StdErr:   "stderr",
		},
		Failure: Failure{
			Type:  "failtype",
			Title: "failtitle",
			Body:  "failbody",
		},
	},
}

var expectedFieldsMap = common.MapStr{
	"name":      "testname",
	"duration":  3.5,
	"classname": "classname",
	"skipped":   true,
	"failed":    false,
	"success":   false,
	"metadata": common.MapStr{
		"runid":       "runid",
		"environment": "env",
		"project":     "project",
		"link":        "link",
	},
	"suite": common.MapStr{
		"duration": 2.2,
		"name":     "suitename",
		"stderr":   "stderr",
		"stdout":   "stdout",
	},
	"failure": common.MapStr{
		"type":  "failtype",
		"title": "failtitle",
		"body":  "failbody",
	},
}

var bt = Testingbeat{
	done: nil,
	config: config.Config{
		Path:        "",
		Type:        "",
		RunId:       "runid",
		Environment: "env",
		Project:     "project",
		Link:        "link",
	},
	client: nil,
}

var events = bt.resultToEvents(testResults)

func TestResultToEvents(t *testing.T) {
	if len(events) != 1 {
		t.Errorf("Size of events not correct, got '%d' - expected '1'", len(events))
	}
}

func TestResultToEvents1(t *testing.T) {
	td.Cmp(t, events[0].Fields, expectedFieldsMap)
}
