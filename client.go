package main
import (
	"./db/user"

	"fmt"
	"reflect"
)

func main() {
	fmt.Println(user_db.GetUUIDById(1))
}