package pluginvalidation

import "fmt"

type Report struct {
	PluginCount int
	Errors      []string
}

func (r Report) OK() bool {
	return len(r.Errors) == 0
}

func (r *Report) add(where, message string) {
	r.Errors = append(r.Errors, fmt.Sprintf("%s: %s", where, message))
}
