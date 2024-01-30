package main

import "fmt"

type nspevent struct {

 date string
 name string
 contact int
 organizingteam
}

type organizingteam struct {
	names []string
	 fun_contact int

}

func main() {

	nsp1 := nspevent{
		date : "22jan2024",
		name : "rakesh",
		contact : 1234,
        organizingteam: organizingteam{
			names :[]string{"one","two"},
            fun_contact: 5678,
        },
        
	}
	fmt.Println(nsp1.date , nsp1.name ,nsp1.organizingteam.fun_contact, nsp1.organizingteam.names)

	fmt.Println(nsp1.date , nsp1.name ,nsp1.organizingteam)
	
	fmt.Println("Structures")

}