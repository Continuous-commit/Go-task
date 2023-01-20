// 課題[4]
// このコードの実行結果はどうなるか予想してください。
// 実際に実行してなぜこのような挙動になるのかを説明してください。 また、正しいと思われる挙動にするための修正をしてください。
// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	users := []User{
// 		{"tarou", 33},
// 		{"zirou", 22},
// 		{"itirou", 11},
// 	}

// 	for _, user := range users {
// 		user.Age = 44
// 	}

// 	fmt.Printf("%v", users)
// }

package main

import "fmt"

type Men struct{
	Name string
	Age  int
}

func main() {
	users := []Men{
		{"tarou", 33},
		{"zirou", 22},
		{"itirou", 11},
	}

	for i, _ := range users {
		users[i].Age = 44
	}

	fmt.Printf("%v", users) 
}
