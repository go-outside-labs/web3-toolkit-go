# ✨ Examples and Boilerplates for Golang ✨ 

[![Go Report Card](https://goreportcard.com/badge/github.com/bt3gl/Awesome_Golang_Examples)](https://goreportcard.com/report/github.com/bt3gl/Awesome_Golang_Examples)


## In this repository 

* [Start Here](https://github.com/bt3gl/Awesome_Golang_Examples/tree/master/start_here).



--------

## Golang Cheatsheet

Note: everything that has a 🌟 in the title is something special in Go.

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

#### 🌟 Switch

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

* In Go, an `arrray` is aa numbered sequence of elements of a specific length.
* The type of elements and length are both part of the array's part.
* By default, an array is zero-valued.

```
var a [5]int
```

* Use this syntax to declare and initalize an array in one line.

```
b := [5]int{1, 2, 3, 4, 5}
```
* Arrays can be made multi-dimensional:

```
var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
```

#### 🌟 Slices

* Slices are a key data type in golang, giving an interface to sequences.
* Unlike arrays, slices are typed only by the elements they contain (not the number of elements).

```
s := make([]string, 3)
```

* Slices support several operations such as: `append` and  `copy`.
* Slices can be multi-dimensitonal as well:

```
twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
```


#### Maps

* Maps are Go's built-in *associative data types*, similar to `dicts` in Python.
* To create an empty map, use `make`:

```
m := make(map[string]int)
```

* The builtin `delete` removes key/value pairs from a map.
* The optional return value when getting a value from a map indicates if the key was present in the map.

```
_, exist := m["key"]
```

* You can declare and initialize a new map in the same line:

```
n := map[string]int{"foo": 1, "bar": 2}
```



----


## References

* [Golang Official documentation](https://golang.org/).
* [Go by Example](https://gobyexample.com/).
* [/r/golang)](https://www.reddit.com/r/golang/).
* [Go weekly](https://golangweekly.com/).










---


<a href="https://www.buymeacoffee.com/miavonpizza" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/arial-pink.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" ></a>
