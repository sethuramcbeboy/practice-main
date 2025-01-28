package practice

import "fmt"

func IsPalindrome(x int) bool {
    a,sum,v:=0,0,0
	v=x
	if x<0{
        return false
    }
    for{
        if x>0{
        a=x%10
        sum=sum*10+a
        x/=10
		fmt.Println(sum,x)
        }else{
            break
        }
    }
    if v==sum{
        return true
    }else{
        return false
    }
}