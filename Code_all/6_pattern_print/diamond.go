package practice

import "fmt"

func Pattern_daimond(l int){
	for i:=0;i<l;i++{
		for j:=l-i;j>0;j--{
			fmt.Print(" ")
		}
		for k:=0;k<=i*2;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
    for i:=1;i<=l;i++{
		for j:=0;j<i;j++{
			fmt.Print(" ")
		}
		for k:=0;k<=(l-i)*2;k++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}