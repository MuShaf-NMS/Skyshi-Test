package handler

import (
	"fmt"
	"strconv"

	"github.com/MuShaf-NMS/Skyshi-Test/dto"
	"github.com/MuShaf-NMS/Skyshi-Test/helper"
	"github.com/MuShaf-NMS/Skyshi-Test/service"
	"github.com/gin-gonic/gin"
)

type ActivityHandler interface {
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	GetOne(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type activityHandler struct {
	service service.ActivityService
}

func (ah *activityHandler) GetAll(ctx *gin.Context) {
	activities, err := ah.service.GetAll()
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", activities)
	ctx.JSON(200, res)
}

func (ah *activityHandler) Create(ctx *gin.Context) {
	var activity dto.CreateActivity
	ctx.BindJSON(&activity)
	if err := helper.Validate(activity); err != nil {
		errs := helper.ValidationError(err)
		res := helper.ErrorResponseBuilder("Bad Request", errs[0])
		ctx.JSON(400, res)
		return
	}
	a, err := ah.service.Create(activity)
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", a)
	ctx.JSON(201, res)
}

func (ah *activityHandler) GetOne(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
		ctx.JSON(400, res)
		return
	}
	activity, err := ah.service.GetOne(int(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", activity)
	ctx.JSON(200, res)
}

func (ah *activityHandler) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
		ctx.JSON(400, res)
		return
	}
	var activity dto.UpdateActivity
	ctx.BindJSON(&activity)
	if err := helper.Validate(activity); err != nil {
		errs := helper.ValidationError(err)
		fmt.Println(errs)
		res := helper.ErrorResponseBuilder("Bad Request", errs[0])
		ctx.JSON(400, res)
		return
	}
	a, err := ah.service.Update(activity, int(id))
	if err != nil {
		e, _ := err.(*helper.CustomError)
		res := helper.ErrorResponseBuilder(e.Status, e.Message)
		ctx.JSON(e.Code, res)
		return
	}
	res := helper.ResponseBuilder("Success", a)
	ctx.JSON(200, res)
}

func (ah *activityHandler) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		res := helper.ErrorResponseBuilder("Not Found", fmt.Sprintf("Todo with ID %d Not Found", id))
		ctx.JSON(400, res)
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

func NewActivityHandler(service service.ActivityService) ActivityHandler {
	return &activityHandler{
		service: service,
	}
}
