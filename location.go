package bullhorn

type Location struct {
	Id      int `json:"id"`
	Address struct {
		Address1    string      `json:"address1"`
		Address2    interface{} `json:"address2"`
		City        string      `json:"city"`
		State       string      `json:"state"`
		StateID     int         `json:"stateID"`
		StateName   string      `json:"stateName"`
		Zip         string      `json:"zip"`
		CountryID   int         `json:"countryID"`
		CountryName string      `json:"countryName"`
		CountryCode string      `json:"countryCode"`
	} `json:"address"`
	Candidate      interface{} `json:"candidate"`
	ClientContacts struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"clientContacts"`
	ClientCorporation struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"clientCorporation"`
	CustomDate1      interface{} `json:"customDate1"`
	CustomDate2      interface{} `json:"customDate2"`
	CustomDate3      interface{} `json:"customDate3"`
	CustomFloat1     interface{} `json:"customFloat1"`
	CustomFloat2     interface{} `json:"customFloat2"`
	CustomFloat3     interface{} `json:"customFloat3"`
	CustomInt1       interface{} `json:"customInt1"`
	CustomInt2       interface{} `json:"customInt2"`
	CustomInt3       interface{} `json:"customInt3"`
	CustomText1      interface{} `json:"customText1"`
	CustomText10     interface{} `json:"customText10"`
	CustomText11     interface{} `json:"customText11"`
	CustomText12     interface{} `json:"customText12"`
	CustomText13     interface{} `json:"customText13"`
	CustomText14     interface{} `json:"customText14"`
	CustomText15     interface{} `json:"customText15"`
	CustomText16     interface{} `json:"customText16"`
	CustomText17     interface{} `json:"customText17"`
	CustomText18     interface{} `json:"customText18"`
	CustomText19     interface{} `json:"customText19"`
	CustomText2      interface{} `json:"customText2"`
	CustomText20     interface{} `json:"customText20"`
	CustomText3      interface{} `json:"customText3"`
	CustomText4      interface{} `json:"customText4"`
	CustomText5      interface{} `json:"customText5"`
	CustomText6      interface{} `json:"customText6"`
	CustomText7      interface{} `json:"customText7"`
	CustomText8      interface{} `json:"customText8"`
	CustomText9      interface{} `json:"customText9"`
	CustomTextBlock1 interface{} `json:"customTextBlock1"`
	CustomTextBlock2 interface{} `json:"customTextBlock2"`
	CustomTextBlock3 interface{} `json:"customTextBlock3"`
	DateAdded        int64       `json:"dateAdded"`
	DateLastModified int64       `json:"dateLastModified"`
	Description      interface{} `json:"description"`
	EffectiveDate    string      `json:"effectiveDate"`
	EffectiveEndDate string      `json:"effectiveEndDate"`
	ExternalID       string      `json:"externalID"`
	IsBillTo         bool        `json:"isBillTo"`
	IsDeleted        bool        `json:"isDeleted"`
	IsSoldTo         bool        `json:"isSoldTo"`
	IsWorkSite       bool        `json:"isWorkSite"`
	Owner            struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"owner"`
	Status    string `json:"status"`
	Title     string `json:"title"`
	VersionID int    `json:"versionID"`
	Versions  struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"versions"`
}
