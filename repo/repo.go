package repo

import (
	"strconv"
)

// IDGenerator is a function that can be used by repo to generate ids
type IDGenerator func() string

// Sequential returns ids from a sequence starting from 1
func Sequential(ids ...string) IDGenerator {
	var i = 1

	return func() string {
		var id = strconv.Itoa(i)
		i++
		return id
	}
}
