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
