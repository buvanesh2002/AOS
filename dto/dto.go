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
	StartAge   int          `json:"startAge,omitempty" bson:"startAge,omitempty"`
	Age        int          `json:"age,omitempty" bson:"age"`
	NoBirds    int          `json:"openingBirds,omitempty" bson:"openingBirds,omitempty"`
	ShedNumber string       `json:"shedNumber,omitempty" bson:"shedNumber,omitempty"`
	Active     string       `json:"active,omitempty" bson:"active,omitempty"`
	CreatedAt  string       `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt  string       `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
	Entry      []DailyEntry `json:"entry,omitempty" bson:"entry,omitempty"`
	ListEntry  []ListEntry `json:"listentry,omitempty" bson:"listentry,omitempty"`
}


type DailyEntry struct {
	ID string  `json:"id" bson:"id"`
	Date      string `json:"date,omitempty" bson:"date,omitempty"`
	Mortality int `json:"mortality,omitempty" bson:"mortality,omitempty"`
	Eggs int `json:"extraeggs,omitempty" bson:"extraeggs,omitempty"`
	Feed int `json:"feed,omitempty" bson:"feed,omitempty"`
	BirdsSold int `json:"birdssold,omitempty" bson:"birdssold,omitempty"`
	CountErr  int `json:"counterr,omitempty" bson:"counterr,omitempty"`
	Remarks   string `json:"remarks,omitempty" bson:"remarks,omitempty"`
	Trays     int `json:"trays,omitempty" bson:"trays,omitempty"`
}


type Reminder struct {
	ReminderId string `json:"reminderId,omitempty" bson:"reminderId,omitempty"`
	Name string `json:"remindername" bson:remindername"`
	BeforeDate string `json:"beforedate" bson:"beforedate"`
	AfterDate string `json:"afterdate" bson:"afterdate"`
	Date string `json:"reminderdate" bson:"reminderdate"`
	Remarks string `json:"remarks" bson:"remarks"`
	Status string `json:"status" bson:"status"`
}
type ListEntry struct {
	EntryDate string `json:"entrydate" bson:"entrydate"`
	Age int `json:"age" bson:"age"`
	OpeningBirds int `json:"openingbirds" bson:"openingbirds"`
	Mortality int `json:"mortality" bson:"mortality"`
	BirdsSold int `json:"birdssold" bson:"birdssold"`
	ClosingBirds int `json:"closingbirds" bson:"closingbirds"`
	CumMortality int `json:"cummortality" bson:"cummortality"`
	MortalityPer int `json:"mortality%" bson:"mortality%"`
	EggsPerDay int `json:"eggsperDay" bson:"eggsperDay"`
	EggProducion int `json:"eggproducion" bson:"eggproducion"`
	ProductionPer int `json:"production%" bson:"production%"`
	// HHP string  `json:"curHHP" bson:"curHHP"`
	// CumHHP string `json:"cumHHP" bson:"cumHHP"`
	Feed int 	`json:"feed" bson:"feed"`
	FeedPerBird int `json:"feedperBird" bson:"feedperBird"`
	FeedPerEgg int `json:"feedperEgg" bson:"feedperEgg"`
	CumFPE int `json:"cumFPE" bson:"cumFPE"`
	TotalFeed int `json:"totalFeed" bson:"totalFeed"`
}

