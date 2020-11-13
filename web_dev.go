package main

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "log"
)

type Page struct { // to store Page data structure
  Title string
  Body []byte
}

// method of type page
func (p *Page) save() error  {
  filename := p.Title + ".txt"
  return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
  filename := title + ".txt"
  body, err := ioutil.ReadFile(filename)
  if err != nil {
    return nil, err
  }
  return &Page{Title: title, Body: body}, err
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
  title := (*r).URL.Path[len("/view/"):]
  p, _ := loadPage(title)
  fmt.Fprintf(w, string(p.Body))
}

func main() {
  // p1 := &Page{Title: "TestPage", Body: []byte("this is sample page")}
  // p1.save()
  //
  // p2, _ := loadPage("TestPage")
  // fmt.Println(string(p2.Body))

  http.HandleFunc("/view/", viewHandler)
  http.ListenAndServe(":8080", nil)
}
