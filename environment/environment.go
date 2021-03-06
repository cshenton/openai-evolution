// Package environment defines an interface for assessing the fitness of agents
// on particular environments, as well as some concrete environments, such as
// a simple in-memory MNIST evaluation as well as an rpc link to OpenAI gym.
package environment

import "github.com/cshenton/evolution/agent"

// Environment defines an environment for evaluating agents.
type Environment interface {
	// Eval evaluates agent a and returns a fitness f.
	Eval(a agent.Agent) (f float64)
}

// Result is the result of evaluating an agent.
type Result struct {
	Seed    int64
	Fitness float64
}
