# ✨ Examples and Boilerplates for Golang ✨ 

[![Go Report Card](https://goreportcard.com/badge/github.com/bt3gl/Awesome_Golang_Examples)](https://goreportcard.com/report/github.com/bt3gl/Awesome_Golang_Examples)  ![License: WTFPL](https://img.shields.io/badge/License-WTFPL-brightgreen.svg) [![Awesome](https://cdn.rawgit.com/sindresorhus/awesome/d7305f38d29fed78fa85652e3a63e154dd8e8829/media/badge.svg)](https://github.com/bt3gl/Awesome_Entrepreneur)  [![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/Naereen/StrapDown.js/graphs/commit-activity) 




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


#### Structs

* Go's structs are typed collection of fields. They are useful for grouping data together to form records.

```
type person struct {
    name string
    age int
}
```

* Struct fields can be accessed with a dot.
* An `&` prefix yields a pointer to the struct.
* Structs are mutable.


##### Methods

* Go supports methods defined on struct types.

```
type rect struct {
    width, height int
}

func (r *rect) area() int {
    return r.width * r.height
}
```

* You may want to use a pointer receiver type to avoid copying on methods calls or to allow the method to mutate the receiving struct.

#### Interfaces

* Interfaces are named collections of method signatures.

``` 
type geometry interface {
    area() float64
    perim() float64
}

func measure(g geometry) {
    fmt.Printl(g.area())
}
```

#### Errors

* It's idiomatic to communicate errors via an explicit return value.

```
func f1(arg int) (int, error) {
    if arg == 43 {
        return -1, errors.New("can't do 43")
    }
}
```



#### 🌟 WaitGroups

* To wait for multiple goroutines to finish, we can use a *WaitGroup*.

```
func worker(id int, wg *sync.WaitGrpup) {
    defer wg.Done()
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg)
    }
    wg.Wait()
}
```


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

#### 🌟 Signals

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

#### Creating an HTTP Client

* Issuing an HTTP GET request to a server:

```
resp, err := http.Get("http://curiee.com")
if err != nil {
    panic(err)
}
defer resp.Body.Close()
fmt.Println("Response status:", resp.Status)

// Print the response body
scanner := bufio.NewScanner(resp.Body)
for i := 0; scanner.Scan() && i < 5; i++ {
    fmt.Println(scanner.Text())
}
if err := scanner.Err(); err != nil {
    panic(err)
}
```


#### Creating an HTTP Server

* A fundamental concept in `net/http` servers is **handlers**. A handler is an object implementing the `http.Handler` interface.
* Functions serving as handlers take a `http.ResponseWriter` and `http.Request` as arguments. The response writer is used to fill in the HTTP response.

```
func hello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello\n")
}
```

* Handlers can also read HTTP request headers and echo them into the response body:

```
func headers(w http.ResponseWritter, req *http.Request) {
    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprint(w, "%v: %v\n", name, h)
        }
    }
}
```

* Handlers on server routes can be registered using `http.HandleFunc`, which sets up the default route and takes a function as an argument:

```
func main() {
    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", header)
    http.ListenAndServe(":8090", nil)
}
```



#### 🌟 Context

* A `Context` carries deadlines, cancellation signals, and other request-scoped values accross API boundaries and goroutines.
* In the example below, a `context.Context` is created for each request by the `net/http` machinery, and it's available withe `Context()` method.

```
import (
    "fmt"
    "net/http"
    "time"
)

func hello(w http.ResponseWriter, req *http.Request) {
    
    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")
    
    select {
        case <-time.After(10 * time.Second):
            fmt.Fprintf(w, "hello\n")
        case <-ctx.Done():
            err := ctx.Err()
            fmt.Println("server:", err)
            internalError := http.Status
    
    }
}

func main() {
    http.HandleFunc("/hello", hello)
    http.ListenAndServe(":80900", nil)
}
```



#### 🌟 Exec'ing Processes

* To replace the current Go process with another (Go or non-Go), we can use `exec`.
* In the snippet below, we get an absolute path for a binary (`ls`) and then create a slice with the arguments, load environmnet variables, and then run `exec`:

```
package main

import (
    "os"
    "os/exec"
    "syscall"
)

func main() {

    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }
    args := []string{"ls", "-a", "-l", "-h"}
    env := os.Environ()
    
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
```

#### 🌟 Spawning Processes

* It's possible to spawn processes from go using `exec.Command()`. This command creates an object to represent the external process.
```
dateCmd := exec.Command("date")
dateOut, err := dateCmd.Output()
if err != nil {
    panic(err)
}

fmt.Println(string(dateOut))
```

* The method `.Output` is a helper that handles the common case of running a command, waiting for it to finish, and collecting its output.
* In another example we can pipe data to external processes on its `STDIN` and `STDOUT`. 

```
grepCmd := exec.Command("grep", "hello")
grepIn, _ := grepCmd.StdinPipe()
grepOut, _ := grepCmd.StdoutPipe()
grepCmd.Start()
grepIn.Write([]byte("hello")
grepIn.Close()
grepByte, _ := ioutil.ReadAll(grepOut)
grepCmd.Wait()
fmt.Println(string(grepBytes))
```

----


## References


### Basic

* [Golang Official documentation](https://golang.org/).
* [Go by Example](https://gobyexample.com/).

### Communities

* [/r/golang)](https://www.reddit.com/r/golang/).
* [Go weekly](https://golangweekly.com/).


### Good readings

* [Error handling and Go](https://blog.golang.org/error-handling-and-go).









---


<a href="https://www.buymeacoffee.com/miavonpizza" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/arial-pink.png" alt="Buy Me A Coffee" style="height: 51px !important;width: 217px !important;" ></a>
