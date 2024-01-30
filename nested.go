package main

import "fmt"

type nspevent struct {

 date string
 name string
 contact int
members organizingteam
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
        members: organizingteam{
			names :[]string{"one","two"},
            fun_contact: 5678,
        },
        
	}
	fmt.Println(nsp1.date , nsp1.name ,nsp1.members.fun_contact, nsp1.members.names)
	
	fmt.Println("Structures")

}