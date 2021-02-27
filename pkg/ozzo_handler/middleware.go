// Package accesslog provides a middleware that records every RESTful API call in a log message.
package ozzo_handler

import (
	routing "github.com/go-ozzo/ozzo-routing/v2"
)

// Handler returns a middleware with setted HTTP header.
func SetHeader(key string, value string) routing.Handler {
	return func(c *routing.Context) error {
		c.Response.Header().Set(key, value)
		return nil
	}
}
