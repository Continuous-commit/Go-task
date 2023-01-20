// 課題[5]
// CanAccess()を別の方法で書き換えてください。引数や戻り値は変えずに関数の内部のみ変更すること。
// package main

// import "fmt"

// type Role int

// const (
// 	STUDENT  Role = 1
// 	MENTOR   Role = 2
// 	GRADUATE Role = 3
// 	ADMIN    Role = 4
// )

// type User struct {
// 	Role
// 	Name string
// }

// func (user *User) CanAccess() bool {
// 	if user.Role == MENTOR || user.Role == ADMIN {
// 	 return true
//  }
// return false
// }

// func main() {
// 	user := User{Role: 2, Name: "tarou"}
// 	fmt.Println(user.CanAccess())
// }

package main

import "fmt"

type Role int

const (
	STUDENT  Role = 1
	MENTOR   Role = 2
	GRADUATE Role = 3
	ADMIN    Role = 4
)

type User struct {
	Role
	Name string
}

func (user *User) CanAccess() bool {
	return user.Role == MENTOR || user.Role == ADMIN
}

func main() {
	user := User{Role: 2, Name: "tarou"}
	fmt.Println(user.CanAccess())
}
