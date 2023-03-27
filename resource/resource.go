package resource

import (
	"instasafe1/common"
	"instasafe1/middlewares"
	"instasafe1/service"

	"github.com/gin-gonic/gin"
	"github.com/go-chassis/openlog"
)

type Resource struct {
	ser service.Service
}

func (r *Resource) CreateEndUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		openlog.Info("Got a request for CreateEndUser")
		payload, _ := c.Get("Payload")
		httpres := r.ser.CreateEndUser(payload.(map[string]interface{}), "en")
		c.JSON(httpres.Status, httpres)
	}
}

func (r *Resource) CreateTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		openlog.Info("Got a request for CreateTransaction")
		payload, _ := c.Get("Payload")
		httpres := r.ser.CreateTransaction(payload.(map[string]interface{}), "en")
		c.JSON(httpres.Status, httpres)
	}
}

func (r *Resource) GetStatistics() gin.HandlerFunc {
	return func(c *gin.Context) {
		openlog.Info("Got a request for GetStatistics")
		uid := c.Param("uid")
		if filter, ok := c.GetQuery("city"); ok {
			httpres := r.ser.GetStatistics(uid, filter, "en")
			c.JSON(httpres.Status, httpres)
		} else {
			httpres := common.ErrorHandler("110", nil, 0, "en")
			c.JSON(httpres.Status, httpres)
		}
	}
}

func (r *Resource) DeleteAllTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		openlog.Info("Got a request for DeleteAllTransactions")
		httpres := r.ser.DeleteAllTransactions("en")
		c.JSON(httpres.Status, httpres)
	}
}

func (r *Resource) AddLoaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		openlog.Info("Got a request for AddLoaction")
		payload, _ := c.Get("Payload")
		uid := c.Param("uid")
		httpres := r.ser.AddLoaction(uid, payload.(map[string]interface{}), "en")
		c.JSON(httpres.Status, httpres)
	}
}

func (r *Resource) ResetLoaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		openlog.Info("Got a request for ResetLoaction")
		payload, _ := c.Get("Payload")
		uid := c.Param("uid")
		httpres := r.ser.ResetLoaction(uid, payload.(map[string]interface{}), "en")
		c.JSON(httpres.Status, httpres)
	}
}

func (r *Resource) URLRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/user/signUp", middlewares.PayloadValidator(), r.CreateEndUser())
	incomingRoutes.POST("/transactions", middlewares.PayloadValidator(), r.CreateTransaction())
	incomingRoutes.GET("/user/:uid/statistics", r.GetStatistics())
	incomingRoutes.DELETE("/transactions", r.DeleteAllTransactions())
	incomingRoutes.POST("/user/:uid/addLoaction", middlewares.PayloadValidator(), r.AddLoaction())
	incomingRoutes.PUT("/user/:uid/resetLoaction", middlewares.PayloadValidator(), r.ResetLoaction())
}
