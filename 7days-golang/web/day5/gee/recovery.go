package gee

import (
	"fmt"
	"log"
	"net/http"
)

func Recovery() HandlerFunc {
	return func(c *Context) {
		defer func() {
			if err := recover(); err != nil {
				msg := fmt.Sprintf("%s", err)
				log.Printf("%s\n\n", msg)
				c.String(http.StatusInternalServerError, "%s", "Internal Server Error")
			}
		}()
		c.Next()
	}
}
