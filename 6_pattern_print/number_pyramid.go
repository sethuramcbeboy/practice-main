package practice

import "fmt"

func Number_pyramid(l int){
	for i:=0;i<=l;i++{
		for k:=i;k<l;k++{
			fmt.Print(" ")
		}
		for j:=1;j<=i;j++{
			fmt.Print(j," ")
		}
		fmt.Println()
	}
}