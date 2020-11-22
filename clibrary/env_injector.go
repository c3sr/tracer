package main

import (
	"os"

	"github.com/c3sr/tracer/clibrary/env"
)

func GetTraceIdEnv() string {
	return os.Getenv(env.TraceIdEnv)
}
