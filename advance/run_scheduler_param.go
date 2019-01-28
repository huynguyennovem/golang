package main

import "fmt"

/**
	purpose:
		- passing separate parameters into a Object type
		- func does not need to define too many parameters
	for example:
		- passing params:	"scheduler name", 100
		- result: 			schedulerOptions("scheduler name", 100) and New func has just defined a param (opts)
	design pattern:
		- TBD
 */

type schedulerOptions struct {
	schedulerName      string
	bindTimeoutSeconds int64
}

type Option func(*schedulerOptions)

func WithName(schedulerName string) Option {
	return func(o *schedulerOptions) {
		o.schedulerName = schedulerName
	}
}

func WithBindTimeoutSeconds(bindTimeoutSeconds int64) Option {
	return func(o *schedulerOptions) {
		o.bindTimeoutSeconds = bindTimeoutSeconds
	}
}

func New(opts ...func(o *schedulerOptions)) {
	var scheOpt = schedulerOptions {
		schedulerName:      "schedulerName",
		bindTimeoutSeconds: 10,
	}
	for _, opt := range opts {
		opt(&scheOpt)
	}

	fmt.Println(scheOpt.schedulerName)
	fmt.Println(scheOpt.bindTimeoutSeconds)
}

func main() {
	New(WithName("scheduler name"), WithBindTimeoutSeconds(100))
}
