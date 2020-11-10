// to work with custom packages create a module, its collection of packages
// in the project root directory go mod init <module_name> // same as name of project folder
// all custom packages needs to in subdirectories.
// all export functions needs to start with caps.
// import "module_name/subdirectory"

package main

import(
    b "go_learning/test_pack"
)
func main()  {
  b.Display("sample data")
}
