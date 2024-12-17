package main

import (
	"net/http"
	"plat/framework"
)

func SubjectGetController(c *framework.Context) error {
	c.Json(http.StatusOK, "Subject Get")
	return nil
}

func SubjectPutController(c *framework.Context) error {
	c.Json(http.StatusOK, "Subject Put")
	return nil
}

func SubjectDelController(c *framework.Context) error {
	c.Json(http.StatusOK, "Subject Delete")
	return nil
}

func SubjectListController(c *framework.Context) error {
	c.Json(http.StatusOK, "Subject List")
	return nil
}

func SubjectNameController(c *framework.Context) error {
	c.Json(http.StatusOK, "Subject name")
	return nil
}
