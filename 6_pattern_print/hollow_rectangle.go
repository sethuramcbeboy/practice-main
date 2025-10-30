package practice

import "fmt"


func Pattern_hollow_rectangle(n,m int){
		for i:=0;i<n;i++{
			for j:=0;j<m;j++{
				if i== 0 || i==n-1 || j==0 || j== m-1{
					fmt.Print("*")
				}else{
					fmt.Print(" ")
				}
			}
			fmt.Println(" ")
		}
}