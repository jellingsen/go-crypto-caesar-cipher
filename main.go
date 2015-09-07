package main

import (
	"os"
	"fmt"
	"strconv"
)

const min int = 32
const max int = 127
const diff int = max-min

func crypt(s string, num int) string {
	a := []byte(s)
	var b []byte
	fmt.Println(a)
	for _, c := range a {
		d := int(c)
		if d > max || d < min {
			fmt.Println("Invalid character!")
			break
		}
		e := d + num
		if e > max {
			e -= diff
		}
		if e < min {
			e += diff
		}
		b = append(b,byte(e))
	}
	fmt.Println(b)
	return string(b)
}

func printEx() {
	fmt.Println("Must be used with three command line arguements!")
	fmt.Printf("Use the following syntax: type number(0-%v) text\n", diff)
	fmt.Println("Use either -e for encryption, or -d for decrypting")
	fmt.Println("Example: -d 78 \"This is my plaintext\"")
}

func main() {
	args := os.Args[1:]
	if len(args) != 3 || (args[0] != "-d" && args[0] != "-e") {
		printEx()
	} else {
		num := args[1]
		msg := args[2]

		numr, err := strconv.Atoi(num)

		if err != nil || numr < 1 || numr > 127 {
			printEx()
		} else {
			if args[0] == "-d" {
				fmt.Println(crypt(msg, -numr))
			} else if args[0] == "-e" {
				fmt.Println(crypt(msg, numr))
			} else {
				fmt.Println("This message should never be shown.")
			}
		}
	}
}
