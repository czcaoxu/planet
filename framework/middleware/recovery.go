package middleware

import (
	"log"
	"plat/framework"
)

func Recovery() framework.ControllerHandler {
	return func(c *framework.Context) error {
		log.Println("[Middleware] Recovery")
		defer func() {
			if err := recover(); err != nil {
				log.Fatalf("Panic %v\n", err)
				c.JsonWithStatusCode(500, err)
			}
		}()

		c.Next()

		return nil
	}
}
