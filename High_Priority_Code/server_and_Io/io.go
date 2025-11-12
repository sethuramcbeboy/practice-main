//Demonstrate how to read and write files in Go.

package serverandio

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// Writing to a file
	data := []byte("Hello, Go!")
	ioutil.WriteFile("example.txt", data, 0644) // 0644 -> is file permission

	// Reading from a file
	content, _ := ioutil.ReadFile("example.txt")
	fmt.Println(string(content))
}
