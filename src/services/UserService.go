package services

import (
	userModel "server/src/models/user"

	"github.com/gin-gonic/gin"
	"github.com/tigrisdata/tigris-client-go/fields"
	"github.com/tigrisdata/tigris-client-go/search"
	"github.com/tigrisdata/tigris-client-go/tigris"
)

func Read[T interface{}](c *gin.Context, id int32) (*tigris.Iterator[T], error) {
	return userModel.Read[T](c, id)
}

func Search[T interface{}](c *gin.Context, u search.Request) (*tigris.SearchIterator[T], error) {
	return userModel.Search[T](c, u)
}

func Create[T interface{}](c *gin.Context, u T) (T, error) {
	return userModel.Create(c, u)
}

func Update[T interface{}](c *gin.Context, id int32, u fields.Update) (int32, error) {
	return userModel.Update[T](c, id, u)
}

func Delete[T interface{}](c *gin.Context, id int32) (bool, error) {
	return userModel.Delete[T](c, id)
}
