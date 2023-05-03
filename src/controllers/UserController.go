package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tigrisdata/tigris-client-go/fields"
	"github.com/tigrisdata/tigris-client-go/search"

	models "server/src/models"
	userModel "server/src/models/user"
	services "server/src/services"
)

func InitUserController(r *gin.Engine) {
	v1 := r.Group(fmt.Sprintf("/api/v1/%s", models.ServerUsers))
	SetReadRoute(v1)
	SetSearchRoute(v1)
	SetCreateRoute(v1)
	SetUpdateRoute(v1)
	SetDeleteRoute(v1)
}

// @BasePath /api/v1
// @Summary Read user by id
// @Schemes
// @Description Read user by id
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} SetReadRoute
// @Router /users/:id [get]
func SetReadRoute(r *gin.RouterGroup) {
	r.GET("/:id",
		func(c *gin.Context) {
			id, err := strconv.Atoi(c.Param("id"))
			fmt.Println(id)
			u, err := services.Read[userModel.User](c, int32(id))
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, u)
		})
}

// @BasePath /api/v1
// @Summary Search list users
// @Schemes
// @Description Search list users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} SetSearchRoute
// @Router /users/search [post]
func SetSearchRoute(r *gin.RouterGroup) {
	r.POST("/search", func(c *gin.Context) {

		var u search.Request
		if err := c.Bind(&u); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		it, _ := services.Search[userModel.User](c, u)

		r := &search.Result[userModel.User]{}
		for it.Next(r) {
			c.JSON(http.StatusOK, r)
		}

		if err := it.Err(); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
}

// @BasePath /api/v1
// @Summary Create user
// @Schemes
// @Description Create user
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} SetCreateRoute
// @Router /users/create [post]
func SetCreateRoute(r *gin.RouterGroup) {
	r.POST("/create", func(c *gin.Context) {

		var u userModel.User
		if err := c.Bind(&u); err != nil {
			return
		}

		if _, err := services.Create(c, u); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, u)

	})
}

// @BasePath /api/v1
// @Summary Update user
// @Schemes
// @Description Update user
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} SetUpdateRoute
// @Router /users/update [put]
func SetUpdateRoute(r *gin.RouterGroup) {
	r.PUT("/update/:id", func(c *gin.Context) {
		var u fields.Update
		if err := c.Bind(&u); err != nil {
			return
		}

		id, _ := strconv.Atoi(c.Param("id"))

		if _, err := services.Update[userModel.User](c, int32(id), u); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, u)

	})
}

// @BasePath /api/v1
// @Summary Delete user
// @Schemes
// @Description Delete user
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} SetDeleteRoute
// @Router /users/delete/:id [delete]
func SetDeleteRoute(r *gin.RouterGroup) {
	r.POST("/delete/:id", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		if _, err := services.Delete[userModel.User](c, int32(id)); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Status": "DELETED"})

	})
}
