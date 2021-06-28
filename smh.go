package main 

import "fmt"

var clone Trader

type Trader struct {
	name[] string 
    amount[] float64
    Name string
    Amount float64
}

type trader struct {
    Clone []Trader
}

func main(){
	var name string 
	fmt.Println("Enter a name")
	fmt.Scanln(&name)
    fmt.Println("Enter an amount")
    var amount float64
    fmt.Scanln(&amount)

	temp:=whatever(name)
        for temp!= true {
            Clone= append(name, )
        }

	for temp== false{
        fmt.Println("Enter a name")
        fmt.Scanln(&name)
        fmt.Println("Enter an amount")
        fmt.Scanln(&amount)
    } 
}

func whatever (n string ) bool {
	if n== "quit"{
		return true 
	}
	if n!="quit"{
                clone.name= append(clone.name,Trader{
                    Name: n,
                })

                } 
		return false 
	}

func whatevs( ha float64) bool {
    if ha == "quit"{
        return true 
    }

    if ha!= "quit"{
        clone.amount = append(clone.amount, Trader{
            Amount:ha,
        })
        return false
    }

}