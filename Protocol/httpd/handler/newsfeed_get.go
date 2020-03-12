package handler

import (
			"net/http"
			"protocol/platform/newsfeed"
			"github.com/gin-gonic/gin"
)

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc {
		return		func(c *gin.Context) {
			results := feed.GetAll()
			c.JSON(http.StatusOK, results)


		// c.JSON(http.StatusOK, map[string]string{
		// 	"Hello":"Found me",
		// })
}
}