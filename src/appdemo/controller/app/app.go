package app

import (
	controller "appdemo/controller/base"
)

// Controller user controller
type Controller struct {
	controller.Controller
}

// Get ...
func (c *Controller) Get() {
	c.Post()
	return
}

// Post handler get request
func (c *Controller) Post() {
	switch c.Ctx.Input.Param(":do") {
	case "test1":
		c.test1() // the succ demo
	case "test2":
		c.test2() // the fail demo
	default:
		c.test2()
	}
	return
}
