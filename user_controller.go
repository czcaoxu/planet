package main

import (
	"net/http"
	"plat/framework"
)

func UserLoginController(c *framework.Context) error {
	c.Json(http.StatusOK, "User login")
	return nil
}
