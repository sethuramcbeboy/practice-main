package general

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOMAXPROCS:", runtime.GOMAXPROCS(0))
	fmt.Println("NumCPU:", runtime.NumCPU())
}

/////////////////////////////////////////////////////
func update(x *int) {
	*x = 10
}

func main_() {
	a := 5
	update(&a)
	fmt.Println(a)
}

// Output = 10

/////////////////////////////////////////////////////

func main_c() {
	ch := make(chan int)

	ch <- 5  // It won't goes to next line
	fmt.Println(<-ch)
}

// output = dead lock

// ch <- 5
//blocks because there is no goroutine receiving from the channel at that moment.
//Since both send and receive are in the same goroutine, the program waits forever and causes a deadlock.
//////////////////////////////////////////////////////


//You have an Employees Collection with columns of EmployeeCode , EmployeeName, EmailId, Department, Salary that contains duplicate email addresses. 
// Each email may appear more than once due to data entry errors. Write a Mongodb query to delete all duplicate records, keeping only one record per 
// unique email address.

// db.employee.aggregate(
// [{
// $group:{
//     _id :"$EmailId",
//     ids : { "$push":"$_id"},
//     count:{$sum:1}
// }
// ],
// {
// $match: {count:{$gt:1}}
// }
// ]}.forEach(function(doc){
// docs.ids.shifts()
// db.employess.deletemany([_id :{$in:doc.ids}])
// ])