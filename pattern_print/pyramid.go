package practice

import "fmt"

func Pattern_pyramid(l int){
	for i:=0;i<l;i++{
		for j:=l-i;j>0;j--{
			fmt.Print(" ")
		}
		for k:=0;k<=i*2;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}