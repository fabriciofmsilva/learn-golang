package main

import "github.com/grsmv/goweek"

import "fmt"

// retrieving slice with days (`time.Time`instances) for a given week:
func days() {
	week, _ := goweek.NewWeek(2015, 46)
	// return week.Days()
	fmt.Println(week.Days)
}
