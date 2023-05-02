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
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
}

func SetReadRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.GET(fmt.Sprintf("/%s/read/:id", name), func(c *gin.Context) {
		coll := tigris.GetCollection[T](db)

		u, err := coll.Read(c, filter.Eq("Id", c.Param("id")))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, u)
	})
}

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

func SetUpdateRoute[T interface{}](r *gin.Engine, db *tigris.Database, name string) {
	r.POST(fmt.Sprintf("/%s/update/:id", name), func(c *gin.Context) {
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
	SetReadRoute[T](r, db, name)
	SetSearchRoute[T](r, db, name)
	SetCreateRoute[T](r, db, name)
	SetUpdateRoute[T](r, db, name)
	SetDeleteRoute[T](r, db, name)
}
