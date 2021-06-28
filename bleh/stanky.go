package main 

import "fmt"

type name struct {names[] string}

func main(){
   var Name string
   var na name
   fmt.Println("Enter a name")
   fmt.Scanln(&Name)
   na.names = append(na.names, Name)
   temp:= whatif(Name)
   for temp== false{
      fmt.Println("Enter your name")
      fmt.Scanln(&Name)  
      na.names = append(na.names, Name)
   }
   for temp== true{
      fmt.Println(na.names)
   }
}

func whatif (n string ) bool {
   neh:= false 
   if n== "quit"{
      neh= true
   }
   return neh
}