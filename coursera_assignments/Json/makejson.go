package main

import (
"fmt"
"encoding/json"
)

func main() {
	var mymap map[string]string
        var myname string
	var myaddr string
	fmt.Printf("Enter Name")
	fmt.Scan(&myname)
        fmt.Printf("Enter Address")
        fmt.Scan(&myaddr)

	mymap = map[string]string {"Name": myname, "Address": myaddr}

	myjson, _ := json.Marshal(mymap)
	fmt.Println("The strigified JSON is ", string(myjson))
	fmt.Println("The JSON object is ", myjson)
}
