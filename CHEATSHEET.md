## Golang Cheatsheet

<br>

Note: everything that has a 🌟 in the title is *something special* in Go.

<br>

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
for j := 7; j <= 9, j++ {
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

#### Strings 

* Go offers lots of supports for string formating:

```
package main

import (
    "fmt"
    "os"
)

type point struct {
    x, y int
}

func main() {

    p := point{1, 2}
    fmt.Printf("%v\n", p)

    // For a struct, the %+v variant will include 
    // the struct’s field names
    fmt.Printf("%+v\n", p)
    
    // Print the source code snippet that would 
    // produce that value
    fmt.Printf("%#v\n", p)
    
    // To print the type of a value, use %T
    fmt.Printf("%T\n", p)
    
    // Formatting booleans
    fmt.Printf("%t\n", true)
    
    // base-10 formatting
    fmt.Printf("%d\n", 123)
    
    // prints a binary representation
    fmt.Printf("%b\n", 14)
    
    //  prints the character corresponding to the given int
    fmt.Printf("%c\n", 33)
    
    // provides hex encoding
    fmt.Printf("%x\n", 456)
    
    // several formatting options for float
    fmt.Printf("%f\n", 78.9)

    // basic string printing use %s
    fmt.Printf("%s\n", "\"string\"")
    
    // double-quote strings
    fmt.Printf("%q\n", "\"string\"")
    
    // print a representation of a pointer
    fmt.Printf("%p\n", &p)

    // print error
    fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
```


* The standard library's `strings` package provides many useful string-related functions.


```
package main

import (
    "fmt"
    s "strings"
)

var p = fmt.Println

func main() {

    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
```

prints

```
Contains:   true
Count:      2
HasPrefix:  true
HasSuffix:  true
Index:      1
Join:       a-b
Repeat:     aaaaa
Replace:    f00
Replace:    f0o
Split:      [a b c d e]
ToUpper:    TEST
Len:  5
Char: 101
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

```
func main() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}
```

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


```
type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```


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

* Interfaces are named collections of method signatures and are used to achieve polymorphism in Go.

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

#### 🌟 Panic

* A panic means something went unexpectedly wrong. It's used to fail fast on errors that shouldn't occur during normal operation, or that we aren't prepared to handle gracefully.

```
_, err := os.Create("/tmp/file")
if err != nil {
    panic(err)
}   
```

#### 🌟 Goroutines

* A goroutine is a lightweight thread of execution.

```
go f()
```

* You can also start a goroutine for an anonymous function call:

```
go func(msg string){
    fmt.Prinln(msg)
}("hello")
```



#### 🌟 Defer

* `defer` is used to ensure that a function call is performed later in the program's execution, usually for purposes of cleanup (such as `finally` in other languages).

```
func main() {

    f := createFile("/tmp/defer.txt")
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
    fmt.Println("closing")
    err := f.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}
```


#### 🌟 Channels

* *Channels* are pipes that connect concurrent goroutines. 
* Channels are typed by the values they convey.

```
messages := make(chan string)

go func() { 
    messages <- "ping" 
}()
    
msg := <- messages
fmt.Println(msg)
```
prints
```
ping
```

* In the example above, the 'ping' message is successfully passed from one go routine to another via channel.

* By default, channels are unbuffered: they will only accept sends (`chan <-`) if there is a corresponding receive (`<- chan`) ready to receive the sent value.

* Buffered channels accept a limited number of values (`2` in the example below) without a corresponding receiver for those values.
```
messages := make(chan string, 2)
messages <- "buffered"
messages <- "channel"

fmt.Println(<-messages)
```

* You can use channels to synchronize execution accross goroutines (or ou can use `Waitgroups`).

```
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println("done")
    done <- true
}

func main() {
    // Start a worker goroutine, giving 
    // it the channel to notify on.
    done := make(chan bool, 1)
    go worker(done)
    
    // Block until we receive a notification from 
    // the worker on the channel
    <- done
}
```


#### 🌟 Select

* Go's `select` lets you wait on multiple channel operations.
* In the example below, select across two channels:

```
c1 := make(chan string)
c2 := make(chan string)

go func() {
    time.Sleep(1)
    c1 <- "one"
}()

go func() {
    time.Sleep(2)
    c2 <- "two"
}()

for i := 0; i < 2; i++ {
    // use select to await both values simultaneously
    // printing each one as it arrives.
    select {
        case msg1 := <-c1:
            fmt.Println("received", msg1)
        case msg2 := <-c2:
            fmt.Println("received", msg2)
    }
}
```


#### Timeouts

* Timeouts are important for programs that connect external resources or that need to bound time.

```
c1:= make(chan string, 1)
go func() {
    time.Sleep(2 * time.Second)
    c1 <= "result 1"
}()

select {
case res := <-c1:
    fmt.Println(res)
case <- time.After(1 * time.Second):
    fmt.Println("timeout 1")

}
```

#### Non-Blocking Channel Operations

* Basic sends and receives on channels are blocking. We can use `select` with a `default` clause to implement *non-blocking* sends, received, and even non-blocking multi-way `selects`.

```
messages := make(chan string)
signals := make(chan bool)

select {
case msg := <-messages:
    fmt.Println("received message", msg)
default:
    fmt.Println("no message received")
}
```

#### Closing Channels

* Closing a channel indicates that no more values will be sent on it and it can communicate completation to the channel's receivers.

```
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)
    
    // Worker goroutine, receiving several jobs.
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // Sends 3 jobs to the worker over the jobs channel,
    // then closes it.
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // Await the worker, using the syncronization approach.
    <-done
}
```

#### Channels as Queues

* We can use `range` to interate over values received from a channel.

```
func main() {

    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)

    // Because we closed the channel above,
    // the iteration terminates after receiving
    // the 2 elements.
    for elem := range queue {
        fmt.Println(elem)
    }
}
```

#### 🌟 Timers & Tickers

* Go's built-in `timer` and `ticker` features make it possible to execute Go code at some point in the future or repeatedly.

* Timers represent a single event in the future. You tell the timer how long you want to wait, and it provides a channel that will be notified at that time.

```
timer1 := time.NewTimer(2 * time.Second)

// This blocks on the timer's channel C until it 
// sends a value indicating that the timer fired.
<-timer1.C
```

* Tickers are for when you want to do something repeatedly at regular intervals.
* The mechanism is similar to timers: a channel that is sent values.

```
ticker := time.NewTicker(500 * time.Millisecond)
done := make(chan bool)

go func() {
    for {
        select {
            case <=done:
                return
            case t := <-ticker.C:
                fmt.Println("Tick at", t)
        }
    }
}()

ticker.Sleep(1600 * time.Millisecond)
ticker.Stop()
done <- true
fmt.Println("Ticker stopped")
```

#### Worker Pools

* We can implement a woker pool in Golang using goroutines and channels:

```
// These workers receive work on the jobs channel and
// send the corresponding results on results.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "started  job", j)
        time.Sleep(time.Second)
        fmt.Println("worker", id, "finished job", j)
        results <- j * 2
    }
}

func main() {

    // In order to use our pool of workers, we need to send their
    // work and collect their results. We make 2 channels for this:
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    // Starts up 3 workeds, initially blocked since
    // there are no jobs yet.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Send 5 jobs and then close that channel.
    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    // Collect all the results of the work.
    // This also ensures that the worker goroutines has finished.
    for a := 1; a <= numJobs; a++ {
        <-results
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
