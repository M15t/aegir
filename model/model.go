package model

// NutritionalFact represents the nutritional facts of a food product
type NutritionalFact struct {
	ProductInfo            ProductInfo         `json:"product_info"`
	NutritionalComposition map[string]Nutrient `json:"nutritional_composition"`
	SugarContent           map[string]Nutrient `json:"sugar_content"`
	Minerals               map[string]Nutrient `json:"minerals"`
	FattyAcids             map[string]Nutrient `json:"fatty_acids"`
	CholesterolPhytosterol map[string]Nutrient `json:"cholesterol_phytosterol"`
	Vitamins               map[string]Nutrient `json:"vitamins"`
	AminoAcids             map[string]Nutrient `json:"amino_acids"`
}

// Nutrient represents a nutrient in a food product
type Nutrient struct {
	Unit  string   `json:"unit"`
	Value *float64 `json:"value"`
}

// ProductInfo represents the basic information of a food product
type ProductInfo struct {
	NameVN      string `json:"name_vn"`
	NameEN      string `json:"name_en"`
	ServingSize string `json:"serving_size"`
	Code        string `json:"code"`
	FoodWaste   string `json:"food_waste"`
}
