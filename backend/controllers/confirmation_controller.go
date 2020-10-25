package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tew147258/app/ent"
	"github.com/tew147258/app/ent/borrow"
	"github.com/tew147258/app/ent/confirmation"
	"github.com/tew147258/app/ent/stadium"
	"github.com/tew147258/app/ent/user"
)

// ConfirmationController defines the struct for the confirmation controller
type ConfirmationController struct {
	client *ent.Client
	router gin.IRouter
}
type Confirmation struct {
	User         int
	Stadium      int
	Borrow       int
	Adddate      string
	Bookingstart string
	Bookingend   string
	Hourstime    string
}

// CreateConfirmation handles POST requests for adding confirmation entities
// @Summary Create confirmation
// @Description Create confirmation
// @ID create-confirmation
// @Accept   json
// @Produce  json
// @Param confirmation body Confirmation true "Confirmation entity"
// @Success 200 {object} ent.Confirmation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations [post]
func (ctl *ConfirmationController) CreateConfirmation(c *gin.Context) {
	obj := Confirmation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "confirmation video binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
		return
	}

	s, err := ctl.client.Stadium.
		Query().
		Where(stadium.IDEQ(int(obj.Stadium))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "stadium not found",
		})
		return
	}

	b, err := ctl.client.Borrow.
		Query().
		Where(borrow.IDEQ(int(obj.Borrow))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "borrow not found",
		})
		return
	}

	t1 := time.Now()
	t2 := t1.Format("2006-01-02T15:04:05Z07:00")
	time1, err := time.Parse(time.RFC3339, t2)
	time2, err := time.Parse(time.RFC3339, obj.Bookingstart)
	time3, err := time.Parse(time.RFC3339, obj.Bookingend)
	time4 := time2.After(time3)
	time5 := time2.Before(time1)
	timehours := time3.Sub(time2)
	timeH := float32(timehours) / 3600000000000

	if time4 {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	} else {
		if time5 {
			c.JSON(400, gin.H{
				"error": "saving failed",
			})
			return
		} else {
			if timeH == 1 || timeH == 2 || timeH == 3 || timeH == 4 || timeH == 5 || timeH == 6 || timeH == 7 || timeH == 8 {
				co, err := ctl.client.Confirmation.
					Create().
					SetConfirmationUser(u).
					SetConfirmationStadium(s).
					SetConfirmationBorrow(b).
					SetAdddate(time1).
					SetBookingstart(time2).
					SetBookingend(time3).
					SetHourstime(int(timeH)).
					Save(context.Background())

				if err != nil {
					c.JSON(400, gin.H{
						"error": "saving failed",
					})
					return
				}

				c.JSON(200, co)
			} else {
				c.JSON(400, gin.H{
					"error": "saving failed",
				})
				return
			}
		}
	}

}

// GetConfirmation handles GET requests to retrieve a confirmation entity
// @Summary Get a confirmation entity by ID
// @Description get confirmation by ID
// @ID get-confirmation
// @Produce  json
// @Param id path int true "Confirmation ID"
// @Success 200 {object} ent.Confirmation
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations/{id} [get]
func (ctl *ConfirmationController) GetConfirmation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	co, err := ctl.client.Confirmation.
		Query().
		Where(confirmation.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, co)
}

// ListConfirmation handles request to get a list of confirmation entities
// @Summary List confirmation entities
// @Description list confirmation entities
// @ID list-confirmation
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Confirmation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations [get]
func (ctl *ConfirmationController) ListConfirmation(c *gin.Context) {
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

	confirmations, err := ctl.client.Confirmation.
		Query().
		WithConfirmationUser().
		WithConfirmationStadium().
		WithConfirmationBorrow().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, confirmations)
}

// DeleteConfirmation handles DELETE requests to delete a confirmation entity
// @Summary Delete a confirmation entity by ID
// @Description get confirmation by ID
// @ID delete-confirmation
// @Produce  json
// @Param id path int true "Confirmation ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations/{id} [delete]
func (ctl *ConfirmationController) DeleteConfirmation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Confirmation.
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

// NewConfirmationController creates and registers handles for the confirmation controller
func NewConfirmationController(router gin.IRouter, client *ent.Client) *ConfirmationController {
	coc := &ConfirmationController{
		client: client,
		router: router,
	}
	coc.register()
	return coc
}

// InitConfirmationController registers routes to the main engine
func (ctl *ConfirmationController) register() {
	confirmations := ctl.router.Group("/confirmations")

	confirmations.GET("", ctl.ListConfirmation)

	// CRUD
	confirmations.POST("", ctl.CreateConfirmation)
	confirmations.GET(":id", ctl.GetConfirmation)
	confirmations.DELETE(":id", ctl.DeleteConfirmation)
}
