package beater

import (
	"errors"
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
		return nil, fmt.Errorf("Error reading config file: %v", err)
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
	logger.Info("RunId is: ", bt.config.RunId, "Environment is: ", bt.config.Environment)

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
				bt.send(b, testResult)
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
		return nil, errors.New("")
	}
}

func (bt *Testingbeat) send(b *beat.Beat, result []TestResult) {
	// Make a Map from a TestResult
	for _, value := range result {
		event := beat.Event{
			Timestamp: time.Now(),
			Fields: common.MapStr{
				"type":      b.Info.Name,
				"name":      value.Name,
				"duration":  value.Duration,
				"classname": value.Classname,
				"skipped":   value.Skipped,
				"failed":    value.Failed,
				"success":   value.Success,
				"stderr":    value.StdErr,
				"stdout":    value.StdOut,
				"metadata": common.MapStr{
					"runid":       bt.config.RunId,
					"environment": bt.config.Environment,
					"project":     bt.config.Project,
					"link":        bt.config.Link,
				},
				"suite": common.MapStr{
					"duration": value.Suite.Duration,
					"name":     value.Suite.Name,
				},
				"failure": common.MapStr{
					"type":  value.Failure.Type,
					"title": value.Failure.Title,
					"body":  value.Failure.Body,
				},
			},
		}
		bt.client.Publish(event)
	}
}

// Stop stops testingbeat.
func (bt *Testingbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
