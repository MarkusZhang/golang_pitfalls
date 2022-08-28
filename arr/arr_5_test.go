package arr

import (
	"fmt"
	"testing"
)

func addPerson(arr []string, person string) []string {
	return append(arr, person)
}

func TestArr_5(t *testing.T) {
	oldTeam := []string{"a","b","c"}
	newTeam := addPerson(oldTeam, "d")
	oldTeam = append(oldTeam, "e")

	fmt.Println("old team: ", oldTeam)
	fmt.Println("new team: ", newTeam)
}