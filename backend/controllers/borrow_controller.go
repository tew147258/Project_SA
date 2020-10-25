package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tew147258/app/ent"
	"github.com/tew147258/app/ent/borrow"
)

// BorrowController defines the struct for the borrow controller
type BorrowController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBorrow handles POST requests for adding borrow entities
// @Summary Create borrow
// @Description Create borrow
// @ID create-borrow
// @Accept   json
// @Produce  json
// @Param borrow body ent.Borrow true "Borrow entity"
// @Success 200 {object} ent.Borrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /borrows [post]
func (ctl *BorrowController) CreateBorrow(c *gin.Context) {
	obj := ent.Borrow{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "borrow binding failed",
		})
		return
	}

	b, err := ctl.client.Borrow.
		Create().
		SetType(obj.Type).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, b)
}

// GetBorrow handles GET requests to retrieve a borrow entity
// @Summary Get a borrow entity by ID
// @Description get borrow by ID
// @ID get-borrow
// @Produce  json
// @Param id path int true "Borrow ID"
// @Success 200 {object} ent.Borrow
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /borrows/{id} [get]
func (ctl *BorrowController) GetBorrow(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	b, err := ctl.client.Borrow.
		Query().
		Where(borrow.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, b)
}

// ListBorrow handles request to get a list of borrow entities
// @Summary List borrow entities
// @Description list borrow entities
// @ID list-borrow
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Borrow
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /borrows [get]
func (ctl *BorrowController) ListBorrow(c *gin.Context) {
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

	borrows, err := ctl.client.Borrow.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, borrows)
}

// NewBorrowController creates and registers handles for the borrow controller
func NewBorrowController(router gin.IRouter, client *ent.Client) *BorrowController {
	bc := &BorrowController{
		client: client,
		router: router,
	}
	bc.register()
	return bc
}

// InitBorrowController registers routes to the main engine
func (ctl *BorrowController) register() {
	borrows := ctl.router.Group("/borrows")

	borrows.GET("", ctl.ListBorrow)

	// CRUD
	borrows.POST("", ctl.CreateBorrow)
	borrows.GET(":id", ctl.GetBorrow)
}
