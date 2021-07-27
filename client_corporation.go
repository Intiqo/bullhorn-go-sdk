package bullhorn

type ClientCorporation struct {
	Id      int `json:"id"`
	Address struct {
		Address1    string `json:"address1"`
		Address2    string `json:"address2"`
		City        string `json:"city"`
		CountryCode string `json:"countryCode"`
		CountryID   int    `json:"countryID"`
		CountryName string `json:"countryName"`
		State       string `json:"state"`
		Timezone    string `json:"timezone"`
		Zip         string `json:"zip"`
	} `json:"address"`
	AnnualRevenue       int         `json:"annualRevenue"`
	BillingAddress      interface{} `json:"billingAddress"`
	BillingContact      interface{} `json:"billingContact"`
	BillingFrequency    interface{} `json:"billingFrequency"`
	BillingPhone        interface{} `json:"billingPhone"`
	Branch              interface{} `json:"branch"`
	BusinessSectorList  interface{} `json:"businessSectorList"`
	CertificationGroups struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"certificationGroups"`
	Certifications struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"certifications"`
	ChildClientCorporations struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"childClientCorporations"`
	ClientContactNotes struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"clientContactNotes"`
	ClientContacts struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"clientContacts"`
	CompanyDescription interface{} `json:"companyDescription"`
	CompanyURL         interface{} `json:"companyURL"`
	Competitors        interface{} `json:"competitors"`
	Culture            interface{} `json:"culture"`
	CustomDate1        interface{} `json:"customDate1"`
	CustomDate2        interface{} `json:"customDate2"`
	CustomDate3        interface{} `json:"customDate3"`
	CustomFloat1       interface{} `json:"customFloat1"`
	CustomFloat2       interface{} `json:"customFloat2"`
	CustomFloat3       interface{} `json:"customFloat3"`
	CustomInt1         interface{} `json:"customInt1"`
	CustomInt2         interface{} `json:"customInt2"`
	CustomInt3         interface{} `json:"customInt3"`
	CustomText1        interface{} `json:"customText1"`
	CustomText10       interface{} `json:"customText10"`
	CustomText11       interface{} `json:"customText11"`
	CustomText12       interface{} `json:"customText12"`
	CustomText13       interface{} `json:"customText13"`
	CustomText14       interface{} `json:"customText14"`
	CustomText15       interface{} `json:"customText15"`
	CustomText16       interface{} `json:"customText16"`
	CustomText17       interface{} `json:"customText17"`
	CustomText18       interface{} `json:"customText18"`
	CustomText19       interface{} `json:"customText19"`
	CustomText2        interface{} `json:"customText2"`
	CustomText20       interface{} `json:"customText20"`
	CustomText3        interface{} `json:"customText3"`
	CustomText4        interface{} `json:"customText4"`
	CustomText5        interface{} `json:"customText5"`
	CustomText6        interface{} `json:"customText6"`
	CustomText7        interface{} `json:"customText7"`
	CustomText8        interface{} `json:"customText8"`
	CustomText9        interface{} `json:"customText9"`
	CustomTextBlock1   interface{} `json:"customTextBlock1"`
	CustomTextBlock2   interface{} `json:"customTextBlock2"`
	CustomTextBlock3   interface{} `json:"customTextBlock3"`
	CustomTextBlock4   interface{} `json:"customTextBlock4"`
	CustomTextBlock5   interface{} `json:"customTextBlock5"`
	DateAdded          int64       `json:"dateAdded"`
	DateFounded        int64       `json:"dateFounded"`
	DateLastModified   int64       `json:"dateLastModified"`
	Department         struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"department"`
	ExternalID          interface{} `json:"externalID"`
	FacebookProfileName interface{} `json:"facebookProfileName"`
	Fax                 interface{} `json:"fax"`
	FeeArrangement      int         `json:"feeArrangement"`
	FileAttachments     struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"fileAttachments"`
	Funding       interface{} `json:"funding"`
	IndustryList  interface{} `json:"industryList"`
	InvoiceFormat interface{} `json:"invoiceFormat"`
	Leads         struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"leads"`
	LinkedinProfileName interface{} `json:"linkedinProfileName"`
	Locations           struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"locations"`
	Name         string      `json:"name"`
	Notes        interface{} `json:"notes"`
	NumEmployees int         `json:"numEmployees"`
	NumOffices   int         `json:"numOffices"`
	Owners       struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"owners"`
	Ownership               string      `json:"ownership"`
	ParentClientCorporation interface{} `json:"parentClientCorporation"`
	Phone                   interface{} `json:"phone"`
	Requirements            struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"requirements"`
	Revenue       interface{} `json:"revenue"`
	Status        string      `json:"status"`
	TaxRate       int         `json:"taxRate"`
	TickerSymbol  interface{} `json:"tickerSymbol"`
	TrackTitle    interface{} `json:"trackTitle"`
	TwitterHandle interface{} `json:"twitterHandle"`
	WorkWeekStart interface{} `json:"workWeekStart"`
}
