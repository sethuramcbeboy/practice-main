package practice

import "reflect"

func Equal_Inbuild(req1,req2 []int) bool{
	check:=reflect.DeepEqual(req1,req2)
	return check
}

//Without using Inbuild function

func Equal(req1,req2 []int) bool{
	if len(req1)!=len(req2){
		return false
	}
	for i,_:=range req1{
		if req1[i]!=req2[i]{
			return false
		}
	}
	return true
}