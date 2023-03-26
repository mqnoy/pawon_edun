package delivery

import (
	"example/web-service-gin/domain"
	utils "example/web-service-gin/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type CookerHandler struct {
	g        *gin.Engine
	CUsecase domain.CookerUsecase
}

func NewCookerHandler(g *gin.Engine, cu domain.CookerUsecase) {
	handler := CookerHandler{
		g:        g,
		CUsecase: cu,
	}

	api := g.Group("/api/v1/cookers")
	{
		api.GET(":cookerId/recipes", handler.FetchAllRecipe)
	}

}

func (ch *CookerHandler) FetchAllRecipe(c *gin.Context) {
	cookerId, _ := strconv.Atoi(c.Params.ByName("cookerId"))

	pagination := utils.Pagination(c)
	logrus.Info("cookerId:", cookerId)
	recipes, err := ch.CUsecase.ListAllRecipes(uint(cookerId), &pagination)

	if err != nil {
		logrus.Error(err)
	}

	c.JSON(200, gin.H{
		"message": "TODO: FetchAllRecipe",
		"code":    http.StatusOK,
		"status":  true,
		"data":    recipes,
	})
}
