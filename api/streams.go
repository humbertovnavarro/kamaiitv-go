package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func validKey(key string) bool {
	return true
}

func validPushSig(sig string) bool {
	return true
}

func addToDB(channel string, server string, ip string) error {
	return nil
}

func removeFromDB(channel string, server string, ip string) error {
	return nil
}

func createVOD(channel string) error {
	return nil
}

func publish(c *gin.Context) {
	fmt.Println("Start RTMP stream")
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(404)
		return
	}
	ip := c.Request.Form.Get("addr")
	channel := c.Request.Form.Get("name")
	server := c.Request.Form.Get("tcurl")
	pushSig := c.Request.Form.Get("push")
	key := c.Request.Form.Get("k")
	if ip == "" || channel == "" || server == "" {
		c.AbortWithStatus(404)
	}
	if validPushSig(pushSig) {
		c.AbortWithStatus(201)
		return
	}
	if validKey(key) {
		err = addToDB(channel, server, ip)
		if err != nil {
			c.AbortWithStatus(201)
			return
		}
		c.AbortWithStatus(500)
		return
	}
	c.AbortWithStatus(404)
}

func done(c *gin.Context) {
	fmt.Println("End RTMP stream")
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	ip := c.Request.Form.Get("addr")
	channel := c.Request.Form.Get("name")
	server := c.Request.Form.Get("tcurl")
	if ip == "" || channel == "" || server == "" {
		c.AbortWithStatus(400)
		return
	}
	err = removeFromDB(channel, server, ip)
	if err != nil {
		c.AbortWithStatus(500)
		return
	}
	err = createVOD(channel)
	if err != nil {
		c.AbortWithStatus(500)
	}
	c.AbortWithStatus(201)
}
