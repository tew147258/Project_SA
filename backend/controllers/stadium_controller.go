package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tew147258/app/ent"
	"github.com/tew147258/app/ent/stadium"
)

// StadiumController defines the struct for the stadium controller
type StadiumController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateStadium handles POST requests for adding stadium entities
// @Summary Create stadium
// @Description Create stadium
// @ID create-stadium
// @Accept   json
// @Produce  json
// @Param stadium body ent.Stadium true "Stadium entity"
// @Success 200 {object} ent.Stadium
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stadiums [post]
func (ctl *StadiumController) CreateStadium(c *gin.Context) {
	obj := ent.Stadium{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "stadium binding failed",
		})
		return
	}

	s, err := ctl.client.Stadium.
		Create().
		SetNamestadium(obj.Namestadium).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, s)
}

// GetStadium handles GET requests to retrieve a stadium entity
// @Summary Get a stadium entity by ID
// @Description get stadium by ID
// @ID get-stadium
// @Produce  json
// @Param id path int true "Stadium ID"
// @Success 200 {object} ent.Stadium
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stadiums/{id} [get]
func (ctl *StadiumController) GetStadium(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	s, err := ctl.client.Stadium.
		Query().
		Where(stadium.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, s)
}

// ListStadium handles request to get a list of stadium entities
// @Summary List stadium entities
// @Description list stadium entities
// @ID list-stadium
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Stadium
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stadiums [get]
func (ctl *StadiumController) ListStadium(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	stadiums, err := ctl.client.Stadium.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, stadiums)
}

// DeleteStadium handles DELETE requests to delete a stadium entity
// @Summary Delete a stadium entity by ID
// @Description get stadium by ID
// @ID delete-stadium
// @Produce  json
// @Param id path int true "Stadium ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stadiums/{id} [delete]
func (ctl *StadiumController) DeleteStadium(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Stadium.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateStadium handles PUT requests to update a stadium entity
// @Summary Update a stadium entity by ID
// @Description update stadium by ID
// @ID update-stadium
// @Accept   json
// @Produce  json
// @Param id path int true "Stadium ID"
// @Param stadium body ent.Stadium true "Stadium entity"
// @Success 200 {object} ent.Stadium
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /stadiums/{id} [put]
func (ctl *StadiumController) UpdateStadium(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Stadium{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "stadium binding failed",
		})
		return
	}

	s, err := ctl.client.Stadium.
		UpdateOneID(int(id)).
		SetNamestadium(obj.Namestadium).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, s)
}

// NewStadiumController creates and registers handles for the stadium controller
func NewStadiumController(router gin.IRouter, client *ent.Client) *StadiumController {
	sc := &StadiumController{
		client: client,
		router: router,
	}
	sc.register()
	return sc
}

// InitStadiumController registers routes to the main engine
func (ctl *StadiumController) register() {
	stadiums := ctl.router.Group("/stadiums")

	stadiums.GET("", ctl.ListStadium)

	// CRUD
	stadiums.POST("", ctl.CreateStadium)
	stadiums.GET(":id", ctl.GetStadium)
	stadiums.PUT(":id", ctl.UpdateStadium)
	stadiums.DELETE(":id", ctl.DeleteStadium)
}
