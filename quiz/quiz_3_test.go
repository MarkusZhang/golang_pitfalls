package quiz

import (
	"fmt"
	"testing"
)

func TestQuiz_3(t *testing.T) {
	gxsTeam := []string{"a","b","c"}
	males := gxsTeam[:2]
	females := gxsTeam[2:]

	updatedMales := append(males, "d")
	updatedMales[1] = "e"
	updatedFemales := append(females,"f")

	fmt.Println("males before update: ",males)
	fmt.Println("males after update: ",updatedMales)
	fmt.Println("females before update: ",females)
	fmt.Println("females after update: ",updatedFemales)
}
