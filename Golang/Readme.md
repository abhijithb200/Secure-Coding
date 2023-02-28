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