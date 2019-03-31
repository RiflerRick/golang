_go cli_

```bash
go build # compiles a program

go run # used for compiling and executing a file

go fmt # formats all the code in each file in the current directory

go install # compiles and installs a package

go get # downloads the raw source code of someone else's package

go test # run and execute any tests
```

a go package is a collection of common source code files. The very first line of each file must declare the package that it belongs to.

inside of go there are 2 types of packages: an executable type and a reusable type.
An executable type of package is the one that when compiled spits out an actual executable of runnable type of file. Executable packages are used for actually doing something.

Reusable packages are code dependencies or libraries.

Its actually the name of the package you use that determines whether you are making an executable or a reusable type package.

anytime we are using `package main` we are using an executable package. Anything else and we are going to create a reusable package.

anytime we create an executable package, it must always have the function main inside itself.

`import "fmt"`
the import statement is used to give our package access to some code that is written inside another package. So its basically the same as having the import statement of python for instance.

`fmt` is the name of the standard library package that is included by default with the go programming language.

Inside of go there are essentially 2 ways of defining a list like datastructure. The first one is called a good old array which can be of a fixed length only.
The second one is called a slice that can be of variable length and has other properties. In slices however we can only store variables of the same type.

Object Oriented approach vs Go approach

Go is not an object oriented programming language and so there is no idea of classes in go
In case of the deck of cards example that we are dealing with here in this project, an object oriented approach might be the following:

There would be a Deck class with the following data members:

- cards: string type storing all the cards

And the following methods:

- print: print the deck of cards
- shuffle: suffle the deck of cards
- saveToFile: save the entire deck to a file

However in case of Go, the approach would be a bit different:

We are going to have to create custom data types(which in theory are just classes). However unlike classes in which we can have our own data members, in case of go we literally extend characteristics of a base type.

For instance if we use `type deck []string` we are saying that we have a special type deck which extends a type string

Analogous to methods we have receiver functions in go

One important thing to note when it comes to go packages is that if we have multiple files in the same folder, we can use functions declared in some other file. When we are using `go run`, all we have to do is simply include that file as part of the `go run` command.

Slicing in go works pretty much the same way as in python.
For a slice in go if we want to get a portion of it
we would do `a[1:3]` which would give us back all elements from index 1 to index 2, it would not include 3

also `a[:3]` would be from 0 to 2
also `a[1:]` would be 1 till the end

Packages for interacting with files `ioutil`

`byte slice`: byte slice is basically a string.
We can type cast strings to byte slices

```go
greeting := "Hi there!"
fmt.Println([]byte(greeting))
```

## Pointers in golang

```go
func main() {
    jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000, // this comma is required
		}, //even if it is the last value in the struct, the comma is required
	}

	jim.updateNamePassByValue("jimmy") // when we call this function, from the struct actually its a pass by value
	jimPointer := &jim
	jim.print()
	jimPointer.updateNamePassByRef("jimmy")
	jim.print()
}

//receiver functions for structs
func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p person) updateNamePassByValue(newFirstName string) {
	p.firstName = newFirstName
}

func (p *person) updateNamePassByRef(newFirstName string) { // understand that p is here a pointer type
    // object
	(*p).firstName = newFirstName
}
```

The `&` operator works in the same way as it works in c. It gives us the memory address of a variable
The `*` operator also works in the same way as it works in c. It gives us the variable that is pointed to by the memory address

When we are calling the function `updateNamePassByRef(newFirstName string)`, it is possible to actually
call the function simply using the object itself. for instance `jim.updateNamePassByRef(<string>)`. go
would figure out that we are dealing with a pointer in the actual function and go will automatically
do the needful to give you a pointer inside the function.

For slices if we pass a slice to a function it is a pass by reference and not a pass by value.

Strings, floats, ints, structs, bools are all **passed by value**

## Slices

When we actually instantiate and initialize a slice, go is actually creating 2 separate datastructures for us. The first one is a slice and the second one is an array.
`The slice has a pointer, a capacity number and a length`. The length is how many elements are there in the slice. The capacity is how many elements the slice can hold and the ptr in the pointer to the head of the actual array
Whenever we create a slice, the slice and the array are existing in 2 separate locations in memory. Go is a **pass by value** language and so when we pass the slice over to the function, the slice is still copied however the new copied slice still has the pointer to the head as the pointer to the same array that was pointed to by the old slice. This makes the difference and this is why it appears that we are passing by reference.

This kind of behaviour is prevalent in many other data structures as well. These `types` are called **reference types**.

![value types and reference types](https://raw.githubusercontent.com/RiflerRick/golang/master/static/Screenshot%20from%202019-03-31%2021-11-25.png)

## Maps

Maps in go are basically like python dictionaries. One important thing is that in maps, all the keys must be of the same type and all the values must be of the same type.
They can be declared statically in the following way:

```go
static initialization of maps
colors := map[string]string {
    "red": "<hex code of red>",
    "blue": "<hex code of blue>",
}

var colors map[string]string // both keys and values are of type string

// another way of doing the same thing
colors := make(map[string]string)

colors["green"] = "<hex code for white>" // modifying an existing map

delete(colors, "green") // delete the "green" key

// iterating over maps
for color, hex := range colors {
    fmt.Println("Hex code for the given color", color, "is", hex)
}

```
