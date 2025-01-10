//go:build ignore

package main

import (
	"fmt"
	"sort"
)

type Intervals struct {
	Start int
	End   int
}

func canAttendMeetings(intervals []Intervals) bool {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Start < intervals[j].End
	})
	for i := 0; i < len(intervals)-1; i++ {
		if intervals[i].End >= intervals[i+1].Start {
			return false
		}
	}
	return true
}

func main() {
	intervals1 := []Intervals{
		{Start: 0, End: 30},
		{Start: 5, End: 10},
		{Start: 15, End: 20},
	}
	fmt.Println(canAttendMeetings(intervals1))
	intervals2 := []Intervals{
		{Start: 5, End: 8},
		{Start: 9, End: 15},
	}
	fmt.Println(canAttendMeetings(intervals2))
}
