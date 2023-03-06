## Constants

```golang
const myConst int = 42
// myConst = 5 # compiler throws an error because it is constant
fmt.Printf('%v, %T' ,myConst, myConst)
```

iota is  used to represent constant increasing sequence and iota is scoped to constant block. iota is initalized with 0 in integer.

```golang
const (
    a = iota
    b
    c
)
func main(){
fmt.Printf('%v' ,a) // 0
fmt.Printf('%v' ,b) // 1
fmt.Printf('%v' ,c) // 2
}
```

_ is a write only variable in go

```golang
const (
    _ = iota
    a 
    b
    c
)
func main(){
fmt.Printf('%v' ,a) // 1
fmt.Printf('%v' ,b) // 2
fmt.Printf('%v' ,c) // 3
}
```

## Arrays 

- Only hold same type. Need to specify the number of elements

`grades := [3]int{97,95,90}`

- No need to specify the number of elements, replaced by three dots

`grades := [...]int{97,95,90}`

- Declare array without elements

`var students [3]string`

`len(students)` -> length of the array

```golang
var students [3]string
students[0] = "Lisa"
```

## Slice

Slice can be initialized by,

`a := []int{1,2,3}`

Slice is a small part of array. Array will make a duplication but slice reference to the same memory address

```golang
func main(){
    a := []int{1,2,3,4,5,6,7,8,9,10} // its a slice
    b := a[:] // slice of all elements
    c := a[3:] // slice from 4th element
}
```

```golang
func main(){
    a := [...]int{1,2,3,4,5,6,7,8,9,10} // its an array now
    b := a[:] // slice of all elements
    c := a[3:] // slice from 4th element
}
```

Use make function to create a slice,

- Make an integer slice with len() and cap() = 3

    `a := make([]int,3)`

- Make an integer slice with len() = 3 and cap() = 100. Unlike arrays slice need not have fixed size because we can add or delete element from slice

    `a := make([]int, 3, 100)`

- Add an element to the slice

    `a = append(a,10)`
    
    `a = append(a,10,11,12,13)`
    
- Delete centre element from slice 

```golang
func main(){
    a := []int{1,2,3,4,5} 
    b := append(a[:2],a[3:]...) // ... will spread the appended elements
    fmt.Println(b)
}
```

## Maps

- key -> value

```go
func main() {
    statePopulation := map[string]int{
    "California": 200,
    "Texas":      300,
    "New York":   400,
    }
    fmt.Println(statePopulation)
}
```

- Use make function

`statePopulation := make(map[string]int)`

`statePopulation := make(map[string]int,10)`

- Get the value using key

`fmt.Println(statePopulation["Texas"])`

- Add a new value with key

`statePopulation["Georgia"] = 1012`

- Delete an entry from map

`delete(statePopulation, "Georgia")`

- Check if the key is there or not

`_, ok := statePopulation("Ohio") //ok -> true if value is there, false if not there`

- Length of a map

`len(statePopulation)`

## Structs

```golang
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John",
		companions: []string{
			"Abhi",
			"Akhil",
		},
	}

	fmt.Println(aDoctor)
}
```

- Mix any kind of data together
- To get actorName -> `aDoctor.actorName`
- To get companion slice -> `aDoctor.companions`
- Positional declaration without specifying the varible name which works on specified order. NOT RECOMMEND TO USE IT

```golang
func main() {
	aDoctor := Doctor{
		   3,
		"John",
		[]string{
			"Abhi",
			"Akhil",
		},
	}

	fmt.Println(aDoctor)
}
```

- Anonymous struct - without type keyword and cannot be used again and again

`aDoctor := struct{name string}{name: "John"}`

- Unlike maps and slices, structs are using independent dataset. When manipulating, copies are created

```golang
func main(){
    aDoctor := struct{name string}{name: "John"}
    anotherDoctor := aDoctor
    anotherDoctor.name = "Tom"
    fmt.Println(aDoctor)        //John
    fmt.Println(anotherDoctor)  //Tom
}
```
- Go doesnt support inheritance and all but using struct another struct can be embedded

```golang
type Animal struct {
	Name   string
	Origin string
}

type Bird struct {
	Animal  // get the property of another struct : Embedding another Animal
	Speed  float32
	Canfly bool
}

func main() {
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.Speed = 10
	b.Canfly = false

	fmt.Println(b.Name)
}

```

- The above can be simplyfied with : 

```golang
func main() {
	b := Bird{
		Animal: Animal{Name: "Emu", Origin: "Australia"},
		Speed:  30,
		Canfly: false,
	}

	fmt.Println(b.Name)
}
```

## If statements

```golang
if true {
	fmt.Println("The test is true)
}
```

- Get the boolean result by interogating the map and using that result as the test for if

```golang
func main() {
    statePopulation := map[string]int{
    "California": 200,
    "Texas":      300,
    "New York":   400,
    }

	if pop,ok := statePopulation["Texas"]; ok {
		fmt.Println(pop)  // 300
	}
}
```

- Comparison 

```golang
func main(){
	number := 50
	gueass := 30
	if guess < number{
		fmt.Println("Too low")
	}
	if guess < 1 || guess > 100 {
		fmt.Println("The guess must be between 1 and 100)
	}
}
```

```golang
if {}
else if {}
else {
	if{}
	if{}
	.
	.
}
```

## Switch Statement

```golang
func main() {
	switch 2 {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two") 	// this will be executed
	default:
		fmt.Println("not one or two")
	}
}
```

- Support multiple conditions

```golang
func main() {
	switch 5 {
	case 1, 5, 10:
		fmt.Println("one, five or ten")	// this will work
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}
```

- Using an expression

```golang
func main() {
	switch i := 2+3;i {
	case 1, 5, 10:
		fmt.Println("one, five or ten")	// this will work (2+3 = 5)
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not one or two")
	}
}
```

- Tagless syntax

```golang
func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")	// this will work, because its seen first
	case i<=20:
		fmt.Println("less than or equal to twenty")
	default:
		fmt.Println("greater than twenty")
	}
}
```

- Use fallthrough to continue checking the conditions without break

```golang
func main() {
	i := 10
	switch {
	case i <= 10:
		fmt.Println("less than or equal to ten")	// this will work, because its seen first
		fallthrough
	case i<=20:
		fmt.Println("less than or equal to twenty") // this will also work, due to fallthrough , indentionally make next case after fallthrough to execute
	default:
		fmt.Println("greater than twenty")
	}
}
```

- Type switching

```golang
func main() {
	var i interface{} = 1
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")	// this will work because i is integer
	case float64:
		fmt.Println("i is float64")
	case string:
		fmt.Println("i is string")
	default:
		fmt.Println("another type")
	}
}
```

## Looping

```golang
func main() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}
```

- Two variables

```golang
func main() {
	for i,j := 0,0 ; i < 5; i,j = i+1, j+1 {
		fmt.Println(i,j)
	}
}
```

- Naked for loop

```golang
func main() {
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i)
	}
	fmt.Println(i) // the scope of i will be get outside the loop
}
```

- Another naked loop - do while loop

```golang
func main() {
	i := 0
	for  i < 5  {
		fmt.Println(i)
		i++
	}
}
```

- Infinite loop

```golang
func main() {
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 5 {
			break
		}
	}
}
```

- `continue` statement is also used to skip the iteration
- Nested loop

```golang
func main() {
	for i:= 0 ; i < 5; i++ {
		for j:=0; j < 5; j++ {
			fmt.Println(i,j)
		}
	}
}
```

- Breaking the outer loop

```golang
func main() {
Loop:
	for i:= 0 ; i < 5; i++ {
		for j:=0; j < 5; j++ {
			fmt.Println(i,j)
			if i*j >= 3{
				break Loop	// break the outer loop
			}
		}
	}
}
```

- Work with collections - slice

```golang
func main(){
	s := []int{1,2,3}
	for k,v := range s {  // key(index) and value
		fmt.Println(k,v)
	}
}
```

- Work with collections - map

```go
func main() {
    statePopulation := map[string]int{
    "California": 200,
    "Texas":      300,
    "New York":   400,
    }
	for k,v := range statePopulation {
    	fmt.Println(k,v)
	}
}
```

- Work with string

```go
func main() {
   s : = "hello"
	for k,v := range s {
    	fmt.Println(k, string(v))
	}
}
```

## Control Flow

### Defer

- Executes after the main function and before it return any result to the calling function
- Open a resource, check a resource and close the resoure (with defer keyword)
- Defer not take a function instead take a function call

```golang
func main(){
	fmt.Println("start")
	defer fmt.Println("middle") // print atlast after those two
	fmt.Println("end")
}
```

- Takes value at the time defer is called and not at the time the function is executed

```golang
func main(){
	a := "start"
	defer fmt.Println(a) // will print "start" and not "end"
	a := "end"
}
```

### Panic

- When go can't continue to function and cannot figure out what it needed to do
- Program exits with panic string

```golang
func main(){
	fmt.Println("start")
	panic("Something bad happened")
	fmt.Println("end)
}
```

- if defer and panic is there, defer will be executed before the function panic
- function will stop executing but deferred function will still fire

### Recover

- recover() returns nil if the application is panicking and isn't nil then it returns the error that causing the application to panic 
- only useful in deferred function
- current function will not attempt to continue, but higher functions in call stack will

## Pointer

```golang
func main(){
	var a int = 42
	var b *int = &a // declaring b as a pointer varible
	fmt.Println(a,*b) //use * to dereference the variable, get the value

	a = 50 // will change both variable
	*b = 30 // will change both variable 
}
```

```golang
func main(){
	a := [3]int{1,2,3}
	b := &a[0]
	c := &a[1]
	fmt.Printf("%v %p %p" , a,b,c)
}
```

```golang
type myStruct struct{
	foo int
}

func main(){
	var ms *myStruct
	ms = new(myStruct) // initialize the struct with 0
	(*ms).foo = 42 // modify the variable
	fmt.Println((*ms).foo)
}
```

- the above ugly syntax can be corrected by : 

```golang
type myStruct struct{
	foo int
}

func main(){
	var ms *myStruct
	ms = new(myStruct) // initialize the struct with 0
	ms.foo = 42 // this will also work 
	fmt.Println(ms.foo)
}
```

## Function

```golang
func main(){
	sey("hi","hello")
}
func sey(a ,b string){  // because both are string only specify one
	fmt.Println(a,b)
}
```

- Pointers in function

```golang
func main(){
	g := "hello"
	h := "Abhi"
	sey(&g , &h)
	fmt.Println(h) // print "Abhijith"
}

func sey(g ,h *string){
	*h = "Abhijith"
	fmt.Println(*h)  // print "Abhijith"
}
```

- For a large datastructure, it is better to pass a pointer
- Vairadic Parameter

```golang
func main(){
	sum("The sum is",1,2,3,4,5)
}
func sum(msg string, values ...int){ // variatic paramter need to be at last
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result += v
	}
	fmt.Println(result,msg)
}
```

- Returning  as a pointer

```golang
func main(){
	s := sum(1,2,3,4,5)
	fmt.Println("Sum is", *s)
}

func sum(values ...int) *int { 
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result += v
	}
	return &result
}
```

- Named return value

```golang
func main(){
	s := sum(1,2,3,4,5)
	fmt.Println("Sum is", s)
}

func sum(values ...int) (result int) { 
	fmt.Println(values)
	result := 0
	for _,v := range values {
		result += v
	}
	return 
}
```
- Returning multiple value

```golang
func main() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) { // returning multiple value
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
```

- Anonymous function

```golang
func main(){

	func (){
		fmt.Println("Hello")
	}() // calling the function after defining it
}
```

- Function as a variable

```golang
func main(){
	f := func(){ // var f func() = func() {
		fmt.Println("Hello")
	}
	f()
}
```

- Method
    - It is a function in a known context

```golang
type greeter struct {
	greeting string
	name     string
}
// type count int -> also works
func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

func (g greeter) greet() { // we are getting the copy of struct
	fmt.Println(g.greeting, g.name)
}
``` 

```golang
type greeter struct {
	greeting string
	name     string
}
// type count int -> also works
func main() {
	g := greeter{
		greeting: "Hello",
		name:     "Go",
	}
	g.greet()
}

func (g *greeter) greet() { // we are getting not copy of struct actual location
	fmt.Println(g.greeting, g.name)
}
``` 

## Interfaces

- Instead of data, interfaces stores method definition

```golang
func main() {
	var s Shapes = circle{2}
	var d Shapes = square{2, 5}
	fmt.Println(s.area())
	fmt.Println(d.area())
}

type circle struct {
	radius int
}
type square struct {
	length  int
	breadth int
}

type Shapes interface {
	area() int
}

func (c circle) area() int {
	return c.radius * c.radius
}

func (s square) area() int {
	return s.length * s.breadth
}

```

```golang
func main() {
	s := circle{2}
	d := square{2, 5}
	shape := []Shapes{s, d} // using the interface type for array
	fmt.Println(shape[0].area())
	fmt.Println(shape[1].area())
}

type circle struct {
	radius int
}
type square struct {
	length  int
	breadth int
}

type Shapes interface {
	area() int
}

func (c circle) area() int {
	return c.radius * c.radius
}

func (s square) area() int {
	return s.length * s.breadth
}
```

```golang
type Shape interface {
    area() float64
}

type Rectangle struct {
    length float64
    width  float64
}

func (r Rectangle) area() float64 {
    return r.length * r.width
}

type Circle struct {
    radius float64
}

func (c Circle) area() float64 {
    return math.Pi * c.radius * c.radius
}

type Square struct {
    side float64
}

func (s Square) area() float64 {
    return s.side * s.side
}

func main() {
    shapes := []Shape{
        Rectangle{length: 2, width: 3},
        Circle{radius: 5},
        Square{side: 4},
    }

    for _, shape := range shapes {
        fmt.Println(shape.area())
    }
}
```

```golang
type Car struct {
    name      string
    model     string
    engine    Engine
}

type Engine interface {
    start() string
}

type GasEngine struct {
    typeOfFuel string
}

func (ge GasEngine) start() string {
    return "Starting gas engine..."
}

type ElectricEngine struct {
    voltage float64
}

func (ee ElectricEngine) start() string {
    return "Starting electric engine..."
}

func main() {
    gasEngine := GasEngine{typeOfFuel: "petrol"}
    electricEngine := ElectricEngine{voltage: 12.5}

    car1 := Car{name: "Honda", model: "Civic", engine: gasEngine}
    car2 := Car{name: "Tesla", model: "Model 3", engine: electricEngine}

    fmt.Println(car1.engine.start())
    fmt.Println(car2.engine.start())
}
```

```golang
type Car struct {
    name   string
    model  string
    engine []Engine
}

type Engine interface {
    start() string
}

type GasEngine struct {
    typeOfFuel string
}

func (ge GasEngine) start() string {
    return "Starting gas engine..."
}

type ElectricEngine struct {
    voltage float64
}

func (ee ElectricEngine) start() string {
    return "Starting electric engine..."
}

func main() {
    gasEngine := GasEngine{typeOfFuel: "petrol"}
    electricEngine := ElectricEngine{voltage: 12.5}

    car1 := Car{name: "Honda", model: "Civic", engine: []Engine{gasEngine}}
    car2 := Car{name: "Tesla", model: "Model 3", engine: []Engine{electricEngine}}

    for _, e := range car1.engine {
        fmt.Println(e.start())
    }

    for _, e := range car2.engine {
        fmt.Println(e.start())
    }
}
```

## Goroutines

- Implement concurrency
- Go routine will execute separately from the main function

```golang
func main() {
	go sayHello()
	time.Sleep(100 * time.Millisecond)
}

func sayHello() {
	fmt.Println("Hello")
}

```

- Using anonymous function

```golang
func main(){
	var msg = "Hello"
	go func(){
		fmt.Println(msg)
	}()
	time.Sleep(100 * time.Millisecond)
}
```

```golang
func main(){
	var msg = "Hello"
	go func(){
		fmt.Println(msg)
	}()
	msg = "Goodbye" // it prints instead of "Hello"
	time.Sleep(100 * time.Millisecond)
}
```

- Solution for the above problem

```golang
func main(){
	var msg = "Hello"
	go func(msg string){
		fmt.Println(msg)
	}(msg)
	msg = "Goodbye" // it prints "Hello"
	time.Sleep(100 * time.Millisecond)
}
```

- Use waitgroup to remove usage of sleep

```golang
func main(){
	var msg = "Hello"
	wg.Add(1) // increment 1
	go func(msg string){
		fmt.Println(msg)
		wg.Done() // decrement 1
	}(msg)
	msg = "Goodbye" // it prints "Hello"
	wg.Wait()
}
```

- 20 go routines - but raising each other and no synchronization

```golang
var wg = sync.WaitGroup{}
var counter = 0

func main(){
	for i:= 0; i < 10; i++{
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello(){
	fmt.Printf("Hello #%v\n", counter)
	wg.Done()
}

func increment(){
	counter ++
	wg.Done()
}
```

- For solving the above issue, use Mutex
 	- Only one routine can access the data at a single time and lock the rest
	- If anything is reading anything can write it into it 
	- Still out of sync but removed the random order

```golang
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RwMutex{} //Creating read-write mutex

func main(){
	runtime.GOMAXPROCS(100)
	for i:= 0; i < 10; i++{
		wg.Add(2)
		go sayHello()
		go increment()
	}
	wg.Wait()
}

func sayHello(){
	m.RLock() // read lock - locking to read counter
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() // releasing the lock with runlock
	wg.Done() 
}

func increment(){
	m.Lock() // write lock - where the varible is mutating
	counter ++
	m.Unlock()
	wg.Done()
}
```

- The solution for the above problem is:
	- Locking everything in single context
	- But it is not working parallely

```golang
var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RwMutex{} //Creating read-write mutex

func main(){
	runtime.GOMAXPROCS(100) // forcing the system to use 100 core 
	for i:= 0; i < 10; i++{
		wg.Add(2)
		m.RLock() // read lock - locking to read counter
		go sayHello()

		m.Lock() // write lock - where the varible is mutating
		go increment()
	}
	wg.Wait()
}

func sayHello(){
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() // releasing the lock with runlock
	wg.Done() 
}

func increment(){
	counter ++
	m.Unlock()
	wg.Done()
}
```

## Channels

- synchronize the data transmission multiple go routine
- create a channel : specify the datatype that flow through the channel

`ch := make(chan int)`

```golang
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() { // the receiving goroutine
		i := <-ch // pulling data from the channel and assigning to variabl i
		fmt.Println(i)
		wg.Done()
	}()

	go func() { // the sending goroutine
		ch <- 42 // putting data into the channel
		wg.Done()
	}()

	wg.Wait()

}
```

- Both are readers and writers at the same time

```golang
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func() {
		i := <-ch      // 2
		fmt.Println(i) // 3
		ch <- 27       // 4
		wg.Done()
	}()

	go func() {
		ch <- 42          // 1
		fmt.Println(<-ch) // 5
		wg.Done()
	}()

	wg.Wait()

}
```

- Making read only and write only channels:

```golang
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int)
	wg.Add(2)

	go func(ch <-chan int) { // receive only - data flow out of channel
		i := <-ch
		fmt.Println(i)
		wg.Done()
	}(ch)

	go func(ch chan<- int) { // send only - data flow into the channel
		ch <- 42
		wg.Done()
	}(ch)

	wg.Wait()
}
```

- Buffered Channel
	- Maintain a data buffer to store the data and can be retrieved
	- Used when the sender and the receiver works at the different frequency

```golang
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50) // the buffer can store 50 integers
	wg.Add(2)
	go func(ch <-chan int) {
		i := <-ch // receiving 42 and prints
		fmt.Println(i)

		i = <-ch // receiving 27 and prints
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan<- int) {
		ch <- 42
		ch <- 27
		wg.Done()
	}(ch)
	wg.Wait()
}
```

- Continue monitoring the data flow

```golang
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for i := range ch { // range for infinite times until the channel stop putting data
			fmt.Println(i)
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27

		close(ch) // close the channel AFTER USE
		wg.Done()
	}(ch)

	wg.Wait()

}
```

```golang
var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)

	go func(ch <-chan int) {
		for  { 
			if i,ok := <- ch; ok{ // i will get the value and ok get the error; if no error, continue execution
				fmt.Println(i)
			}else{
				break
			}
		}
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27

		close(ch) // close the channel AFTER USE
		wg.Done()
	}(ch)

	wg.Wait()

}
```

- Select statement

```golang
const (
	logInfo    = "INFO"
	logWarning = "WARNING"
	logError   = "ERROR"
)

type logEntry struct {
	severity string
	message  string
}

var logCh = make(chan logEntry, 50)
var doneCh = make(chan struct{}) // used as a switch to send close singal

func main() {
	go logger()
	logCh <- logEntry{logInfo, "App is starting"}
	logCh <- logEntry{logInfo, "App is shutting down"}

	time.Sleep(100 * time.Millisecond)

	doneCh <- struct{}{} // passing an empty struct signal to doneCh
}

func logger() {
	for {
		select {
		case entry := <-logCh: // open when theere is anyting from logCh
			fmt.Printf("%v - %v\n", entry.severity, entry.message)
		case <-doneCh: // open when there is some singnal in doneCh
			break
		}
	}
}

```

## Context type 
