package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

// How do we run the code in our project ???
// -  go run <file-name>

// GO Commands
// 1) go build => compiles a bunch of go source code files
// 2) go run => compiles and executes one or two files
// 3) go fmt => Formats all the code in each file in the current directory
// 4) go install  => compiles and installs a package
// 5) go get  => Downloads the raw source code of someone else's package
// 6) go test  =>  runs any tests associated with the current project

// pakage <name>  => we need to declare it in every go file at the start
// => Suppose if a go file has package temp   => That file belongs to package temp

// Types of Packages in Go
// 1) Executable  => Generates a file that we can run
// 2) Reusable  =>  Code used as "helpers". Good place to put reusable logic

// - How to Know which type of package we are using ????
// IMPP :  If we mention  package main => Then, it is used to make an executable type package

// => Other than package main, any other package <pckg-name>  => can be identified as re-usable package

// What does import "fmt" mean ??  => It is a standard library
