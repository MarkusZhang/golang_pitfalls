package arr

import (
	"fmt"
	"testing"
)

func addPerson_3(arr []string, person string) []string {
	return append(arr, person)
}

func TestArr_7(t *testing.T) {
	oldTeam := []string{}
	oldTeam = append(oldTeam, "a","b","c")
	newTeam := addPerson_3(oldTeam, "d")
	oldTeam = append(oldTeam, "e")

	fmt.Println("old team: ", oldTeam)
	fmt.Println("new team: ", newTeam)
}

