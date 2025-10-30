package practice

func Intersection(req1,req2 []int) []int{
	s,l:=req1,req2
	m,n,k:=0,0,0
	m,n=len(l),len(s)
	if len(s)>len(l){
		m,n=len(s),len(l)
	}
	for i:=0;i<m;i++{
        for j:=0;j<n;j++{
            if s[i]==l[j]{
				s[k]=l[j]
				k++
			}
		}
	}
	return s[:k]
}