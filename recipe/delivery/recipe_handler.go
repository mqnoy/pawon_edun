package delivery

import (
	"example/web-service-gin/domain"

	"github.com/gin-gonic/gin"
)

type RecipeHandler struct {
	RUsecase domain.RecipeUsecase
}

func NewRecipeHandler(g *gin.Context, ru domain.RecipeUsecase) {
	// handler := &RecipeHandler{
	// 	RUsecase: ru,
	// }

	// g.Get()

}

// func () FetchRecipe() error {

// }
