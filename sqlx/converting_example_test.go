package sqlx

import "fmt"

func ExamplePointerToSqlNull() {
	var x int = 7
	fmt.Println(PointerToSqlNull(&x))
	//Output: {7 true}
}
