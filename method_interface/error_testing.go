package main

import(
  "fmt"
  "regexp"
)

type AdultAgeException int
func (exp AdultAgeException) Error()string{
  return fmt.Sprintf("Age %v is less than 18", exp)
}


type InvalidNameException struct{
  name string
}
func (exp * InvalidNameException) Error()string{
  return fmt.Sprintf("Name %v can't start with integer", exp.name)
}


type Person struct{
  name string
  age int
}

func test_age(p *Person)(string, error){
  age:=(*p).age
  if age < 18{
    return  "", AdultAgeException(age)
  }
  return "Adult", nil
}


func test_name(p *Person)(string, error){
  name:=(*p).name
  invalid_name:=regexp.MustCompile(`^[0-9]`)
  matched := invalid_name.MatchString(name)
  if matched{
    return "", &InvalidNameException{name}
  }
fmt.Println(name, matched)

  return "Correct Name", nil
}


func main(){
  p := Person{"23arn", 11}

  _, err:= test_age(&p)
  if err!= nil{
    fmt.Println("Error seen with age")
  }
  _, err2:= test_name(&p)
  if err2!= nil{
    fmt.Println("Error seen with name")
  }

}
