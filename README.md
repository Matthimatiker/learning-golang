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

default: pass by value

_ = unused var placeholder  

*type => pointer
&var => reference to var
(*var) => de-reference var
created by make() => pointer to data

### Data types ###

``:=`` initializes a variable with type inference:
  
    x := 42
    // <=>
    var y int
    y = 42
  
The default int type is platform dependent (int64 vs. int32).  

#### Strings ####

A string is a byte list.  
A ``rune`` represents a unicode character (can contain multiple bytes).  

Backticks are used to assign multiline strings: 

    myString := `hello
    world`
    
``len()`` returns the length in bytes:
    
    lengthInBytes := len(myString)
 
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

#### Arrays ####

Array: [size]type{“a”, “b”}  

#### Slices ####

A slice is a list of variable length.

Declaration: ``[]type{}``

- ``newSlice = append(slice, “a”, “b”)``
- ``newSlice = slice[start : end]``
    - end optional: ``slice[start:]``
- ``slice := make(type, len, capacity)``
- cap(slice), len(slice)
- capacity is doubled when reached


for key, value := range list {
}
infinite: for { code; }

CLI parameters: os.Args
Unfolding: ``slice...``

#### Type Alias ####
          
Create a type alias:

    type Foo string

#### Type Casting ####

Cast to compatible type:

    casted, ok := value.(target type)
    
Type check language construct:
    
    switch casted := value.(type) {
    case string:
        //  casted of type string
    case int:
        // casted of type int
    default:
        // any other type
    }


ToDo: move the following section

finds possible problems: go vet . 

Executed on import possible per file:

    func init() {
    }

Dependency Management?

godep -> Ende 2017

dependencies of current project and sub projects

    go get ./... 
    
#### Visibility ####

Types (and attributes) starting with uppercase letter are *public*:

    type MyType struct {
        Value string
    }
    
Types (and attributes) starting with lowercase character have package visibility:

    type myType struct {
        value string
    }

### Testing ###

Running tests of sub package with verbose output:

    go test -v github.com/matthimatiker/learning-golang/key-value-store

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

Test setup/tearDown pattern:

    setUp := func () (func ()) {
        return func () {
        
        }
    }
    tearDown := setUp()
    defer tearDown()

### Error Handling ###

Handle panics:

    defer func () {
        if p := recover(); p != nil {
            // handle panic
        }
    }()
    panic("oh no!")
    
Prevent errors with static code analyis:
    
    go vet
    
Tool: errcheck (ToDo: find link)



Übung:

Multiplikationstabelle: n
Files: cat, tac (umgekehrt?), wc wordcount
Package bufio (Scanner)
Key/Value store on CLI

KeyValueStore: define interface

