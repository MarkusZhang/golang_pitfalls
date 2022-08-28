package loop

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLoop_2(t *testing.T) {
	team := []string{"a", "b", "c","d"}
	var wg sync.WaitGroup
	wg.Add(1)
	for _,name := range team {
		go func(){
			fmt.Printf("award bonus to %v\n", name)
			time.Sleep(time.Millisecond*100)
			wg.Done()
		}()
	}
	wg.Wait()
}

