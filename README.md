hello-go
========

Code for the "[Tour of Go](https://tour.golang.org)"

Build
-----

The code for each chapter is its own file, and because there's a main function in each file, each file is in its own directory.

To build the code file for one chapter, enter the directory and call the `build` command, like this:

```bash
$ cd basics
$ go build
```

This will produce an executable file, which you can execute like this:

```bash
$ ./basics
```

> Note: Some files import external dependencies (libraries that are not part of the standard library).  
> For example, `moretypes.go` imports `"golang.org/x/tour/pic"`.  
> You have to download the dependency first, like this:  
> `go get "golang.org/x/tour/pic"`
