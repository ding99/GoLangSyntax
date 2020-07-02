package main

import . "fmt"

func main(){
  define()
  del()
}

func define(){
  Println("== define")
  
  var capitals map[string]string
  capitals = make(map[string]string)
  
  capitals["France"] = "Paris"
  capitals["Italy"] = "Roma"
  capitals["Japan"] = "Tokyo"
  capitals["India"] = "New Delhi"
  
  for country := range capitals{ Println(country, "-", capitals[country], " ") }
  
  capital, ok := capitals["American"]
  if(ok){ Println("American Capital ", capital)
  } else { Println("Ameraican Capital does not existing") }
}

func del(){
  Println("== delete")
  
  capitals := map[string]string{"France":"Paris","Italy":"Roma","Japan":"Tokyo","India":"New Delhi"}
  Println("-- original map")
  for country := range capitals{ Print(country, "-", capitals[country],", ") }
  Println()
  
  delete(capitals, "France")
  Println("France Item was deleted")
  
  Println("-- updated map")
  for country := range capitals{ Print(country, "-", capitals[country],", ") }
  Println()
}
