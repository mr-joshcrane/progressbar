package progressbar

import (
	"time"
)

type ProgressBar struct {
	started  time.Time
	elapsed  int64
	finished bool
}

func New() ProgressBar {
	return ProgressBar{time.Now(), 0, false}
}

func (bar *ProgressBar) Stop() {
	started := bar.started.Unix()
	elapsed := time.Now().Unix() - started
	bar.elapsed = elapsed
	bar.finished = true
}

func (bar *ProgressBar) State() (started time.Time, elapsed int64, finished bool) {
	return bar.started, bar. elapsed, bar.finished	
}