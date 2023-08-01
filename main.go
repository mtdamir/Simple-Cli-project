package main

import (
	"flag"
	"fmt"
)

// func main() {
// 	fmt.Println(os.Args)
// 	fmt.Printf("\n\n")
// 	args := os.Args
// 	fmt.Println("len(Args)", len(args))

// 	if len(args) < 2 {
// 		// fmt.Println("error  - there is no input")
// 		// unreachable code
// 		// os.Exit(1)
// 		log.Fatalf("there is no input - len of args: %d", len(args))
// 	}

// 	name1 := args[1]
// 	fmt.Println("Hello", name1)
// 	name2 := args[2]
// 	fmt.Println("Hello2", name2)
// 	name1 = strings.Replace(name1, "--first_name=", "", 1)
// 	if name1 == "" {
// 		fmt.Println("Hello there")
// 	} else {
// 		fmt.Println("Hello", name1)
// 	}

// }

func main() {
	firstName := flag.String("word", "default value", "description")
	age := flag.Int("age", 23, "age of user")
	flag.Parse()
	fmt.Println("name", *firstName, *age)

	rflags := flag.Args()
	fmt.Println("len remaining flags", len(rflags))
	fmt.Println("len remaining flags", rflags)

	fmt.Println("please enter ypur name")
	var firstname, lastName string
	// Scanln use for input yor cli program
	fmt.Scanln(&firstname, &lastName)
	fmt.Println("name", firstname, lastName)
}
