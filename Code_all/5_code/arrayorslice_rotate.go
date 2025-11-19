package practice

func TwoD_SliceRotate(req [][]int) [][]int {
	l := len(req)
    res:=make([][]int,l)
	for i, _ := range res {
		res[i] = make([]int, l)
	}
	//log.Println("return ........")
	for i := 0; i < l; i++ {
		for j := 0; j < l; j++ {
			res[j][i] = req[i][j]
		}
	}
	return res
}
