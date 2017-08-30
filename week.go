// importing:
package main

import "github.com/grsmv/goweek"

import "fmt"

// initializing goweek.Week struct:
//                          year
//                          |     week number (starting from 1)
//                          |     |

// retrieving slice with days (`time.Time` instances) for a given week:
func days() {
	week, _ := goweek.NewWeek(2015, 46)
	// 	return week.Days()
	fmt.Println(week.Days)
}

// // retrieving `goweek.Week` instance for a next week:
// nextWeek, err := week.Next()

// // retrieving `goweek.Week` instance for a previous week:
// previousWeek, err := week.Previous()
