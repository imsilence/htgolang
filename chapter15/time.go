package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Date(2019, time.April, 25, 11, 18, 20, 0, time.UTC))

	fmt.Println(time.Now())

	fmt.Println(time.Unix(0, 0))

	fmt.Println(time.Parse("2006-01-02 15:04:05", "2019-04-05 11:20:30"))

	now := time.Now()
	fmt.Println(now.Format("2006/01/02 15:04:05"))

	fmt.Println(now.Unix())
	fmt.Println(now.UnixNano())

	tomorrow := now.Add(24 * time.Hour)
	fmt.Println(tomorrow)

	fmt.Println(now.After(tomorrow))
	fmt.Println(now.Before(tomorrow))

	fmt.Println(now.Clock())
	fmt.Println(now.Date())

	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	fmt.Println(now.String())

}
