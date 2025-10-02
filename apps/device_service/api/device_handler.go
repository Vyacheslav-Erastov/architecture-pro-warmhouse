package api

import (
	"context"
	"device_service/models"
	"device_service/mq"
	"device_service/service"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type DeviceHandler struct {
	service  *service.DeviceService
	producer *mq.Producer
}

func NewDeviceHandler(service *service.DeviceService, producer *mq.Producer) *DeviceHandler {
	return &DeviceHandler{service: service, producer: producer}
}

func (h *DeviceHandler) RegisterRoutes(r *gin.Engine) {
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.GET("/devices", h.getDevices)
	r.GET("/devices/:id", h.getDeviceByID)
	r.POST("/devices", h.createDevice)
	r.PATCH("/devices/:id/status", h.updateStatus)
	r.POST("/devices/:id/commands", h.sendCommand)
}

func (h *DeviceHandler) getDevices(c *gin.Context) {
	devices, err := h.service.GetDevices(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, devices)
}

func (h *DeviceHandler) createDevice(c *gin.Context) {
	var d models.DeviceCreate
	if err := c.BindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, err := h.service.CreateDevice(context.Background(), d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, id)
}

// Получить устройство по ID
func (h *DeviceHandler) getDeviceByID(c *gin.Context) {
	id := c.Param("id")
	d, err := h.service.GetDevice(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "device not found"})
		return
	}
	c.JSON(http.StatusOK, d)
}

// Обновить статус устройства
func (h *DeviceHandler) updateStatus(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Status json.RawMessage `json:"status"`
	}
	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.UpdateStatus(context.Background(), id, string(body.Status)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "status updated"})
}

// Отправка команды устройству
func (h *DeviceHandler) sendCommand(c *gin.Context) {
	idStr := c.Param("id")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid device ID format",
			"details": "Device ID must be an integer",
		})
		return
	}
	var cmd mq.Command
	if err := c.BindJSON(&cmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cmd.DeviceID = id

	if err := h.producer.SendCommand(cmd); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message":   "command sent",
		"device_id": id,
		"type":      cmd.Type,
	})
}
