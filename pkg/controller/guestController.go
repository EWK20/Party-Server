package controller

import (
	"log"
	"net/http"

	"github.com/getground/tech-tasks/backend/pkg/dto"
	"github.com/getground/tech-tasks/backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type GuestController interface {
	GetGuests(ctx *gin.Context)
	CreateGuest(ctx *gin.Context)
	Checkin(ctx *gin.Context)
	Checkout(ctx *gin.Context)
	GetArrivedGuests(ctx *gin.Context)
}

type guestController struct {
	guestService service.GuestService
}

func NewGuestController(guestS service.GuestService) GuestController {
	return &guestController{
		guestService: guestS,
	}
}

func (c *guestController) GetGuests(ctx *gin.Context) {
	res, err := c.guestService.FindAll()
	if err != nil {
		log.Println("Could not retrieve guests")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Println("Get Guests Controller - Successfully retrieved guests")
	ctx.IndentedJSON(http.StatusFound, res)
}

func (c *guestController) CreateGuest(ctx *gin.Context) {
	var req dto.GuestReqDto
	var emptyRes dto.GuestResDto

	err := ctx.BindJSON(&req)
	if err != nil {
		log.Println("Could not retrieve guest data")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	name := ctx.Param("name")

	req.Name = name

	res, err := c.guestService.Save(req)
	if err != nil {
		log.Println("Create Guest Controller - Could not create new guest")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if res == emptyRes {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "too many guests"})
	} else {
		log.Println("Create Guest Controller - Successfully added to guest list")
		ctx.IndentedJSON(http.StatusCreated, res)
	}
}

func (c *guestController) Checkin(ctx *gin.Context) {
	var req dto.GuestReqDto
	var emptyRes dto.GuestResDto

	err := ctx.BindJSON(&req)
	if err != nil {
		log.Println("Checkin Controller - Could not retrieve guest data")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	name := ctx.Param("name")

	req.Name = name

	res, err := c.guestService.Checkin(req)
	if err != nil {
		log.Println("Checkin Controller - Could not update guest")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	if res == emptyRes {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "too many guests"})
	} else {
		log.Println("Checkin Controller - Successfully checked in guest")
		ctx.IndentedJSON(http.StatusCreated, res)
	}
}

func (c *guestController) Checkout(ctx *gin.Context) {
	name := ctx.Param("name")

	err := c.guestService.Checkout(name)
	if err != nil {
		log.Println("Checkout Controller - Could not delete guest")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Println("Checkout Controller - Successfully deleted guest")
	ctx.IndentedJSON(http.StatusNoContent, nil)
}

func (c *guestController) GetArrivedGuests(ctx *gin.Context) {
	res, err := c.guestService.GetArrivedGuests()
	if err != nil {
		log.Println("Could not retrieve guests")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Println("Get Guests Controller - Successfully retrieved guests")
	ctx.IndentedJSON(http.StatusFound, res)
}
