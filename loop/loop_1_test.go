package loop

import (
	"fmt"
	"testing"
	"time"
)

func TestLoop_1(t *testing.T) {
	team := []string{"a", "b", "c","d"}
	for i,name := range team {
		go func(){
			fmt.Printf("award bonus to %v %v\n",i, name)
		}()
	}
	time.Sleep(time.Second)
}
