package arr

import (
	"fmt"
	"testing"
)

func addPerson_2(arr []string, person string) []string {
	return append(arr, person)
}

func TestArr_6(t *testing.T) {
	oldTeam := []string{}
	oldTeam = append(oldTeam, "a")
	oldTeam = append(oldTeam, "b")
	oldTeam = append(oldTeam, "c")
	newTeam := addPerson_2(oldTeam, "d")
	oldTeam = append(oldTeam, "e")

	fmt.Println("old team: ", oldTeam)
	fmt.Println("new team: ", newTeam)
}
