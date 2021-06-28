package main

import (
	"fmt"
	"strconv"
)

type Trader struct {
	Name   string
	Amount float64
}

// how do i use a nested for loop to check if both name and amounr are true?
func main() {
	var name string
	var amount string
	var Anything []Trader
	

	for  {
		fmt.Println("Enter a name:\n")
		fmt.Scanln(&name)
		if name=="quit"{
			break 
		}

		fmt.Println("Enter an amount\n")
		fmt.Scanln(&amount) 
		if amount == "quit"{
			break
		}
		
		f, err := strconv.ParseFloat(amount, 64)

		if err!=nil{
			fmt.Println("you entered an incorrect number")
			break
		} 

		Anything = append(Anything, Trader{
			Name: name,
			Amount:f,
			
		})
	
	}

	for i, v := range Anything{
		fmt.Printf("%d) %s %g\n", i+1,v.Name,v.Amount)
		
	}
	
}

