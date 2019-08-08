package beater

type TestResult struct {
	Suite     Suite
	Name      string
	Classname string
	Duration  float64
	Skipped   bool
	Failed    bool
	Success   bool
	StdOut    string
	StdErr    string
	Failure   Failure
}

type Suite struct {
	Name     string
	Duration float64
	Metadata []Metadata
}

type Metadata struct {
	Key   string
	Value string
}

type Failure struct {
	Type  string
	Title string
	Body  string
}