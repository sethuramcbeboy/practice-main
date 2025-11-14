package practice

import "fmt"

func Flayord_triangle(l int){
	v:=1
	for i:=0;i<l;i++{
		// for k:=0;k<l-i;k++{
		// 	fmt.Print(" ")
		// }
		for j:=0;j<=i;j++{
			fmt.Print(v," ")
			v++
		}
		fmt.Println()
	}
}