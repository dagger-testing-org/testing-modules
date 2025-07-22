package main

import (
	"context"
)

type Dependent struct{}

func (m *Dependent) BreakageBye() (string, error) {
	return dag.Breakage().Bye(context.Background())
}
