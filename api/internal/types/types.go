// Code generated by goctl. DO NOT EDIT.
package types

type AllIngredientsReply struct {
	IngredientList []IngrediendReply `json:"ingredientList"`
}

type AddIngredientReq struct {
	Name string `json:"name"`
}

type IngrediendReply struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type Recipe struct {
	Id             int64                    `json:"id"`
	Title          string                   `json:"title"`
	Description    string                   `json:"description"`
	IngredientList []IngredientWithQuantity `json:"ingredientList"`
	StageList      []Stage                  `json:"stageList"`
}

type IngredientWithQuantity struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type Stage struct {
	Order       int32  `json:"order"`
	Description string `json:"description"`
}

type RecipesReply struct {
	RecipeList []Recipe `json:"recipeList"`
}

type IngredientConstraintsReq struct {
	WithIngredientIdList    []int64 `json:"withIngredientIdList"`
	WithoutIngredientIdList []int64 `json:"withoutIngredientIdList"`
}

type AddRecipeReq struct {
	Title          string                       `json:"title"`
	Description    string                       `json:"description"`
	IngredientList []LinkIngredientWithQuantity `json:"ingredientList"`
	StageList      []Stage                      `json:"stageList"`
}

type LinkIngredientWithQuantity struct {
	IngredientId int64   `json:"ingredientId"`
	Quantity     float32 `json:"quantity"`
	Unit         string  `json:"unit"`
}

type AddRecipeReply struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type AuthReq struct {
	UserId int64 `json:"userId"`
}
