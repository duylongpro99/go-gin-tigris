package models

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tigrisdata/tigris-client-go/fields"
	"github.com/tigrisdata/tigris-client-go/filter"
	"github.com/tigrisdata/tigris-client-go/search"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

type User struct {
	ID          int32 `tigris:"primaryKey,autoGenerate"`
	FullName    string
	Email       string
	PhoneNumber string
}

type RegisterUser struct {
	FullName string
	Email    string
	Pwd      string
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
func SetReadRoute[T interface{}](v1 *gin.RouterGroup, db *tigris.Database, name string) {
	v1.GET(":id",
		func(c *gin.Context) {
			coll := tigris.GetCollection[T](db)
			u, err := coll.Read(c, filter.Eq("Id", c.Param("id")))
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
func SetSearchRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.POST(fmt.Sprintf("/%s/search", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)

		var u search.Request
		if err := c.Bind(&u); err != nil {
			_ = c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		it, err := coll.Search(c, &u)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		r := &search.Result[T]{}
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
// @Schemes models.RegisterUser
// @Description Create user
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {string} SetCreateRoute
// @Router /users/create [post]
func SetCreateRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.POST(fmt.Sprintf("/%s/create", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)
		var u T
		if err := c.Bind(&u); err != nil {
			return
		}

		if _, err := coll.Insert(c, &u); err != nil {
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
func SetUpdateRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.PUT(fmt.Sprintf("/%s/update/:id", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)
		var u fields.Update
		if err := c.Bind(&u); err != nil {
			return
		}

		if _, err := coll.Update(c, filter.Eq("Id", c.Param("id")), &u); err != nil {
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
func SetDeleteRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.POST(fmt.Sprintf("/%s/delete/:id", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)

		if _, err := coll.Delete(c, filter.Eq("Id", c.Param("id"))); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"Status": "DELETED"})

	})
}

func SetCRUDRoutes[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	//set up swagger
	v1 := r.Group(fmt.Sprintf("/api/v1/%s", name))
	SetReadRoute[T](v1, db, name)
	SetSearchRoute[T](r, db, name)
	SetCreateRoute[T](r, db, name)
	SetUpdateRoute[T](r, db, name)
	SetDeleteRoute[T](r, db, name)
}
