# Golang notes

Some files might not compile or run the main as intended, most of the files are intended for reminder of the Golang programming syntax and maybe some best practices.

## Some notes about Go:

* Strings are immutable.
* Go has grabage collector.
* `defers` functions are pushed into a stack.
* Everything that starts with uppercase letter is accesible from other packages, otherwise, only from within the current package.
* For objects, Go uses structs and functions as methods and as constructors
* You can add a name to an imported package, useful when the package name is different of the function/object you are importing.
