# go_learning
repo to store go scripts in the progress of learning

## important pointers
- Go is static typed language.
- we need to have entry point in go like java.
- Entry point should be main package and have main function.
- Main function should not return values.
- Most used: fmt package has basic i/o utilities.
- Casting of data types are similar like other languages, but for int <--> string conversions, check the snippet
  `strconv.Itoa(values)` strconv is builtin package
- Default value of all variables are zero. No garbage values.
- If we define a variable we need to use it, otherwise compiler will through error.
- Defining varables:
   ``` var variable_name type = value // we are defining expicitly
       variable_name := value // compiler infers the type```
-  Go has interesting operator called, iota. It generates successive + ve interger constant : 0, 1,2,3...
```package main
	import (
		"fmt")

	// kind of fill in excel
	const (
		_ = iota
		a = 10* iota + 4
		b
		c
	)

	func main(){
		fmt.Println(a,b,c)
	}
	// output is 
	14, 24, 34 
```
- Important pointers for constants
    - we should assign during compile time.
    - shadow works for constant ( we can overwrite)
   
