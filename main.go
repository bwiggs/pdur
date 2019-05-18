package main

import (
	"time"
	"fmt"
	"io"
	"os"
	"bufio"
)

func main()  {
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		in := s.Text()
		t, err := time.ParseDuration(in)
		if err != nil {
			io.WriteString(os.Stderr, "(ignoring): "+err.Error())
			continue
		}
		fmt.Println(t.Nanoseconds() / 1000000)
	}
}