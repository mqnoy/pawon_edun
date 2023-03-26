package utils

import (
	"example/web-service-gin/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Pagination(c *gin.Context) model.Pagination {
	limit := 1
	page := 1
	sort := `created_at DESC`
	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}

	return model.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
