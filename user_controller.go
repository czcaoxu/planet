package main

import (
	"net/http"
	"plat/framework"
	"time"
)

func UserLoginController(c *framework.Context) error {
	time.Sleep(time.Second * 10)
	c.JsonWithStatusCode(http.StatusOK, "User login")
	return nil
}
