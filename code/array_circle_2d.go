package practice

/*
10 20 30 40 
11 22 33 44

if p = 10 then we need to avoid 20,11,22 circle line value, remaining value we need to find max and return max+p
*/

func Circle_avoid_remaining_value_max_add(p int,arr [][]int) int{
	rows:=len(arr)
	columns:=len(arr[0])
	r,c:=0,0

	//first finding p values index
	for i:=0;i<rows;i++{
		for j:=0;j<columns;j++{
			if arr[i][j]==p{
				r,c=i,j	
				break		
			}
		}
	}
	
	//finding the max value in array except p value and its adjacent values
	var max int
	t:=Get(arr,r,c,rows,columns)
	for i:=0;i<rows;i++{
		for j:=0;j<columns;j++{
			if Check(arr[i][j],t){
				if arr[i][j]>max{
					max=arr[i][j]
				}
			}
		}
	}
	return max+p
}


func Get(arr [][]int,r int,c int,rows int,columns int) []int{
	var t []int		
	if c-0>0{							//left
		t=append(t, arr[r][c-1])
	}
	t=append(t, arr[r][c])				//current (p)
	if c+1 <= columns{
		t = append(t, arr[r][c+1])		//right
	}
	if r+1<=rows && c-1>=0{
		t=append(t, arr[r+1][c-1])		//below left
	}
	t=append(t, arr[r+1][c])			//below
	if r+1<=rows && c+1<=columns{
		t = append(t, arr[r+1][c+1])	//below right
	}
	if r-1 >= rows && c-1 >=0{
		t = append(t, arr[r-1][c-1])	//above left
	}
	if r-1>=rows{
		t=append(t, arr[r-1][c])		//above of p
	}
	if r-1>=rows && c+1 <= columns{
		t=append(t, arr[r-1][c+1])
	}
	return t
}

func Check(a int,t []int) bool{
	for _,v:=range t{
		if v==a{
			return false
		}
	}
	return true
}



/*
10 20 30 40 
11 22 33 44

if 10 is given as p we need avoid 20,11,22 circle line value remainder value we need to find max and return max+p
*/