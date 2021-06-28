package main 

import "fmt"

type Name struct {names[] string}

func main(){
   var name string
   var Anything Name
   fmt.Println("Enter a name")
   fmt.Scanln(&name)
   Anything.names= append(Anything.names, name)
   temp:= whatif(name)
   for temp== false{
      fmt.Println("Enter your name")
      fmt.Scanln(&name)  
      Anything.names = append(Anything.names, name)
      temp=whatif(name)
   }
   for temp!= false{
      fmt.Println(Anything.names)
   }
}

func whatif (n string ) bool {
   neh:= false 
   if n== "quit"{
      neh= true
   }
   return neh  
}