package practice

func Swap_variabel(a,b int) (int,int){
	 a=a^b
	 b=a^b
	 a=a^b
	 return a,b
}