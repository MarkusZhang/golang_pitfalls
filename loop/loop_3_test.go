package loop

import (
	"fmt"
	"testing"
	"time"
)

func TestLoop_3(t *testing.T) {
	team := []string{"a", "b", "c","d"}
	for _,name := range team {
		go awardBonus(name)
	}
	time.Sleep(time.Millisecond*100)
}

func awardBonus(name string) {
	fmt.Printf("award bonus to %v\n", name)
}
