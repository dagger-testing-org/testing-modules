package main

import (
	"context"
)

type Mymod struct{}

// +check
func (m *Mymod) CheckTest(ctx context.Context) (string, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", "Test passed!"}).Stdout(ctx)
}
