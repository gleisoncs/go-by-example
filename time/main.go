package main

import (
	"fmt"
	"time"
)

func main() {
	func1()
}

func func1() {
	t := time.Now().UTC()
	s1 := t.String()
	fmt.Printf("s1: %s\n", s1)

	s2 := t.Format("2006-01-02")
	fmt.Printf("s2: %s\n", s2)
}

func func2() {
	p := fmt.Println

	now := time.Now()
	p(now)

	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))
}
