package practice

import "fmt"

func Pattern_print(r,c int){
	for i:=0;i<r;i++{
		for j:=0;j<c;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}