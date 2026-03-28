package sunware

import (
	"log"
	//"net/http"
)

type Resp struct {
	Name     string
	Password string
	Age      int
	Address  *Address
}

type Address struct {
	Addressline1 string
	Addressline2 string
	City         string
	Country      string
	Zip          string
}

func Getuserapi() (Resp, error) {
	// ctr := context.WithTimeout(5 * time.Minute)
	// client, err := db.Connect(config.URL)
	// if err != nil {
	// 	return nil, err
	// }
	// defer close(client)
	// res := client.findAll(&Resp{}, ctx)
	// return res
	return Resp{},nil
}

func main() {
	log.Println("Starting the application")
	//http.HandleFunc("/app/Getuser", Getuserapi())
	//http.ListenAndServe(":8080")
}
