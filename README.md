# Learning Golang #

Run a go program:

    go run [path to *.go file] arguments

Example:

    go run mult/mult.go 5

The run command *cannot* be used when passing the path to another ``*.go`` file as argument:

    go run cat/cat.go mult/mult.go

Running tests of sub package with verbose output:

    go test -v github.com/matthimatiker/learning-golang/key-value-store


## Notices ##

Windows-Binary (exe):

GOOS=

### Data types ###

:= init with type inference  
default int: platform dependent  
rune = unicode character  
string = byte list 
 
#### Maps ####

Declaration:

    map[key_type]value_type{}  

Common operations:

- len(map)
- delete(map, “key”)
- value, exists := map[“key”]

Iteration over map:

    for key, value := range myMap {
    }

Order of map iteration is random:

- https://nathanleclaire.com/blog/2014/04/27/a-surprising-feature-of-golang-that-colored-me-impressed/

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

Backticks: multiline string `hello
world`

finds possible problems: go vet . 

Executed on import possible per file:

    func init() {
    }

Dependency Management?

godep -> Ende 2017

dependencies of current project and sub projects

    go get ./... 

### Testing ###

Improved assertions are available in the following package:

    stretchr/testify

Generate test coverage from cover file:

    go tool cover -html=cover.out
    
In an test, ``TestMain`` can be used to define setup and tear down for the *whole* test case:
    
    func TestMain(m *testing.M) {
        // setup
    
        // *all* tests are running here:
        returnCode := m.Run()
    
        // tear down
    
        os.Exit(returnCode)
    }
    
This construct is comparable to ``setUpClass`` and ``tearDownClass`` in PHPUnit.



Übung:

Multiplikationstabelle: n
Files: cat, tac (umgekehrt?), wc wordcount
Package bufio (Scanner)
Key/Value store on CLI



