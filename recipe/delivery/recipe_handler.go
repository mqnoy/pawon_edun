package delivery

import (
	"example/web-service-gin/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	g        *gin.Engine
	RUsecase domain.RecipeUsecase
}

func NewRecipeHandler(g *gin.Engine, ru domain.RecipeUsecase) {
	handler := &RecipeHandler{
		g:        g,
		RUsecase: ru,
	}

	api := g.Group("/api/v1/recipes")
	{
		api.GET("/:id", handler.FetchOneRecipe)
		api.PATCH("/:id", handler.UpdateRecipe)
		api.POST("/:id", handler.CreateRecipe)
		api.DELETE("/:id", handler.DeleteRecipe)
		api.GET("/", handler.FetchAllRecipe)
	}

}

func (rh *RecipeHandler) FetchOneRecipe(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "FetchOneRecipe",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *RecipeHandler) UpdateRecipe(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "UpdateRecipe",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *RecipeHandler) CreateRecipe(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "CreateRecipe",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *RecipeHandler) DeleteRecipe(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "DeleteRecipe",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}

func (rh *RecipeHandler) FetchAllRecipe(g *gin.Context) {
	// TODO: code here
	g.JSON(http.StatusOK, gin.H{
		"message": "FetchOneRecipe",
		"code":    http.StatusOK,
		"status":  true,
		"data":    nil,
	})
}
