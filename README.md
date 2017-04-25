# Learning Golang #

Run a go program:

    go run [path to *.go file] arguments

Example:

    go run mult/mult.go 5

The run command *cannot* be used when passing the path to another ``*.go`` file as argument:

    go run cat/cat.go mult/mult.go



## Notices ##

Windows-Binary (exe):

GOOS=

Data types:

:= init with type inference
default int: platform dependent
rune = unicode character
string = byte list
map[key_type]value_type{}

- len(map)
- delete(map, “key”)
- value, exists := map[“key”]

_ = unused var placeholder
Array: [size]type{“a”, “b”}
Slice: []type{}

- newSlice = append(slice, “a”, “b”)
- newSlice = slice[start : end]
    - end optional: slice[start:]
- slice := make(type, len, capacity)
- cap(slice), len(slice)
- capacity is doubled when reached

pass by value

for key, value := rang list {
}
infinite: for { code; }

CLI parameters: os.Args
Unfolding: slice…

Attribute ucfirst => public
          lcfirst => package visibility

*type => pointer
&var => reference to var
(*var) => de-reference var
created by make() => pointer to data

Package Assertions: stretchr/testify

Backticks: multiline string `hello
world`

finds possible problems: go vet . 

Executed on import possible per file:
func init() {
}

Dependency Management?

godep -> Ende 2017

go get ./… => dependencies of current project and sub projects


Übung:

Multiplikationstabelle: n
Files: cat, tac (umgekehrt?), wc wordcount
Package bufio (Scanner)
Key/Value store on CLI



