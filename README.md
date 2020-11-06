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
   ` var variable_name type = value // we are defining expicitly 
   ` variable_name := value // compiler infers the type
   
