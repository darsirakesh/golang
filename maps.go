package main

import "fmt"

func main() {

	family := make(map[string][]string)
	family["SR"] = []string{"7750","7950" , "7450"}

	fmt.Println("map:", family)

	family["SAS"] = []string{"sasx", "sasm" ,"sasdxp"}

	fmt.Println("map:", family)
	fmt.Println("now iterating through the map")

	for k, v := range family { 
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	fmt.Println("finding the node")

	result, ok := family["SR"]

	if(ok){
		for i := 0 ; i< len(result) ; i++ {
			if(result[i] == "7250"){
				fmt.Println("node is not present in the family")
			}else {
				result = append(result , "7250")
				  family["SR"] = result
				  break
			  }
	     }
		 
    } 

	for k, v := range family { 
		fmt.Printf("key[%s] value[%s]\n", k, v)
	}

	fmt.Println("deleting the node")

	result2, ok := family["SR"]

	if(ok){
		for k := 0 ; k< len(result2) ; k++ {
			if(result2[k] == "7450"){
				result2 = append(result2[:k], result2[k+1:]...)
				family["SR"] = result2
				break
			}
			
    } 
		}

		for k, v := range family { 
			fmt.Printf("key[%s] value[%s]\n", k, v)
		}


}