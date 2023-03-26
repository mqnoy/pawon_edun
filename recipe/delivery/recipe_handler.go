package delivery

import (
	"example/web-service-gin/domain"

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
}

func (rh *RecipeHandler) UpdateRecipe(g *gin.Context) {
}

func (rh *RecipeHandler) CreateRecipe(g *gin.Context) {
}

func (rh *RecipeHandler) DeleteRecipe(g *gin.Context) {
}

func (rh *RecipeHandler) FetchAllRecipe(g *gin.Context) {
}
