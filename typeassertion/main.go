// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

//	func main(){
//		data := `{"name": "Yash", "age": 21}`
//
// var result map[string]interface{}
// json.Unmarshal([]byte(data), &result)
// fmt.Printf("%v %T",result,result)
// // name := result["name"].(string)   // assert it's a string
// // age := int(result["age"].(float64)) // JSON numbers are float64
// // fmt.Println("Name:", name, "Age:", age)
// }
package main

import (
	"fmt"
	"log"

	"github.com/Yash-Watchguard/typeassertion/Greetings"
)

func main() {
    // Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
    log.SetPrefix("greetings: ")
    log.SetFlags(0)

    // A slice of names.
    names := []string{"Gladys", "Samantha", "Darrin"}

    // Request greeting messages for the names.
    messages, err := Greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    // If no error was returned, print the returned map of
    // messages to the console.
    fmt.Println(messages)
}