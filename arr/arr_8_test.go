package arr

import (
	"fmt"
	"testing"
)

func TestArr_8(t *testing.T) {
	team := []string{"a","b","c","d"}
	sub := team[1:3]
	sub[0] = "e"
	sub = append(sub, "xxx")
	fmt.Println("team is: ", team)
	fmt.Println("sub is: ", sub)
}
