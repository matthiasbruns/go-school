package api

type mealsResponse struct {
	Meals []Meal `json:"meals"`
}
type Meal struct {
	IDMeal                      string `json:"idMeal"`
	StrMeal                     string `json:"strMeal"`
	StrDrinkAlternate           any    `json:"strDrinkAlternate"`
	StrCategory                 string `json:"strCategory"`
	StrArea                     string `json:"strArea"`
	StrInstructions             string `json:"strInstructions"`
	StrMealThumb                string `json:"strMealThumb"`
	StrTags                     any    `json:"strTags"`
	StrYoutube                  string `json:"strYoutube"`
	StrIngredient1              string `json:"strIngredient1"`
	StrIngredient2              string `json:"strIngredient2"`
	StrIngredient3              string `json:"strIngredient3"`
	StrIngredient4              string `json:"strIngredient4"`
	StrIngredient5              string `json:"strIngredient5"`
	StrIngredient6              string `json:"strIngredient6"`
	StrIngredient7              string `json:"strIngredient7"`
	StrIngredient8              string `json:"strIngredient8"`
	StrIngredient9              string `json:"strIngredient9"`
	StrIngredient10             string `json:"strIngredient10"`
	StrIngredient11             string `json:"strIngredient11"`
	StrIngredient12             string `json:"strIngredient12"`
	StrIngredient13             string `json:"strIngredient13"`
	StrIngredient14             string `json:"strIngredient14"`
	StrIngredient15             string `json:"strIngredient15"`
	StrIngredient16             string `json:"strIngredient16"`
	StrIngredient17             string `json:"strIngredient17"`
	StrIngredient18             string `json:"strIngredient18"`
	StrIngredient19             string `json:"strIngredient19"`
	StrIngredient20             string `json:"strIngredient20"`
	StrMeasure1                 string `json:"strMeasure1"`
	StrMeasure2                 string `json:"strMeasure2"`
	StrMeasure3                 string `json:"strMeasure3"`
	StrMeasure4                 string `json:"strMeasure4"`
	StrMeasure5                 string `json:"strMeasure5"`
	StrMeasure6                 string `json:"strMeasure6"`
	StrMeasure7                 string `json:"strMeasure7"`
	StrMeasure8                 string `json:"strMeasure8"`
	StrMeasure9                 string `json:"strMeasure9"`
	StrMeasure10                string `json:"strMeasure10"`
	StrMeasure11                string `json:"strMeasure11"`
	StrMeasure12                string `json:"strMeasure12"`
	StrMeasure13                string `json:"strMeasure13"`
	StrMeasure14                string `json:"strMeasure14"`
	StrMeasure15                string `json:"strMeasure15"`
	StrMeasure16                string `json:"strMeasure16"`
	StrMeasure17                string `json:"strMeasure17"`
	StrMeasure18                string `json:"strMeasure18"`
	StrMeasure19                string `json:"strMeasure19"`
	StrMeasure20                string `json:"strMeasure20"`
	StrSource                   string `json:"strSource"`
	StrImageSource              any    `json:"strImageSource"`
	StrCreativeCommonsConfirmed any    `json:"strCreativeCommonsConfirmed"`
	DateModified                any    `json:"dateModified"`
}
