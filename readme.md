# Learning Golang

## Introduction

In Go, every file belongs to a package. package is a keyword used to define a package name. main is a special package name.

```packageMain
package main
```

It tells the Go compiler - "This package is an executable program, not a reusable library". It will produce a runnable program (binary/executable). Go will look for the main() function to start execution.

*Example of other packages*

```examplesOfOtherPackages
package mathutils
package models
package helpers
```

*In simple words*

* package main - Marks the program as executable.

* func main() - The starting point of the program.

* import "fmt" - Imports printing functions :=Declare variable quickly.

* fmt.Println() - Print text to console.

## Internal Memory Management

* Code segment - All functions, constants (all immutable segments) goes to code segment.
* Data segment - All global variables (all mutable segments) goes to data segment.
* Stack - All executable functions goes to stack frame.
* Heap - Escape analysis for closer on stack frame
