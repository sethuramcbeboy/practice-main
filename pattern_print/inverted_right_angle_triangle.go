package practice

import "fmt"

func Pattern_inverted_right_angle_triangle(l int){
	for i:=0;i<l;i++{
		for j:=l-i;j>0;j--{
			fmt.Print("*")
		}
		fmt.Println()
	}
}