package handlers

import (
	"github.com/labstack/echo/v4"
)

type User struct {
	Name string 
	Id string 
	Addr string 
	Job string 
}

var userList map[string]User

userList["a"] = User{"shudin","a","1254/1 malibag"}
userList["b"] = User{"shudin","a","1254/1 malibag"}

//exported function
func GetUserDetails(c echo.Context) error {
	user_id := c.Param("userId")
}
