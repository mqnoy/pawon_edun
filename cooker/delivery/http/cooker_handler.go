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
		api.GET("/:id/recipes", handler.FetchAllRecipe)
		api.GET("/:id", handler.FetchOneCooker)
		api.PATCH("/:id", handler.UpdateCooker)
		api.POST("/:id", handler.CreateCooker)
		api.DELETE("/:id", handler.DeleteCooker)
		api.GET("/", handler.FetchAllCooker)
	}

}

func (ch *CookerHandler) FetchAllRecipe(c *gin.Context) {
	cookerId, _ := strconv.Atoi(c.Params.ByName("id"))

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

func (rh *CookerHandler) FetchOneCooker(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "FetchOneCooker",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *CookerHandler) UpdateCooker(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "UpdateCooker",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *CookerHandler) CreateCooker(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "CreateCooker",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *CookerHandler) DeleteCooker(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "DeleteCooker",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *CookerHandler) FetchAllCooker(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "FetchOneCooker",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}
