package practice

import "reflect"

func Equal_Inbuild(req1,req2 []int) bool{
	s:=req1
    l:=req2
	check:=reflect.DeepEqual(s,l)
	return check
}

//Without using Inbuild function

func Equal(req1,req2 []int) bool{
	s:=req1
	l:=req2
	if len(s)!=len(l){
		return false
	}
	for i,_:=range s{
		if s[i]!=l[i]{
			return false
		}
	}
	return true
}