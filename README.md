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

We are going to have to create custom data types(which in theory are just classes)
