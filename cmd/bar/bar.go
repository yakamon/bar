package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	rep, n, message := ParseArgs()
	bar := Bar(rep, n, message)
	fmt.Println(bar)
	os.Exit(0)
}

// ParseArgs parses command line arguments.
func ParseArgs() (rep string, n int, message string) {
	cmd := os.Args[0]

	if len(os.Args) < 2 {
		fmt.Printf("Usage: %s <rep> <n> [<message>]\n", cmd)
		os.Exit(0)
	}

	rep = os.Args[1]

	n, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("<n> must be integer: %s\n", n)
		os.Exit(1)
	}

	if len(os.Args) >= 4 {
		message = os.Args[3]
	}

	return rep, n, message
}

// Bar returns rep repeated n times.
// If message is not empty, insert it in the middle of repetition.
func Bar(rep string, n int, message string) string {
	b := bytes.Repeat([]byte(rep), n)
	if len(message) == 0 {
		return string(b)
	}

	l := len(b)
	m := fmt.Sprintf(" %s ", message)
	ml := len(m)
	i := (l - ml) / 2
	copy(b[i:i+ml], m)
	return string(b)
}
