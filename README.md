## Go 

- To instantiate the module
    - every project in Go should have a module. It creates go.mod file

```sh
go mod init demo
```
- To run 

```sh
go run main.go
```

- To run , if main is in multiple files

```sh
go run main.go greet.go
```
 
or

```sh
go run .
```

- Go build

```sh
go build .
```

- build with output name

```sh
go build -o hello .
```

## debug mode 

- symbol information and also DWARF information for the debugging purpose 

## Build for the release mode

- This build is for release but not for debugging purpose. You cannot do debugging using tools like dlb or gdb or lldb etc..

```sh
go build -ldflags="-s -w" -o hello .
```

- compile and link 

```sh
go tool compile -o main.o main.go 
go tool link -o demo main.o  
```

## Keywords (2 of 25)

func, package


## Builtin functions (2 of 18)

print,println