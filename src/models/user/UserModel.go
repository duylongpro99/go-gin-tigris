package models

import (
	"net/http"
	"server/config"

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

func Read[T interface{}](c *gin.Context, id int32) (*tigris.Iterator[T], error) {
	coll := tigris.GetCollection[T](config.DB)
	u, err := coll.Read(c, filter.Eq("Id", id))
	return u, err
}

func Search[T interface{}](c *gin.Context, u search.Request) (*tigris.SearchIterator[T], error) {
	coll := tigris.GetCollection[T](config.DB)
	it, err := coll.Search(c, &u)
	if err != nil {
		return nil, err
	}
	return it, nil
}

func Create[T interface{}](c *gin.Context, u T) (T, error) {
	coll := tigris.GetCollection[T](config.DB)

	if _, err := coll.Insert(c, &u); err != nil {
		return u, err
	}

	return u, nil
}

func Update[T interface{}](c *gin.Context, id int32, u fields.Update) (int32, error) {
	coll := tigris.GetCollection[T](config.DB)
	if _, err := coll.Update(c, filter.Eq("Id", id), &u); err != nil {
		return id, err
	}

	return id, nil

}

func Delete[T interface{}](c *gin.Context, id int32) (bool, error) {
	coll := tigris.GetCollection[T](config.DB)

	if _, err := coll.Delete(c, filter.Eq("Id", id)); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return false, err
	}

	return true, nil
}
