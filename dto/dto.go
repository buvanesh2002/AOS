package dto

type Logindata struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Flockdata struct {
	ID         string       `json:"id,omitempty" bson:"_id,omitempty"`
	FlockName  string       `json:"flockName,omitempty" bson:"flockName,omitempty"`
	BreedName  string       `json:"breedName,omitempty" bson:"breedName,omitempty"`
	StartDate  string       `json:"startDate,omitempty" bson:"startDate,omitempty"`
	Age        string       `json:"startAge,omitempty" bson:"startAge,omitempty"`
	NoBirds    string       `json:"openingBirds,omitempty" bson:"openingBirds,omitempty"`
	ShedNumber string       `json:"shedNumber,omitempty" bson:"shedNumber,omitempty"`
	Active     string         `json:"active,omitempty" bson:"active,omitempty"`
	CreatedAt  string       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  string       `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Entry      []DailyEntry `json:"entry,omitempty" bson:"entry,omitempty"`
}

type DailyEntry struct {
	Date      string `json:"date,omitempty" bson:"date,omitempty"`
	Mortality string `json:"mortality,omitempty" bson:"mortality,omitempty"`
	ExtraEggs string `json:"extraeggs,omitempty" bson:"extraeggs,omitempty"`
	ExtraFeed string `json:"extrafeed,omitempty" bson:"extrafeed,omitempty"`
	BirdsSold string `json:"birdssold,omitempty" bson:"birdssold,omitempty"`
	CountErr  string `json:"counterr,omitempty" bson:"counterr,omitempty"`
	Remarks   string `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Trays     string `json:"trays,omitempty" bson:"trays,omitempty"`
}

