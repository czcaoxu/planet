package middleware

import (
	"log"
	"plat/framework"
	"time"
)

func Cost() framework.ControllerHandler {
	return func(c *framework.Context) error {
		log.Println("[Middleware] Cost")
		startTime := time.Now()

		c.Next()

		log.Printf("[Middleware] Cost api uri %s, cost: %v\n", c.GetRequest().RequestURI, time.Since(startTime))
		return nil
	}
}
