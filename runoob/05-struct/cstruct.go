package main

import . "fmt"

type Books struct {
  title string
  author string
  subject string
  book_id int
}

func main(){
  Println(Books{"Go Language","www.runoob.com","Go Language Tutorial", 6495407})
  Println(Books{title:"Go Language",author:"www.runoob.com",subject:"Go Language Tutorial", book_id:6495407})
  Println(Books{title:"Go Language",author:"www.runoob.com"})
}