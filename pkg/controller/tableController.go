//The controller handles all incoming requests and outbound responses
package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/getground/tech-tasks/backend/pkg/dto"
	"github.com/getground/tech-tasks/backend/pkg/service"
	"github.com/gin-gonic/gin"
)

type TableController interface {
	GetTables(ctx *gin.Context)
	GetATable(ctx *gin.Context)
	CreateTable(ctx *gin.Context)
	GetSpace(ctx *gin.Context)
}

type tableController struct {
	tableService service.TableService
}

func NewTableController(tableS service.TableService) TableController {
	return &tableController{
		tableService: tableS,
	}
}

func HandlerPing(ctx *gin.Context) {
	fmt.Fprintf(ctx.Writer, "Hello World\n")
}

func (c *tableController) GetTables(ctx *gin.Context) {
	res, err := c.tableService.FindAll()
	if err != nil {
		log.Println("Create Table Controller - Could not create new table")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Println("Create Table Controller - Successfully retrieved all tables")
	ctx.IndentedJSON(http.StatusOK, res)
}

func (c *tableController) GetATable(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))

	res, err := c.tableService.FindById(id)
	if err != nil {
		log.Println("Get Table By Id Controller - Could not get table")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Println("Create Table Controller - Successfully retrieved table")
	ctx.IndentedJSON(http.StatusFound, res)
}

func (c *tableController) CreateTable(ctx *gin.Context) {
	var req dto.TableReqDto

	err := ctx.BindJSON(&req)
	if err != nil {
		log.Println("Could not retrieve table data")
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	res, err := c.tableService.Save(req)
	if err != nil {
		log.Println("Create Table Controller - Could not create new table")
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	log.Println("Create Table Controller - Successfully added table")
	ctx.IndentedJSON(http.StatusCreated, res)
}

func (c *tableController) GetSpace(ctx *gin.Context) {
	space, ok := c.tableService.CheckSpace()
	if !ok {
		log.Println("Available Space Controller - There are no empty seats")
		ctx.IndentedJSON(http.StatusNoContent, gin.H{"seats_empty": 0})
	}

	log.Println("Available Space Controller - There are empty seats")
	ctx.IndentedJSON(http.StatusOK, gin.H{"seats_empty": space})
}
