package beater

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"

	"github.com/vaubarth/testingbeat/config"
)

// Testingbeat configuration.
type Testingbeat struct {
	done   chan struct{}
	config config.Config
	client beat.Client
}

var logger = logp.NewLogger("testingbeat")

// New creates an instance of testingbeat.
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	c := config.DefaultConfig
	if err := cfg.Unpack(&c); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	bt := &Testingbeat{
		done:   make(chan struct{}),
		config: c,
	}
	return bt, nil
}

// Run starts testingbeat.
func (bt *Testingbeat) Run(b *beat.Beat) error {
	logger.Info("testingbeat is running! Hit CTRL-C to stop it.")

	var err error
	bt.client, err = b.Publisher.Connect()
	if err != nil {
		return err
	}

	// Create a fsnotify watcher for the configured path
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()
	if err := watcher.Add(bt.config.Path); err != nil {
		return err
	}

	for {
		select {
		case <-bt.done:
			return nil
		case err := <-watcher.Errors:
			logger.Error(err)
		case event := <-watcher.Events:
			// We only care if a file was created
			if event.Op == fsnotify.Create {
				logger.Info("New file was created: ", event.Name)
				testResult, err := bt.getResults(event)
				if err != nil {
					return err
				}
				runConfig, err := config.GetTestRunConfig(bt.config.RunConfigFile)
				if err != nil {
					return err
				}
				events := bt.resultToEvents(testResult, runConfig)
				for _, event := range events {
					bt.client.Publish(event)
				}
				logger.Info("Event send")
			}
		}

	}
}

// Selects the reporter type on based configured value
func (bt *Testingbeat) getResults(event fsnotify.Event) ([]TestResult, error) {
	switch bt.config.Type {
	// Only junit is available right now
	case "junit":
		return readJunitFile(event.Name).makeJunitReport(), nil
	default:
		return nil, fmt.Errorf("configured type '%s' is not supported", bt.config.Type)
	}
}

// Creates a list of events from TestResults
func (bt *Testingbeat) resultToEvents(result []TestResult, runConfig config.TestRunConfig) []beat.Event {
	var events []beat.Event
	for _, value := range result {
		events = append(events, beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"test": common.MapStr{
					"name":      value.Name,
					"duration":  value.Duration,
					"classname": value.Classname,
					"skipped":   value.Skipped,
					"failed":    value.Failed,
					"success":   value.Success,
					"state":     value.State,
					"metadata": common.MapStr{
						"runid":       runConfig.RunId,
						"environment": runConfig.Environment,
						"project":     runConfig.Project,
						"runner":      runConfig.Runner,
						"owner":       runConfig.Owner,
						"startedBy":   runConfig.StartedBy,
						"link":        runConfig.Link,
					},
					"suite": common.MapStr{
						"duration": value.Suite.Duration,
						"name":     value.Suite.Name,
						"stderr":   value.Suite.StdErr,
						"stdout":   value.Suite.StdOut,
					},
					"failure": common.MapStr{
						"type":  value.Failure.Type,
						"title": value.Failure.Title,
						"body":  value.Failure.Body,
					},
				},
			},
		})
	}
	return events
}

// Stop stops testingbeat.
func (bt *Testingbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
