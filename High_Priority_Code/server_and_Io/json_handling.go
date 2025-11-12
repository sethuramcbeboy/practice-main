//Write a program to encode and decode JSON data.

package serverandio

import (
    "encoding/json"
    "fmt"
)

type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func main_json_handling() {
    p := Person{Name: "Alice", Age: 25}
    
    // Encode to JSON
    jsonData, err := json.Marshal(p)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(jsonData))
    
    // Decode from JSON
    var p2 Person
    err = json.Unmarshal(jsonData, &p2)
    if err != nil {
        fmt.Println(err)
        return
    }
}