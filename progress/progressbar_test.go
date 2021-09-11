package progressbar_test

import (
	// "fmt"
	"math/rand"
	"time"
	pb "github.com/mr-joshcrane/progressbar/progress"
	"testing"
)

func DoWork() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(20)
}

func TestProgress(t *testing.T) {
	bar := pb.New()
	go bar.Start(DoWork)
	err := bar.Render()
	if err != nil {
		t.Error(err)
	}
}
