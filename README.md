# ✨ Examples and Boilerplates for Golang ✨ 

[![Go Report Card](https://goreportcard.com/badge/github.com/bt3gl/Awesome_Golang_Examples)](https://goreportcard.com/report/github.com/bt3gl/Awesome_Golang_Examples)


## In this repository 

* [Start Here](https://github.com/bt3gl/Awesome_Golang_Examples/tree/master/start_here).



--------

## Golang Cheatsheet

Note: everything that has a 🌟 in the title is *something special* in Go.

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


#### 🌟 Maps

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

#### Range

* `range` iterates over elements in a variety of data structures.
* On map, it iterates over key/pair values:

```
kvs := map[string]string{"a": "apple", "b": "banana"}
    for k, v := range kvs {
        fmt.Printf("%s -> %s\n", k, v)
    }
```

* On arrays and slices, it iterates over both the index and the value of each entry:

```
   for i, num := range nums {
        if num == 3 {
            fmt.Println("index:", i)
        }
    }
```

* On strings, iterates over Unicode code points. The first value is the starting byte index of the rune, and the second the rune itself:

```
for i, c := range "go" {
        fmt.Println(i, c)
    }
```

#### 🌟 Functions

* Go requires explicit returns.
```
func plus(a int, b int) int {
   return a + b
}
```

* When you have multiple consecutive parameters of the same type, you may omit the type:

```
func plus(a, b, c int) int {
   return a + b + c
}
```

* Go has built-in support for *multiple return values**.

```
func vals() (int, int) {
    return 3, 7
}
```

* **Variadic functions** can be called with any number of trailing arguments. For example, `fmt.Println` is a variatic function.

* Here is a function that takes an arbritary number of ints:

```
func sum(nums ...int) {
    for _, num := range nums {
        total += num
    }
    return num
}
```

* If you already have multiple args in a slice, you can apply them to a variadic function with `func(slice...)`:

```
nums := []int{1, 2, 3, 4}
sum(nums...)
```

#### 🌟 Closure

* Go supports *anonymous functions* which can form *closures*. 
* A function can return another function that is defined anonymously in the main function's body. The returned function closes over the variable i to form a closure:

```
func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()
    
    fmt.Println(nextInt())
    fmt.Println(nextInt()) 
    fmt.Println(nextInt())
    
    newInts := intSeq()
    fmt.Println(newInts())
}
```

Running this snippet prints:

```
➜ go run closures.go
1
2
3
1
```

#### Recursion

* Here is a classical factorial example for recursive functions:

```
package main
import "fmt"

func fact(n int) int {
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}
func main() {
    fmt.Println(fact(7))
}
```

#### 🌟 Pointers

* Go supports pointers, allowing you to pass references to values and records within a program.
* Assigning a value to a dereferenced pointer changes the value at the referenced address.

```
// Function argument by value
func zeroval(ival int) {
    ival = 0
}

// Function argument by pointer
func zeroptr(iptr *int) {
    *iptr = 0
}
```

* The `&i` syntax gives the memory address of `i`, i.e., a pointer to `i`.


#### 🌟 Channels

* *Channels* are pipes that connect concurrent goroutines. 
* Channels are typed by the values they convey.

```
messages := make(chan string)
go func() { messages <- "ping" }()
msg := <-messages
fmt.Println(msg)
```
prints
```
ping
```

* In the example above, the 'ping' message is successfully passed from one go routine to another via channel.

#### Signals

* Golang can handle Unix signals (e.g., `SIGTERM`, `SIGINT`).
* Signal notification works by sending `os.Signal` values on a channel we create.

```
import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    go func() {
        sig := <= sigs
        fmt.Println(sig)
        done <- true
    }()
    
    fmt.Println("Awaiting signal")
    <-done
    fmt.Println("exiting")
}
```






----


## References

* [Golang Official documentation](https://golang.org/).
* [Go by Example](https://gobyexample.com/).
* [/r/golang)](https://www.reddit.com/r/golang/).
* [Go weekly](https://golangweekly.com/).










---


<a href="https://www.buymeacoffee.com/miavonpizza" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/arial-pink.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" ></a>
