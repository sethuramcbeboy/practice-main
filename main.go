package main

import (
	// run "Practice/code"
	// "fmt"
	//run2 "Practice/concurrency"

	leetcode "Practice/Leetcode"
	"log"
	//mapcode "Practice/Mapcode"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	log.Println("--------------------------Execution Started--------------------")
	//run.Odd()                   	//1
	//run.Merge()             		//2
	//run.Max()						//3
	//run.Sort()					//4

	// var s[]int					//5
	// s=append(s, 1,1,2,3,3,4)
	// count:=run.Occourance(s)
	// fmt.Println(count)

	// var s[]int					//6
	// ctr:=2
	// s=append(s,1,2,3,4)
	// rotate:=run.Rotate(s,ctr)
	// fmt.Println(rotate)

	// var s[]int 					//7
	// s=append(s, 1,2,3,4,5,6,7,8)
	// res:=run.Reverse(s)
	// fmt.Println(res)

	// s:=[]int{1,2,3,5}			//8
	// l:=[]int{1,2,3,5}
	// res:=run.Equal(s,l)
	// fmt.Println(res)

	// s,l:=[]int{3,2,4,6},[]int{1,2,3,5}    //9
	// res:=run.Intersection(s,l)
	// fmt.Println(res)

	// s:=[]int{1,2,3,4,4,5,5,6,6,6,9,10,10,10}    //10
	// res:=run.Dup_remove_Single_Slice(s)
	// fmt.Println(res)

	// arr:=[10]int{1,2,3,4,5,6,7,8,9,10}      //11
	// res:=run.Slice_arr_operation(arr)
	// fmt.Println(res)

	// var s[]int
	// fmt.Scanln(&s)
	// res:=run.Slice_of_slice(s)
	// fmt.Println(res)                        //12

	// data, err := ioutil.ReadFile("/home/sethuram/practice/PrivateKey_327003.pem")
	// if err != nil {
	// 	fmt.Println("Error reading file:", err)
	// 	return
	// }

	// // Convert file content to string
	// fileContent := string(data)

	// // Convert newlines to \n
	// singleLineContent := strings.ReplaceAll(fileContent, "\n", "\\n")

	// // Print the result
	// fmt.Println(singleLineContent)        //13

	// var s string
	// fmt.Scanln(&s)
	// res := run.Comma_string_separate(s)
	// fmt.Println("...........",res)       //14

	// s:="apple banna apple orange apple banna"
	// res:=run.UniqueSortedWords(s)
	// fmt.Println(res)						//15

	// var n int
	// fmt.Scan(&n)
	// matrix := make([][]int, n)
	// for i, _ := range matrix {
	// 	matrix[i] = make([]int, n)
	// }
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		fmt.Scan(&matrix[i][j])
	// 	}
	// }
	// res := run.Diagonal_sum(matrix)
	// fmt.Println(res)                    //16

	// var n int
	// fmt.Scan(&n)
	// mat:=make([][]int,n)
	// for i,_ :=range mat{
	//     mat[i]=make([]int, n)
	// }
	// for i:=0;i<n;i++{
	// 	for j:=0;j<n;j++{
	// 		fmt.Scan(&mat[i][j])
	// 	}
	// }
	// res:=run.TwoD_SliceRotate(mat)
	// fmt.Println(res)						//17

	// var n int
	// fmt.Scan(&n)
	// arr:=make([]int,n)
	// for i:=0;i<n;i++{
	// 	fmt.Scan(&arr[i])
	// }
	// res:=run.Longest_increase_in_arr_or_slice(arr)
	// fmt.Println(res)						//18
	//run.Queue_2_stack_impl() 				//19
	// var v int
	// fmt.Scan(&v)
	// res:=run.IsPalindrome(v)
	// fmt.Println(res)						//20
	//var s string
	// fmt.Scan(&s)
	// res:=run.RomanToInt(s)
	// fmt.Println(res)						//21
	// s := []string{"sts", "lds"}
	// res := run.LongestCommonPrefix(s)
	// fmt.Println(res)						//22
	// var s string
	// fmt.Scan(&s)
	// res:=run.Valid_Paranthesis(s)
	// fmt.Println(res)						//23
	//run.Linked_list()						//24
	//run.Goroutine("new")					//25
	//run.Call_list()						//26
	// list:=run.Create_Linked_list()
	// var v int
	// fmt.Scan(&v)
	// res:=run.LinkedList_insert_at_end(v,list)
	// run.Print_list(res)					//27
	// list:=run.Create_Linked_list()
	// var v int
	// fmt.Scan(&v)
	// res:=run.LinkedList_insert_at_begining(v,list)
	// run.Print_list(res)					//28
	// list:=run.Create_Linked_list()
	// var v int
	// fmt.Scan(&v)
	// res:=run.LinkedList_delete_node(v,list)
	// run.Print_list(res)
	//run2.Go_routines()
	//run2.Pipeline_pattern()
	// var rows, columns, s int
	// fmt.Scanln(&rows, &columns, &s)

	// // Initialize the 2D slice
	// arr := make([][]int, rows)
	// for i := 0; i < rows; i++ {
	// 	arr[i] = make([]int, columns) // Initialize each row slice
	// }

	// // Read the values into the 2D slice
	// for i := 0; i < rows; i++ {
	// 	for j := 0; j < columns; j++ {
	// 		fmt.Scan(&arr[i][j])
	// 	}
	// }

	// // Call the function (assuming it's defined correctly elsewhere)
	// res := run.Circle_avoid_remaining_value_max_add(s, arr)

	// fmt.Println("Result:", res)
	leetcode.Leet_1()
	//mapcode.Map_int()
	// i := 10

	// switch {
	// case i > 10:
	// 	fmt.Println("new")
	// default:
	// 	fmt.Println("not greater")
	// }
}
