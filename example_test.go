package hostpool

import (
	"fmt"
)

func ExampleNewEpsilonGreedy() {
	hp := NewEpsilonGreedy([]string{"a", "b"}, 0, &LinearEpsilonValueCalculator{})
	hostResponse := hp.Get()
	fmt.Println(hostResponse.Host())
	var err error // (make a request with hostname)
	hostResponse.Mark(err)
}
