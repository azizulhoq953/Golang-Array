package main
import "fmt"

func main(){

//primitive data type
//rune,byte=alias

//maps
// Composite Data type
//key value

x := make(map[string]string)

x["name"] = "Azizul Hoq"

x["height"] = "5.7"

//delete Function
delete(x,"height")

fmt.Println(x["name"])


}