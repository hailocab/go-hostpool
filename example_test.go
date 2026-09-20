package hostpool

import (
	"github.com/bitly/go-hostpool"
)

func ExampleNewEpsilonGreedy() {
	hp := hostpool.NewEpsilonGreedy([]string{"a", "b"}, 0, &hostpool.LinearEpsilonValueCalculator{})
	hostResponse := hp.Get()
	hostname := hostResponse.Host()
	// Make a request with hostname and assign its error to err.
	_ = hostname
	var err error
	hostResponse.Mark(err)
}
