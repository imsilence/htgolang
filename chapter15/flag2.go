package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var (
		addr     string
		port     int
		interval time.Duration
		verbose  bool
		help     bool
	)

	flag.StringVar(&addr, "H", "127.0.0.1", "Listen Addr")
	flag.IntVar(&port, "p", 8080, "Listen Port")
	flag.DurationVar(&interval, "i", 3*time.Second, "heartbeat interval")
	flag.BoolVar(&verbose, "v", false, "debug module")
	flag.BoolVar(&help, "h", false, "help")

	flag.Usage = func() {
		fmt.Println(`
Usage: flag [-H addr] [-p port] [-i 3s] [-v] [-h] args...

Options:`)
		flag.PrintDefaults()
	}

	flag.Parse()

	if help {
		flag.Usage()
	} else {
		fmt.Printf("flag: %d, arg: %d\n", flag.NFlag(), flag.NArg())

		fmt.Println("flag:")
		fmt.Printf("\t%s: %#v\n", "name", addr)
		fmt.Printf("\t%s: %#v\n", "port", port)
		fmt.Printf("\t%s: %#v\n", "interval", interval)
		fmt.Printf("\t%s: %#v\n", "verbose", verbose)

		fmt.Println("arg:")

		for i, arg := range flag.Args() {
			fmt.Printf("\t%d: %#v\n", i, arg)
		}
	}
}
