# ✨ Examples and Boilerplates for Golang ✨ 

[![Go Report Card](https://goreportcard.com/badge/github.com/bt3gl/Awesome_Golang_Examples)](https://goreportcard.com/report/github.com/bt3gl/Awesome_Golang_Examples)


## In this repository 

* [Start Here](https://github.com/bt3gl/Awesome_Golang_Examples/tree/master/start_here).



--------

## Golang Cheatsheet

#### Basic Structure
```
package main

import "fmt"

func main() {
    fmt.Println("hello world")
}
```

#### Running and/or building

```
$ go run hello-world.go
```
or
```
$ go build hello-world.go
$ ./hello-world
```

#### Variables

```
var a = 'name'
var b, c int = 1, 2
d : = 'other name'
```

* Go infer the type of initialized variables.
* Syntax `:=` is shorthand for declaring and initializing. 
* Constants can be created with `const`.

#### For loops

```
for i := 7; j <= 9, j++ {
  fmt.Println(j)
}
```

#### If/Else

```
if 7%2 == 0 {
  fmt.Println("7 is even")
} else {
   fmt.Println("7 is odd")
}
```

#### Switch

* Unlike Python, Golang has a `switch` statement to express conditionals across many branches.
* You can use commas to separte multiple expressions in the same `case` statement. 
* You can use the optional `default` case in the example as well.
* You can use a `switch` without an expression as an alternate way to exress `if/else` logic.

```
switch i {
case 1:
   fmt.Println("one")
case 2:
    fmt.Println("two")
case 3:
     fmt.Println("three")
}
```

#### Arrays

* In Go, an ar



----


## References

* [Golang Official documentation](https://golang.org/).
* [Go by Example](https://gobyexample.com/).
* [/r/golang)](https://www.reddit.com/r/golang/).
* [Go weekly](https://golangweekly.com/).










---


<a href="https://www.buymeacoffee.com/miavonpizza" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/arial-pink.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" ></a>
