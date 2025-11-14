package practice

import "fmt"

func Pattern_right_angle_triangle(l int){
	for i:=1;i<=l;i++{
		for j:=0;j<i;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}
}