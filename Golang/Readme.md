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