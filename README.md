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
- Array, Its homogeneous, static and contiguous memory blocks.
  ```
  	grades:=[3]int {90, 89, 78} // array assignment 
	
	fmt.Println(grades)    // similar to python
	fmt.Println(grades[2]) // indexing 
	
	var stud [3]string // only define 
	fmt.Println(stud)
	stud[2] = "gupau"
	fmt.Println(stud)
	fmt.Println(len(stud))
	```
 - Slice, similar to python list it can grow  it size but homogeneous.
   ```
        slc:= make([]int, 1,10)  // creates a slice
	slc:=[]int{1,2,3} // slice  another way
	fmt.Println(slc, len(slc), cap(slc)) // as its dynamic its has a inintial capacity	
	
	slc = append(slc, 10)  // append opeations
	fmt.Println(slc, len(slc), cap(slc))	
	slc = append(slc, []int{12,89,70}...)  // spreading operator
	fmt.Println(slc, len(slc), cap(slc))	
	```
 - Maps, they are similar to python dictionaries.
    - slice can not be map index.
    - map should end with comma.
    - maps are not ordered like python dictionary.
    - delete using builtin function.
    - for index not present it returns zero.
    ```	// initializtion
	sample_map := map[string]int{
	"1": 1,
	"two": 2,
	} 
	```
     - Dealing with empty values
     ```
     _, ispresent := sample_map2["gupai"] // to check if key is present or not.
	fmt.Println(ispresent)	```
 - Interesting switch pointers
     -  No need for break in cases, fallthrough is not compulsory
     -  switch can be used for type check as well 
     ```switch(i){
	case 1,5,10:
		fmt.Println("Ola1")
	case 3:
		fmt.Println("Ola2")
	default:
		fmt.Println("default")
	}```
 - Struct:
    - Acts as pseudo class.
    - Struct should end with comma.
      ```
      type Doctor struct {
		id int
		name string
		grades [5]int
		}
		func main(){
			adoc := Doctor{
				id: 90,
				name: "pupau",
				grades: [5]int {32, 42, 45, 35, 40},
			}

		 fmt.Println(adoc)
		 fmt.Println(adoc.grades)```
    - Exact inheritance is not present in go, we have composition(Works cool)
      ```
      type Animal struct {
			Name string
			Origin string
		}

		type Bird struct {
			Animal
			speed float32
			CanFly bool
		}

		b := Bird{}
		b.Name="Oa"
		b.Origin="India"
		b.speed=89.78
		b.CanFly=true

		fmt.Println(b)
      ```
 - Pointers:
     - pointer stores the address of the types variable
     - & is called aaddress off operator
     - * is for dereference. Another usage: pointer variable declaration.
     - IMP: pointer arithmetic is not allowed in go directly , use unsafe package for these operations.
     - null pointers are assigned value <nil> unlike c or cpp they have null pointers.
	
```
func main(){
	var a int = 42
	var a_ptr *int  // delaire a pointer
	a_ptr = &a  // & returns the address
	fmt.Println(a,a_ptr, *a_ptr, &a_ptr)  // * returns the value of that address
	*a_ptr ++ // increasing the value 
	fmt.Println(a,a_ptr, *a_ptr, &a_ptr)  // * returns the value of that address
}
```
- Defer:
    - statements with defer keyword, before returning a function: defer statement are executed.
    - defer is perfect for cleanup operations.
    - defer is evaluated based on the values at the time when defer is called not during execution.
    - Multiple defers are excuted in LILO
      ```
      func main(){
		fmt.Println(1)
		defer fmt.Println(2)
		fmt.Println(3)
		}
		//output : 1, 3 and 2
      ```
- Additinonal new() is used for creating empty object new(struct) returns object of struct.
- nil in place of None or null
- Functions:	
    - passing a pointer is faster and memory efficient in between function calls
    - maps and slices act as pass by reference in function calls
    - returning a variable as pointer from a function ? go promotes the variable into heap from stack memory.
    - returning multiple values by comma :  return a, b
    - we have method in go, its defined with the context of struct.
    
      ```func main(){
      gr := greeter { 
		name: "French",
		greetings: "OLA",
		}
		gr.greet() 

		}
		func (g greeter) greet(){ // pass by value here 
		fmt.Println(g.name, g.greetings)
		}
      ```
- Sample function
    
	 ```  func test(name string) string {  return name+"---" }```
- Interface
    - Go does not have classes. However, you can define methods on types. 
    - You can declare a method on non-struct types, too. 
    - You cannot declare a method with a receiver whose type is defined in 
        - another package 
        - built-in types such as int
    - You can declare methods with pointer receivers. 
    - Interfaces contains method declaration. 
    - Methods implicitly implements an interface via type context
    - Context of method can be of any type(user defined).
    - When we do interface_type.method() it find for the implementation with defined method with the context.
    - We ability to embedded interfaces like struct, combining interfaces.
    - Examples:
      ```	package main
		import (
			"fmt"
		)

		type Writer interface {
			write(data string) (string, error)
		}

		type StdWriter struct {
			parameter string
		}
		// implementing methods 
		func (st StdWriter)write(value string) (string, error){

			return st.parameter + value, nil
		}

		func main(){
			fmt.Println("Hellow")
			var wr Writer = StdWriter{}
			fmt.Println(wr.write("Ola"))
		}
      ```
 - Packages and Modules:
     - Packages are collection of functions, methods, etcs.
     - Modules are collection of packages.
     - A go project is of strucure 
         - project_root/
         - main.go
         - utility_packages/
         - test/
     - To use any custom packages, we need to create module, in project root
         `go mod init <module_name its mostly the root folder name>`
     - In main we can import packages in utility_packages by 
         `import ("module_name/utility_packages")`
     - To import a custom package, package must be in a subfolder.
     - go modules makes the dependency managent easier.
     
  - Concurrency in Go:
      - Go has builtin high level concurrency.
      - Threads are smallest unit of processing tasks in OS.
      - Threads --> shared memory while process have seperate memory.
      - A process has multiple threads.
      - Linux treats threads and process as same entity, all are tasks.
      - With fork() --> min sharing of resources
      - With pthread_create() --> max sharing of resources.
      - Drawback of threads:
          - Every threads needs a copy of stack(~1MB each). When thread number increases this memory increases drastically.
	  - A thread writes a lot of registers, increasing theads is an overhead.
	  - Creating and teardown required OS calls. Resuing is better.
      - Go comes with goroutines which work in virtual space. OS sees single user requesting.    
      - Global variable GOMAXPROCS --> max # threads to be created by os. For best performance it must be equal to # cores.
      - Channel:
          - Channel are mode for communication between goroutines.
          - Deafult: Sends and receives block until the other side is ready. This 
	    allows goroutines to synchronize without explicit locks or condition variables.
          - we can create buffered channel. when buffer is full receive is blocked.
          - v, ok := <-ch, ok is false if there are no more values to receive and the channel is closed.
          - The loop for i := range c receives values from the channel repeatedly until it is closed.
          - Sender can close the channel, never the receiver. Sending on a closed channel will cause a panic.
          - Channels aren't like files; you don't usually need to close them. Closing is only necessary when the 
            receiver must be told there are no more values coming, such as to terminate a range loop.
       - Select block:
           - A select blocks until one of its cases can run,
           - then it executes that case. It chooses one at random if multiple are ready.
       - mutual exclusion-> We make the go routines mutually exclusive.
           - We achive this my sync.mutex package.
           - we have Two methods:
               - Lock
               - Unlock
           - We can define a block of code to be executed in mutual exclusion by surrounding 
           - it with a call to Lock and Unlock.
           - We can also use defer to ensure the mutex will be unlocked.
