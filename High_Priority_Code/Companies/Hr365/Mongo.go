package hr365

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