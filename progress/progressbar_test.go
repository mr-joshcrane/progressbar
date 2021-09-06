package progressbar_test

import (
	pb "github.com/mr-joshcrane/progressbar"
	"testing"
	"time"
)

func TestProgress(t *testing.T) {
	t.Parallel()
	timeAtStart := time.Now().Unix()
	bar := pb.New()
	time.Sleep(2 * time.Second)
	bar.Stop()

	started, elapsed, finished := bar.State()
	if started.Unix() != timeAtStart {
		t.Fatal("did not capture the start time")

	}
	if elapsed == 0 {
		t.Fatal("elapsed did not update")
	}
	if finished != true {
		t.Fatal("progress bar failed to finish")
	}
}
