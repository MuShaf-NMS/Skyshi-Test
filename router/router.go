package router

import (
	"github.com/MuShaf-NMS/Skyshi-Test/config"
	"github.com/MuShaf-NMS/Skyshi-Test/handler"
	"github.com/MuShaf-NMS/Skyshi-Test/repository"
	"github.com/MuShaf-NMS/Skyshi-Test/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitializeRoute(server *gin.Engine, db *gorm.DB, config config.Config) {
	// Define Repositories
	activityRepository := repository.NewActivityRepository(db)
	todoRepository := repository.NewTodoRepository(db)

	// Define Services
	activityService := service.NewActivityService(activityRepository)
	todoService := service.NewTodoService(todoRepository)

	// Define Handlers
	activityHandler := handler.NewActivityHandler(activityService)
	todoHandler := handler.NewTodoHandler(todoService)

	// activity router
	activityRouter := server.Group("/activity-groups")
	{
		activityRouter.GET("", activityHandler.GetAll)
		activityRouter.POST("", activityHandler.Create)
		activityRouter.GET("/:id", activityHandler.GetOne)
		activityRouter.PATCH("/:id", activityHandler.Update)
		activityRouter.DELETE("/:id", activityHandler.Delete)
	}

	// todo router
	todoRouter := server.Group("/todo-items")
	{
		todoRouter.GET("", todoHandler.GetAll)
		todoRouter.POST("", todoHandler.Create)
		todoRouter.GET("/:id", todoHandler.GetOne)
		todoRouter.PATCH("/:id", todoHandler.Update)
		todoRouter.DELETE("/:id", todoHandler.Delete)
	}

}
