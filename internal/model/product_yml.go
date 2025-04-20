package model

// ProductInfo holds basic product information.
type ProductInfo struct {
	NameVN      string `yaml:"name_vn"`
	NameEN      string `yaml:"name_en"`
	ServingSize string `yaml:"serving_size"`
	Code        string `yaml:"code"`
	FoodWaste   string `yaml:"food_waste"`
}

// NutritionalComposition holds nutritional values.
type NutritionalComposition struct {
	Water        Nutrient `yaml:"water"`
	EnergyKcal   Nutrient `yaml:"energy_kcal"`
	EnergyKj     Nutrient `yaml:"energy_kj"`
	Protein      Nutrient `yaml:"protein"`
	Fat          Nutrient `yaml:"fat"`
	Carbohydrate Nutrient `yaml:"carbohydrate"`
	Fiber        Nutrient `yaml:"fiber"`
	Ash          Nutrient `yaml:"ash"`
}

// Nutrient holds nutrient value and unit.
type Nutrient struct {
	Unit  string  `yaml:"unit"`
	Value float64 `yaml:"value"`
}

// SugarContent holds various sugar types.
type SugarContent struct {
	TotalSugar Nutrient `yaml:"total_sugar"`
	Galactose  Nutrient `yaml:"galactose"`
	Maltose    Nutrient `yaml:"maltose"`
	Lactose    Nutrient `yaml:"lactose"`
	Fructose   Nutrient `yaml:"fructose"`
	Glucose    Nutrient `yaml:"glucose"`
	Sucrose    Nutrient `yaml:"sucrose"`
}

// Minerals holds mineral content.
type Minerals struct {
	Calcium    Nutrient `yaml:"calcium"`
	Iron       Nutrient `yaml:"iron"`
	Magnesium  Nutrient `yaml:"magnesium"`
	Manganese  Nutrient `yaml:"manganese"`
	Phosphorus Nutrient `yaml:"phosphorus"`
	Potassium  Nutrient `yaml:"potassium"`
	Sodium     Nutrient `yaml:"sodium"`
	Zinc       Nutrient `yaml:"zinc"`
	Copper     Nutrient `yaml:"copper"`
	Selenium   Nutrient `yaml:"selenium"`
}

// FattyAcids holds fatty acid content.
type FattyAcids struct {
	TotalSaturatedFattyAcid Nutrient `yaml:"total_saturated_fatty_acid"`
	PalmiticC16_0           Nutrient `yaml:"palmitic_c16_0"`
	MargaricC17_0           Nutrient `yaml:"margaric_c17_0"`
	StearicC18_0            Nutrient `yaml:"stearic_c18_0"`
	ArachidicC20_0          Nutrient `yaml:"arachidic_c20_0"`
	BehenicC22_0            Nutrient `yaml:"behenic_c22_0"`
	LignocericC24_0         Nutrient `yaml:"lignoceric_c24_0"`

	TotalMonounsaturatedFattyAcid Nutrient `yaml:"total_monounsaturated_fatty_acid"`

	MyristoleicC14_1 Nutrient `yaml:"myristoleic_c14_1"`
	PalmitoleicC16_1 Nutrient `yaml:"palmitoleic_c16_1"`
	OleicC18_1       Nutrient `yaml:"oleic_c18_1"`

	TotalPolyunsaturatedFattyAcid Nutrient `yaml:"total_polyunsaturated_fatty_acid"`

	LinoleicC18_2N6  Nutrient `yaml:"linoleic_c18_2_n6"`
	LinolenicC18_3N3 Nutrient `yaml:"linolenic_c18_3_n3"`

	ArachidonicC20_4 Nutrient `yaml:"arachidonic_c20_4"`

	EicosapentaenoicC20_5N3 Nutrient `yaml:"eicosapentaenoic_c20_5_n3"`

	DocosahexaenoicC22_6N3 Nutrient `yaml:"docosahexaenoic_c22_6_n3"`

	TotalTransFattyAcid Nutrient `yaml:"total_trans_fatty_acid"`
}

// CholesterolPhytosterol holds cholesterol and phytosterol content.
type CholesterolPhytosterol struct {
	Cholesterol Nutrient `yaml:"cholesterol"`
	Phytosterol Nutrient `yaml:"phytosterol"`
}

// Vitamins holds vitamin content.
type Vitamins struct {
	VitaminC  Nutrient `yaml:"vitamin_c"`
	VitaminB1 Nutrient `yaml:"vitamin_b1"`
	VitaminB2 Nutrient `yaml:"vitamin_b2"`
	VitaminPP Nutrient `yaml:"vitamin_pp"` // Niacin
	VitaminB5 Nutrient `yaml:"vitamin_b5"` // Pantothenic acid
	VitaminB6 Nutrient `yaml:"vitamin_b6"`
	Folate    Nutrient `yaml:"folate"`
	VitaminA  Nutrient `yaml:"vitamin_a"`
	VitaminE  Nutrient `yaml:"vitamin_e"`
	VitaminK  Nutrient `yaml:"vitamin_k"`
}

// AminoAcids holds amino acid content.
type AminoAcids struct {
	AsparticAcid  Nutrient `yaml:"aspartic_acid"`
	Threonine     Nutrient `yaml:"threonine"`
	Serine        Nutrient `yaml:"serine"`
	GlutamicAcid  Nutrient `yaml:"glutamic_acid"`
	Proline       Nutrient `yaml:"proline"`
	Glycine       Nutrient `yaml:"glycine"`
	Alanine       Nutrient `yaml:"alanine"`
	Cystine       Nutrient `yaml:"cystine"`
	Valine        Nutrient `yaml:"valine"`
	Methionine    Nutrient `yaml:"methionine"`
	Isoleucine    Nutrient `yaml:"isoleucine"`
	Leucine       Nutrient `yaml:"leucine"`
	Tyrosine      Nutrient `yaml:"tyrosine"`
	Phenylalanine Nutrient `yaml:"phenylalanine"`
	Histidine     Nutrient `yaml:"#histidine"`
	Lysine        Nutrient `yamL:"#lysine"`
	Arginine      Nutrient `yamL:"#arginine"`
	Tryptophan    Nutrient `yamL:"#tryptophan"`
}

// CompleteData structure to hold everything together.
type CompleteData struct {
	ProductInfo            ProductInfo            // Basic product information
	NutritionalComposition NutritionalComposition // Holds nutritional values
	SugarContent           SugarContent           // Sugar content details
	Minerals               Minerals               // Mineral content
	FattyAcids             FattyAcids             // Fatty acid details
	CholesterolPhytosterol CholesterolPhytosterol // Cholesterol and phytosterol details
	Vitamins               Vitamins               // Vitamin content
	AminoAcids             AminoAcids             // Amino acid details
}
