package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name string
	Id   string
	Addr string
	Job  string
}

type ErrMsg struct {
	Msg string
}

var userList map[string]User = map[string]User{
	"a": {"shudin", "a", "1254/1 malibag", "Frontend developer"},
	"b": {"shudin1", "b", "uttora", "Android developer"},
	"c": {"shahriar", "c", "dhanmondi", "Backend developer"},
	"d": {"prisoner", "d", "mugda", "Scala developer"},
	"e": {"alex", "e", "USA", "Elixr developer"},
	"f": {"robert", "f", "Unknown", "Rust developer"},
}

// userList["a"] := User{"shudin","a","1254/1 malibag"}
// userList["b"] := User{"shudin","a","1254/1 malibag"}

func GetUserDetails(c echo.Context) error {
	user_id := c.Param("userId")

	data, ok := userList[user_id] // user doesnot exist ok = false
	if !ok {
		msg := ErrMsg{"User does not exist"}
		c.JSON(http.StatusBadRequest, msg)
		return nil
	}
	c.JSON(http.StatusOK, data)
	return nil

}
