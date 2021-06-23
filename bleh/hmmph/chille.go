package main 

import "fmt"

var name string
var amount int 
var quit string


func init(){
	fmt.Println("enter a name")
	fmt.Scanln(&name)
	fmt.Println("Enter the Amount:")
	fmt.Scanln(&amount)
}

func main(){
	fmt.Println("Enter quit to quit and anything else to continue:") 
	fmt.Scanln(&quit)
	trader:= make (map[int]string)
	trade:= make (map[int] int)
	trader[0]= name
	trade[0]= amount 
	for i:=0; quit!="quit"; i++{
		fmt.Println("Enter a name")
		fmt.Scanln(&name)
		fmt.Println("Enter the amount:")
		fmt.Scanln(&amount)
		fmt.Println("Ready to quit now?")
		fmt.Scanln(&quit)
		trader [i]=name
	    trade [i]= amount
	}
	if quit == "quit"{
		for _,k:= range trader{
			fmt.Println(k)
		}
		for _,b:= range trade{
			fmt.Println(b)
		}
	} 
}