package practice

func Reverse(s []int)[]int{
	 l:=len(s)
	 for i:=0;i<(l/2);i++{
         j:=l-i-1 
		s[i],s[j]=s[j],s[i]
	 }
	 return s
}