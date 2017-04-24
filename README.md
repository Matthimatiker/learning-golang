# Learning Golang #

Run a go program:

    go run [path to *.go file] arguments

Example:

    go run mult/mult.go 5

The run command *cannot* be used when passing the path to another ``*.go`` file as argument:

    go run cat/cat.go mult/mult.go
