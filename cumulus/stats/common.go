package stats

import "time"

type operation int

const (
	inc operation = iota
	report

	period = 1 * time.Minute
)

type Reporter interface {
	Report()
}

var reportersChannel = make(chan Reporter)
var reporters []Reporter = make([]Reporter, 0)

func init() {
	go func() {
		for r := range reportersChannel {
			reporters = append(reporters, r)
		}
	}()
}

func Report() {
	for _, r := range reporters {
		r.Report()
	}
}
