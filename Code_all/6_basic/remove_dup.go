package practice

func Dup_remove(req []int)[]int{
      var s[]int
	  idial:=make(map[int]bool)
	  for _,v:=range req{
		if !idial[v]{
           idial[v]=true
		   s=append(s, v)
		}
	  }
	  return s
}

//Without using map and another slice


func Dup_remove_Single_Slice(req []int)[]int{
     s,index:=req,0
     for i,_:=range s{
         check:=false
		 for j:=0;j<index;j++{
			if s[i]==s[j]{
				check=true
				break
			}
		 }    
		 if check==false{
			s[index]=s[i]
			index++
		 }
	 }
	 return s[:index]
}