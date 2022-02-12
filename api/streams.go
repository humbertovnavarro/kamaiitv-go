package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func publish(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	ip := c.Request.Form.Get("addr")
	channel := c.Request.Form.Get("name")
	server := c.Request.Form.Get("tcurl")
	if ip == "" || channel == "" || server == "" {
		c.AbortWithStatus(401)
	}
	// Insert into database
	// Notify socket io clients
	c.AbortWithStatus(201)
}

func done(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	ip := c.Request.Form.Get("addr")
	channel := c.Request.Form.Get("name")
	server := c.Request.Form.Get("tcurl")
	if ip == "" || channel == "" || server == "" {
		c.AbortWithStatus(401)
	}
	// Delete from database
	// Notify socket io clients
	c.AbortWithStatus(200)
}

func viewerConnect(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	ip := c.Request.Form.Get("addr")
	channel := c.Request.Form.Get("name")
	server := c.Request.Form.Get("tcurl")
	if ip == "" || channel == "" || server == "" {
		c.AbortWithStatus(401)
	}
	fmt.Println("got viewer connect")
	c.AbortWithStatus(200)
}

func viewerDisconnect(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(401)
		return
	}
	ip := c.Request.Form.Get("addr")
	channel := c.Request.Form.Get("name")
	server := c.Request.Form.Get("tcurl")
	if ip == "" || channel == "" || server == "" {
		c.AbortWithStatus(401)
	}
	fmt.Println("got viewer connect")
	c.AbortWithStatus(200)
}
