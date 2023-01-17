package main

import (
	"net/http"

	"github.com/getground/tech-tasks/backend/pkg/controller"
	"github.com/getground/tech-tasks/backend/pkg/repository"
	"github.com/getground/tech-tasks/backend/pkg/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var (
	tableRepository repository.TableRepository = repository.NewTableRepository()
	guestRepository repository.GuestRepository = repository.NewGuestRepository()

	tableService service.TableService = service.NewTableService(tableRepository)
	guestService service.GuestService = service.NewGuestService(guestRepository, tableRepository)

	tableController controller.TableController = controller.NewTableController(tableService)
	guestController controller.GuestController = controller.NewGuestController(guestService)
)

func main() {

	// Initializes an instance of the gin engine with logger and recovery functions
	router := gin.Default()

	// test ping
	router.GET("/ping", controller.HandlerPing)

	// Specifying routes
	// Before Party

	router.GET("/tables", tableController.GetTables)
	router.GET("/tables/:id", tableController.GetATable)
	router.POST("/tables", tableController.CreateTable)

	router.GET("/guest_list", guestController.GetGuests)
	router.POST("/guest_list/:name", guestController.CreateGuest)

	//During Party
	router.GET("/guests", guestController.GetArrivedGuests)
	router.GET("/seats_empty", tableController.GetSpace)
	router.PUT("/guests/:name", guestController.Checkin)
	router.DELETE("/guests/:name", guestController.Checkout)

	// Specifies what port the server will listen and answer on
	http.ListenAndServe(":4000", router)
}
