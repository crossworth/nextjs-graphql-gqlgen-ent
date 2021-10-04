package graph

import (
	"strconv"
)

func stringIDToInt(id string) int {
	i, _ := strconv.Atoi(id)
	return i
}
