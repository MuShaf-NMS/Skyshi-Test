package handler

import (
	"fmt"
	"strconv"

	"github.com/MuShaf-NMS/Skyshi-Test/dto"
	"github.com/MuShaf-NMS/Skyshi-Test/helper"
	"github.com/MuShaf-NMS/Skyshi-Test/service"
	"github.com/gin-gonic/gin"
)

type TodoHandler interface {
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type todoHandler struct {
	service service.TodoService
}

func (ah *todoHandler) GetAll(ctx *gin.Context) {
	activityGroupIDString := ctx.Query("activity_group_id")
	activityGroupID := 0
	if activityGroupIDString != "" {
		a, err := strconv.Atoi(activityGroupIDString)
		if err != nil {
			res := helper.ErrorResponseBuilder("Not Found", "Invalid Query Params")
			ctx.JSON(404, res)
			return
		}
		activityGroupID = a
	}
	todos, err := ah.service.GetAll(int(activityGroupID))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", todos)
	ctx.JSON(200, res)
}

func (ah *todoHandler) Create(ctx *gin.Context) {
	var todo dto.CreateTodo
	ctx.BindJSON(&todo)
	if err := helper.Validate(todo); err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder("Bad Request", errs[0])
		ctx.JSON(400, res)
		return
	}
	t, err := ah.service.Create(todo)
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", t)
	ctx.JSON(201, res)
}

func (ah *todoHandler) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
		ctx.JSON(404, res)
		return
	}
	todo, err := ah.service.GetOne(int(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", todo)
	ctx.JSON(200, res)
}

func (ah *todoHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
		ctx.JSON(404, res)
		return
	}
	var todo dto.UpdateTodo
	ctx.BindJSON(&todo)
	if err := helper.Validate(todo); err != nil {
		errs := helper.ValidationError(err)
		fmt.Println(errs)
		res := helper.ErrorResponseBuilder("Bad Request", errs[0])
		ctx.JSON(400, res)
		return
	}
	t, err := ah.service.Update(todo, int(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", t)
	ctx.JSON(200, res)
}

func (ah *todoHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
		ctx.JSON(404, res)
		return
	}
	err = ah.service.Delete(int(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", gin.H{})
	ctx.JSON(200, res)
}

func NewTodoHandler(service service.TodoService) TodoHandler {
	return &todoHandler{
		service: service,
	}
}
