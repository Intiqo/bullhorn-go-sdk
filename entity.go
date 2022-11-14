package bullhorn

type Error struct {
	Key     string `json:"errorMessageKey"`
	Code    int    `json:"errorCode"`
	Message string `json:"errorMessage"`
}

type Options struct {
	Fields  string `json:"fields"`
	Start   int    `json:"start"`
	Count   int    `json:"count"`
	OrderBy string `json:"orderBy"`
}

type Query struct {
	Where string `json:"where"`
	Options
}

type Search struct {
	Query string `json:"query"`
	Options
}

type Ping struct {
	SessionExpires int64 `json:"sessionExpires"`
}

const CountryEntity = "Country"

type Country struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	States []int  `json:"states"`
}

const StateEntity = "State"

type State struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Code    string `json:"code"`
	Country int    `json:"country"`
}

const CandidateEntity = "Candidate"

type Candidate struct {
	Id               int `json:"id"`
	ActivePlacements struct {
		Total int   `json:"total"`
		Data  []int `json:"data"`
	} `json:"activePlacements"`
	Address struct {
		Address1    interface{} `json:"address1"`
		Address2    interface{} `json:"address2"`
		City        interface{} `json:"city"`
		CountryCode string      `json:"countryCode"`
		CountryID   int         `json:"countryID"`
		CountryName string      `json:"countryName"`
		State       interface{} `json:"state"`
		Timezone    interface{} `json:"timezone"`
		Zip         interface{} `json:"zip"`
	} `json:"address"`
	Branch          interface{} `json:"branch"`
	BusinessSectors struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"businessSectors"`
	CanEnterTime    bool        `json:"canEnterTime"`
	CandidateSource interface{} `json:"candidateSource"`
	Categories      struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"categories"`
	Category struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	CertificationList struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"certificationList"`
	Certifications             interface{} `json:"certifications"`
	ClientCorporationBlackList struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"clientCorporationBlackList"`
	ClientCorporationWhiteList struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"clientCorporationWhiteList"`
	Comments                    string      `json:"comments"`
	CompanyName                 interface{} `json:"companyName"`
	CompanyURL                  interface{} `json:"companyURL"`
	CustomDate1                 interface{} `json:"customDate1"`
	CustomDate10                interface{} `json:"customDate10"`
	CustomDate11                interface{} `json:"customDate11"`
	CustomDate12                interface{} `json:"customDate12"`
	CustomDate13                interface{} `json:"customDate13"`
	CustomDate2                 interface{} `json:"customDate2"`
	CustomDate3                 interface{} `json:"customDate3"`
	CustomDate4                 interface{} `json:"customDate4"`
	CustomDate5                 interface{} `json:"customDate5"`
	CustomDate6                 interface{} `json:"customDate6"`
	CustomDate7                 interface{} `json:"customDate7"`
	CustomDate8                 interface{} `json:"customDate8"`
	CustomDate9                 interface{} `json:"customDate9"`
	CustomEncryptedText1        interface{} `json:"customEncryptedText1"`
	CustomEncryptedText10       interface{} `json:"customEncryptedText10"`
	CustomEncryptedText2        interface{} `json:"customEncryptedText2"`
	CustomEncryptedText3        interface{} `json:"customEncryptedText3"`
	CustomEncryptedText4        interface{} `json:"customEncryptedText4"`
	CustomEncryptedText5        interface{} `json:"customEncryptedText5"`
	CustomEncryptedText6        interface{} `json:"customEncryptedText6"`
	CustomEncryptedText7        interface{} `json:"customEncryptedText7"`
	CustomEncryptedText8        interface{} `json:"customEncryptedText8"`
	CustomEncryptedText9        interface{} `json:"customEncryptedText9"`
	CustomFloat1                interface{} `json:"customFloat1"`
	CustomFloat10               interface{} `json:"customFloat10"`
	CustomFloat11               interface{} `json:"customFloat11"`
	CustomFloat12               interface{} `json:"customFloat12"`
	CustomFloat13               interface{} `json:"customFloat13"`
	CustomFloat14               interface{} `json:"customFloat14"`
	CustomFloat15               interface{} `json:"customFloat15"`
	CustomFloat16               interface{} `json:"customFloat16"`
	CustomFloat17               interface{} `json:"customFloat17"`
	CustomFloat18               interface{} `json:"customFloat18"`
	CustomFloat19               interface{} `json:"customFloat19"`
	CustomFloat2                interface{} `json:"customFloat2"`
	CustomFloat20               interface{} `json:"customFloat20"`
	CustomFloat21               interface{} `json:"customFloat21"`
	CustomFloat22               interface{} `json:"customFloat22"`
	CustomFloat23               interface{} `json:"customFloat23"`
	CustomFloat3                interface{} `json:"customFloat3"`
	CustomFloat4                interface{} `json:"customFloat4"`
	CustomFloat5                interface{} `json:"customFloat5"`
	CustomFloat6                interface{} `json:"customFloat6"`
	CustomFloat7                interface{} `json:"customFloat7"`
	CustomFloat8                interface{} `json:"customFloat8"`
	CustomFloat9                interface{} `json:"customFloat9"`
	CustomInt1                  interface{} `json:"customInt1"`
	CustomInt10                 interface{} `json:"customInt10"`
	CustomInt11                 interface{} `json:"customInt11"`
	CustomInt12                 interface{} `json:"customInt12"`
	CustomInt13                 interface{} `json:"customInt13"`
	CustomInt14                 interface{} `json:"customInt14"`
	CustomInt15                 interface{} `json:"customInt15"`
	CustomInt16                 interface{} `json:"customInt16"`
	CustomInt17                 interface{} `json:"customInt17"`
	CustomInt18                 interface{} `json:"customInt18"`
	CustomInt19                 interface{} `json:"customInt19"`
	CustomInt2                  interface{} `json:"customInt2"`
	CustomInt20                 interface{} `json:"customInt20"`
	CustomInt21                 interface{} `json:"customInt21"`
	CustomInt22                 interface{} `json:"customInt22"`
	CustomInt23                 interface{} `json:"customInt23"`
	CustomInt3                  interface{} `json:"customInt3"`
	CustomInt4                  interface{} `json:"customInt4"`
	CustomInt5                  interface{} `json:"customInt5"`
	CustomInt6                  interface{} `json:"customInt6"`
	CustomInt7                  interface{} `json:"customInt7"`
	CustomInt8                  interface{} `json:"customInt8"`
	CustomInt9                  interface{} `json:"customInt9"`
	CustomText1                 string      `json:"customText1"`
	CustomText10                interface{} `json:"customText10"`
	CustomText11                interface{} `json:"customText11"`
	CustomText12                interface{} `json:"customText12"`
	CustomText13                interface{} `json:"customText13"`
	CustomText14                interface{} `json:"customText14"`
	CustomText15                interface{} `json:"customText15"`
	CustomText16                interface{} `json:"customText16"`
	CustomText17                interface{} `json:"customText17"`
	CustomText18                interface{} `json:"customText18"`
	CustomText19                interface{} `json:"customText19"`
	CustomText2                 interface{} `json:"customText2"`
	CustomText20                interface{} `json:"customText20"`
	CustomText21                interface{} `json:"customText21"`
	CustomText22                interface{} `json:"customText22"`
	CustomText23                interface{} `json:"customText23"`
	CustomText24                interface{} `json:"customText24"`
	CustomText25                interface{} `json:"customText25"`
	CustomText26                interface{} `json:"customText26"`
	CustomText27                interface{} `json:"customText27"`
	CustomText28                interface{} `json:"customText28"`
	CustomText29                interface{} `json:"customText29"`
	CustomText3                 string      `json:"customText3"`
	CustomText30                interface{} `json:"customText30"`
	CustomText31                interface{} `json:"customText31"`
	CustomText32                interface{} `json:"customText32"`
	CustomText33                interface{} `json:"customText33"`
	CustomText34                interface{} `json:"customText34"`
	CustomText35                interface{} `json:"customText35"`
	CustomText36                interface{} `json:"customText36"`
	CustomText37                interface{} `json:"customText37"`
	CustomText38                interface{} `json:"customText38"`
	CustomText39                interface{} `json:"customText39"`
	CustomText4                 []string    `json:"customText4"`
	CustomText40                interface{} `json:"customText40"`
	CustomText5                 []string    `json:"customText5"`
	CustomText6                 interface{} `json:"customText6"`
	CustomText7                 interface{} `json:"customText7"`
	CustomText8                 interface{} `json:"customText8"`
	CustomText9                 interface{} `json:"customText9"`
	CustomTextBlock1            interface{} `json:"customTextBlock1"`
	CustomTextBlock10           interface{} `json:"customTextBlock10"`
	CustomTextBlock2            interface{} `json:"customTextBlock2"`
	CustomTextBlock3            interface{} `json:"customTextBlock3"`
	CustomTextBlock4            interface{} `json:"customTextBlock4"`
	CustomTextBlock5            interface{} `json:"customTextBlock5"`
	CustomTextBlock6            interface{} `json:"customTextBlock6"`
	CustomTextBlock7            interface{} `json:"customTextBlock7"`
	CustomTextBlock8            interface{} `json:"customTextBlock8"`
	CustomTextBlock9            interface{} `json:"customTextBlock9"`
	DateAdded                   int64       `json:"dateAdded"`
	DateAvailable               interface{} `json:"dateAvailable"`
	DateAvailableEnd            interface{} `json:"dateAvailableEnd"`
	DateI9Expiration            interface{} `json:"dateI9Expiration"`
	DateLastComment             interface{} `json:"dateLastComment"`
	DateLastModified            int64       `json:"dateLastModified"`
	DateLastPayrollProviderSync interface{} `json:"dateLastPayrollProviderSync"`
	DateNextCall                interface{} `json:"dateNextCall"`
	DateOfBirth                 interface{} `json:"dateOfBirth"`
	DayRate                     int         `json:"dayRate"`
	DayRateLow                  interface{} `json:"dayRateLow"`
	DegreeList                  interface{} `json:"degreeList"`
	Description                 interface{} `json:"description"`
	DesiredLocations            interface{} `json:"desiredLocations"`
	Disability                  interface{} `json:"disability"`
	EducationDegree             interface{} `json:"educationDegree"`
	Educations                  struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"educations"`
	Email                             string      `json:"email"`
	Email2                            interface{} `json:"email2"`
	Email3                            interface{} `json:"email3"`
	EmployeeType                      string      `json:"employeeType"`
	EmploymentPreference              interface{} `json:"employmentPreference"`
	EstaffGUID                        interface{} `json:"estaffGUID"`
	Ethnicity                         interface{} `json:"ethnicity"`
	Experience                        int         `json:"experience"`
	ExternalID                        interface{} `json:"externalID"`
	Fax                               interface{} `json:"fax"`
	Fax2                              interface{} `json:"fax2"`
	Fax3                              interface{} `json:"fax3"`
	FederalAddtionalWitholdingsAmount interface{} `json:"federalAddtionalWitholdingsAmount"`
	FederalExemptions                 interface{} `json:"federalExemptions"`
	FederalExtraWithholdingAmount     interface{} `json:"federalExtraWithholdingAmount"`
	FederalFilingStatus               interface{} `json:"federalFilingStatus"`
	FileAttachments                   struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"fileAttachments"`
	FirstName     string      `json:"firstName"`
	Gender        interface{} `json:"gender"`
	HourlyRate    int         `json:"hourlyRate"`
	HourlyRateLow interface{} `json:"hourlyRateLow"`
	I9OnFile      interface{} `json:"i9OnFile"`
	Interviews    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"interviews"`
	IsAnonymized      bool        `json:"isAnonymized"`
	IsDayLightSavings bool        `json:"isDayLightSavings"`
	IsDeleted         bool        `json:"isDeleted"`
	IsEditable        bool        `json:"isEditable"`
	IsExempt          interface{} `json:"isExempt"`
	IsLockedOut       bool        `json:"isLockedOut"`
	LastName          string      `json:"lastName"`
	Leads             struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"leads"`
	LinkedPerson                    interface{} `json:"linkedPerson"`
	LocalAddtionalWitholdingsAmount interface{} `json:"localAddtionalWitholdingsAmount"`
	LocalExemptions                 interface{} `json:"localExemptions"`
	LocalFilingStatus               interface{} `json:"localFilingStatus"`
	LocalTaxCode                    interface{} `json:"localTaxCode"`
	Locations                       struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"locations"`
	MaritalStatus  interface{} `json:"maritalStatus"`
	MassMailOptOut bool        `json:"massMailOptOut"`
	MasterUserID   interface{} `json:"masterUserID"`
	MiddleName     interface{} `json:"middleName"`
	MigrateGUID    interface{} `json:"migrateGUID"`
	Mobile         string      `json:"mobile"`
	Name           string      `json:"name"`
	NamePrefix     interface{} `json:"namePrefix"`
	NameSuffix     interface{} `json:"nameSuffix"`
	NickName       interface{} `json:"nickName"`
	Notes          struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"notes"`
	NumCategories                   int         `json:"numCategories"`
	NumOwners                       int         `json:"numOwners"`
	Occupation                      interface{} `json:"occupation"`
	OnboardingDocumentReceivedCount int         `json:"onboardingDocumentReceivedCount"`
	OnboardingDocumentSentCount     int         `json:"onboardingDocumentSentCount"`
	OnboardingPercentComplete       int         `json:"onboardingPercentComplete"`
	OnboardingReceivedSent          struct {
		OnboardingDocumentReceivedCount int `json:"onboardingDocumentReceivedCount"`
		OnboardingDocumentSentCount     int `json:"onboardingDocumentSentCount"`
	} `json:"onboardingReceivedSent"`
	OnboardingStatus      interface{} `json:"onboardingStatus"`
	OtherDeductionsAmount interface{} `json:"otherDeductionsAmount"`
	OtherIncomeAmount     interface{} `json:"otherIncomeAmount"`
	Owner                 struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"owner"`
	Pager                  interface{} `json:"pager"`
	PaperWorkOnFile        interface{} `json:"paperWorkOnFile"`
	Password               string      `json:"password"`
	PayrollClientStartDate interface{} `json:"payrollClientStartDate"`
	PayrollStatus          interface{} `json:"payrollStatus"`
	PersonSubtype          string      `json:"personSubtype"`
	Phone                  interface{} `json:"phone"`
	Phone2                 interface{} `json:"phone2"`
	Phone3                 interface{} `json:"phone3"`
	Placements             struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"placements"`
	PreferredContact string `json:"preferredContact"`
	PrimarySkills    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"primarySkills"`
	RecentClientList interface{} `json:"recentClientList"`
	References       struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"references"`
	ReferredBy       interface{} `json:"referredBy"`
	ReferredByPerson interface{} `json:"referredByPerson"`
	Salary           int         `json:"salary"`
	SalaryLow        interface{} `json:"salaryLow"`
	SecondaryAddress struct {
		Address1    interface{} `json:"address1"`
		Address2    interface{} `json:"address2"`
		City        interface{} `json:"city"`
		State       interface{} `json:"state"`
		Zip         interface{} `json:"zip"`
		CountryID   int         `json:"countryID"`
		CountryName string      `json:"countryName"`
		CountryCode string      `json:"countryCode"`
	} `json:"secondaryAddress"`
	SecondaryOwners struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"secondaryOwners"`
	SecondarySkills struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"secondarySkills"`
	Sendouts struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"sendouts"`
	Shifts struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"shifts"`
	SkillSet    interface{} `json:"skillSet"`
	Source      string      `json:"source"`
	Specialties struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"specialties"`
	Ssn                             interface{} `json:"ssn"`
	StateAddtionalWitholdingsAmount interface{} `json:"stateAddtionalWitholdingsAmount"`
	StateExemptions                 interface{} `json:"stateExemptions"`
	StateFilingStatus               interface{} `json:"stateFilingStatus"`
	Status                          string      `json:"status"`
	Submissions                     struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"submissions"`
	Tasks struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tasks"`
	TaxID      interface{} `json:"taxID"`
	TaxState   interface{} `json:"taxState"`
	Tearsheets struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tearsheets"`
	TimeZoneOffsetEST         int         `json:"timeZoneOffsetEST"`
	TobaccoUser               interface{} `json:"tobaccoUser"`
	TotalDependentClaimAmount interface{} `json:"totalDependentClaimAmount"`
	TravelLimit               int         `json:"travelLimit"`
	TravelMethod              interface{} `json:"travelMethod"`
	TwoJobs                   interface{} `json:"twoJobs"`
	Type                      interface{} `json:"type"`
	UserDateAdded             int64       `json:"userDateAdded"`
	UserType                  interface{} `json:"userType"`
	Username                  string      `json:"username"`
	Veteran                   interface{} `json:"veteran"`
	WebResponses              struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"webResponses"`
	WillRelocate   bool `json:"willRelocate"`
	WorkAuthorized bool `json:"workAuthorized"`
	WorkHistories  struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"workHistories"`
	WorkPhone interface{} `json:"workPhone"`
}

const CertificationEntity = "Certification"

type Certification struct {
	Id                  int         `json:"id"`
	Name                string      `json:"name"`
	Category            interface{} `json:"category"`
	CertificationGroups struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"certificationGroups"`
	CountryID              interface{} `json:"countryID"`
	CustomDate1            interface{} `json:"customDate1"`
	CustomDate2            interface{} `json:"customDate2"`
	CustomDate3            interface{} `json:"customDate3"`
	CustomFloat1           interface{} `json:"customFloat1"`
	CustomFloat2           interface{} `json:"customFloat2"`
	CustomFloat3           interface{} `json:"customFloat3"`
	CustomInt1             interface{} `json:"customInt1"`
	CustomInt2             interface{} `json:"customInt2"`
	CustomInt3             interface{} `json:"customInt3"`
	CustomText1            interface{} `json:"customText1"`
	CustomText10           interface{} `json:"customText10"`
	CustomText2            interface{} `json:"customText2"`
	CustomText3            interface{} `json:"customText3"`
	CustomText4            interface{} `json:"customText4"`
	CustomText5            interface{} `json:"customText5"`
	CustomText6            interface{} `json:"customText6"`
	CustomText7            interface{} `json:"customText7"`
	CustomText8            interface{} `json:"customText8"`
	CustomText9            interface{} `json:"customText9"`
	CustomTextBlock1       interface{} `json:"customTextBlock1"`
	CustomTextBlock2       interface{} `json:"customTextBlock2"`
	CustomTextBlock3       interface{} `json:"customTextBlock3"`
	CustomTextBlock4       interface{} `json:"customTextBlock4"`
	CustomTextBlock5       interface{} `json:"customTextBlock5"`
	DateAdded              int64       `json:"dateAdded"`
	DateLastModified       int64       `json:"dateLastModified"`
	Description            interface{} `json:"description"`
	ExpirationDateOptional bool        `json:"expirationDateOptional"`
	MigrateGUID            interface{} `json:"migrateGUID"`
	Specialty              interface{} `json:"specialty"`
	State                  interface{} `json:"state"`
}

const ClientCorporationEntity = "ClientCorporation"

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

const CandidateCertificationEntity = "CandidateCertification"

type CandidateCertification struct {
	BoardCertification interface{} `json:"boardCertification"`
	Candidate          struct {
		FirstName string `json:"firstName"`
		Id        int    `json:"id"`
		LastName  string `json:"lastName"`
	} `json:"candidate"`
	Certification struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"certification"`
	CertificationFileAttachments struct {
		Data  []interface{} `json:"data"`
		Total int           `json:"total"`
	} `json:"certificationFileAttachments"`
	Comments               interface{} `json:"comments"`
	Compact                interface{} `json:"compact"`
	CopyOnFile             interface{} `json:"copyOnFile"`
	CustomDate1            interface{} `json:"customDate1"`
	CustomDate10           interface{} `json:"customDate10"`
	CustomDate2            interface{} `json:"customDate2"`
	CustomDate3            interface{} `json:"customDate3"`
	CustomDate4            interface{} `json:"customDate4"`
	CustomDate5            interface{} `json:"customDate5"`
	CustomDate6            interface{} `json:"customDate6"`
	CustomDate7            interface{} `json:"customDate7"`
	CustomDate8            interface{} `json:"customDate8"`
	CustomDate9            interface{} `json:"customDate9"`
	CustomText1            interface{} `json:"customText1"`
	CustomText10           interface{} `json:"customText10"`
	CustomText2            interface{} `json:"customText2"`
	CustomText3            interface{} `json:"customText3"`
	CustomText4            interface{} `json:"customText4"`
	CustomText5            interface{} `json:"customText5"`
	CustomText6            interface{} `json:"customText6"`
	CustomText7            interface{} `json:"customText7"`
	CustomText8            interface{} `json:"customText8"`
	CustomText9            interface{} `json:"customText9"`
	CustomTextBlock1       interface{} `json:"customTextBlock1"`
	CustomTextBlock10      interface{} `json:"customTextBlock10"`
	CustomTextBlock2       interface{} `json:"customTextBlock2"`
	CustomTextBlock3       interface{} `json:"customTextBlock3"`
	CustomTextBlock4       interface{} `json:"customTextBlock4"`
	CustomTextBlock5       interface{} `json:"customTextBlock5"`
	CustomTextBlock6       interface{} `json:"customTextBlock6"`
	CustomTextBlock7       interface{} `json:"customTextBlock7"`
	CustomTextBlock8       interface{} `json:"customTextBlock8"`
	CustomTextBlock9       interface{} `json:"customTextBlock9"`
	DateAdded              int64       `json:"dateAdded"`
	DateCertified          interface{} `json:"dateCertified"`
	DateExpiration         int64       `json:"dateExpiration"`
	DateLastModified       int64       `json:"dateLastModified"`
	DisplayStatus          string      `json:"displayStatus"`
	ExpirationReminderDate interface{} `json:"expirationReminderDate"`
	FileAttachments        struct {
		Data  []interface{} `json:"data"`
		Total int           `json:"total"`
	} `json:"fileAttachments"`
	Id            int         `json:"id"`
	IsComplete    bool        `json:"isComplete"`
	IsDeleted     bool        `json:"isDeleted"`
	IssuedBy      string      `json:"issuedBy"`
	LicenseNumber string      `json:"licenseNumber"`
	LicenseType   interface{} `json:"licenseType"`
	Location      interface{} `json:"location"`
	MigrateGUID   interface{} `json:"migrateGUID"`
	ModifyingUser struct {
		FirstName string `json:"firstName"`
		Id        int    `json:"id"`
		LastName  string `json:"lastName"`
	} `json:"modifyingUser"`
	Name    string      `json:"name"`
	Results interface{} `json:"results"`
	Status  string      `json:"status"`
}

const ClientContactEntity = "ClientContact"

type ClientContact struct {
	Id               int `json:"id"`
	ActivePlacements struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"activePlacements"`
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
	Appointments struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"appointments"`
	Branch          interface{} `json:"branch"`
	BusinessSectors struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"businessSectors"`
	Categories struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"categories"`
	Category struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"category"`
	Certifications    interface{} `json:"certifications"`
	ClientContactID   int         `json:"clientContactID"`
	ClientCorporation struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"clientCorporation"`
	ClientLocations struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"clientLocations"`
	Comments           interface{} `json:"comments"`
	CompanyName        string      `json:"companyName"`
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
	DateLastComment    int64       `json:"dateLastComment"`
	DateLastModified   int64       `json:"dateLastModified"`
	DateLastVisit      interface{} `json:"dateLastVisit"`
	DeleteMe           interface{} `json:"deleteMe"`
	Description        interface{} `json:"description"`
	DesiredCategories  interface{} `json:"desiredCategories"`
	DesiredLocations   interface{} `json:"desiredLocations"`
	DesiredSkills      interface{} `json:"desiredSkills"`
	DesiredSpecialties interface{} `json:"desiredSpecialties"`
	Division           string      `json:"division"`
	Email              string      `json:"email"`
	Email2             interface{} `json:"email2"`
	Email3             interface{} `json:"email3"`
	ExternalID         interface{} `json:"externalID"`
	Fax                interface{} `json:"fax"`
	Fax2               interface{} `json:"fax2"`
	Fax3               interface{} `json:"fax3"`
	FileAttachments    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"fileAttachments"`
	FirstName  string `json:"firstName"`
	Interviews struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"interviews"`
	IsAnonymized      bool `json:"isAnonymized"`
	IsDayLightSavings bool `json:"isDayLightSavings"`
	IsDefaultContact  bool `json:"isDefaultContact"`
	IsDeleted         bool `json:"isDeleted"`
	IsLockedOut       bool `json:"isLockedOut"`
	JobOrders         struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"jobOrders"`
	JobSubmissions struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"jobSubmissions"`
	LastName string `json:"lastName"`
	Leads    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"leads"`
	LinkedPerson   interface{} `json:"linkedPerson"`
	MassMailOptOut bool        `json:"massMailOptOut"`
	MasterUserID   interface{} `json:"masterUserID"`
	MiddleName     interface{} `json:"middleName"`
	MigrateGUID    interface{} `json:"migrateGUID"`
	Mobile         string      `json:"mobile"`
	Name           string      `json:"name"`
	NamePrefix     interface{} `json:"namePrefix"`
	NameSuffix     interface{} `json:"nameSuffix"`
	NickName       interface{} `json:"nickName"`
	Notes          struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"notes"`
	NumEmployees  int         `json:"numEmployees"`
	Occupation    string      `json:"occupation"`
	Office        interface{} `json:"office"`
	Opportunities struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"opportunities"`
	Owner struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"owner"`
	Pager         interface{} `json:"pager"`
	Password      string      `json:"password"`
	PersonSubtype string      `json:"personSubtype"`
	Phone         string      `json:"phone"`
	Phone2        interface{} `json:"phone2"`
	Phone3        interface{} `json:"phone3"`
	Placements    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"placements"`
	PreferredContact string      `json:"preferredContact"`
	ReferredByPerson interface{} `json:"referredByPerson"`
	ReportToPerson   interface{} `json:"reportToPerson"`
	SecondaryAddress struct {
		Address1    interface{} `json:"address1"`
		Address2    interface{} `json:"address2"`
		City        interface{} `json:"city"`
		State       interface{} `json:"state"`
		Zip         interface{} `json:"zip"`
		CountryID   int         `json:"countryID"`
		CountryName string      `json:"countryName"`
		CountryCode string      `json:"countryCode"`
	} `json:"secondaryAddress"`
	SecondaryOwners struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"secondaryOwners"`
	Sendouts struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"sendouts"`
	SkillSet interface{} `json:"skillSet"`
	Skills   struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"skills"`
	Source      interface{} `json:"source"`
	Specialties struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"specialties"`
	Status string `json:"status"`
	Tasks  struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tasks"`
	Tearsheets struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tearsheets"`
	TimeZoneOffsetEST int         `json:"timeZoneOffsetEST"`
	TrackTitle        interface{} `json:"trackTitle"`
	Type              string      `json:"type"`
	UserDateAdded     int64       `json:"userDateAdded"`
	UserType          interface{} `json:"userType"`
	Username          string      `json:"username"`
}

const LocationEntity = "Location"

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

const JobOrderEntity = "JobOrder"

type JobOrder struct {
	Id      int `json:"id"`
	Address struct {
		Address1    string      `json:"address1"`
		Address2    interface{} `json:"address2"`
		City        string      `json:"city"`
		CountryCode string      `json:"countryCode"`
		CountryID   int         `json:"countryID"`
		CountryName string      `json:"countryName"`
		State       string      `json:"state"`
		Timezone    string      `json:"timezone"`
		Zip         string      `json:"zip"`
	} `json:"address"`
	Appointments struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"appointments"`
	ApprovedPlacements struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"approvedPlacements"`
	AssignedUsers struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"assignedUsers"`
	Benefits           interface{} `json:"benefits"`
	BillRateCategoryID interface{} `json:"billRateCategoryID"`
	BonusPackage       interface{} `json:"bonusPackage"`
	Branch             interface{} `json:"branch"`
	BranchCode         interface{} `json:"branchCode"`
	BusinessSectors    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"businessSectors"`
	Categories struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"categories"`
	CertificationGroups struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"certificationGroups"`
	CertificationList interface{} `json:"certificationList"`
	Certifications    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"certifications"`
	ClientBillRate interface{} `json:"clientBillRate"`
	ClientContact  struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"clientContact"`
	ClientCorporation struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"clientCorporation"`
	CorrelatedCustomDate1      interface{}   `json:"correlatedCustomDate1"`
	CorrelatedCustomDate2      interface{}   `json:"correlatedCustomDate2"`
	CorrelatedCustomDate3      interface{}   `json:"correlatedCustomDate3"`
	CorrelatedCustomFloat1     interface{}   `json:"correlatedCustomFloat1"`
	CorrelatedCustomFloat2     interface{}   `json:"correlatedCustomFloat2"`
	CorrelatedCustomFloat3     interface{}   `json:"correlatedCustomFloat3"`
	CorrelatedCustomInt1       interface{}   `json:"correlatedCustomInt1"`
	CorrelatedCustomInt2       interface{}   `json:"correlatedCustomInt2"`
	CorrelatedCustomInt3       interface{}   `json:"correlatedCustomInt3"`
	CorrelatedCustomText1      interface{}   `json:"correlatedCustomText1"`
	CorrelatedCustomText10     interface{}   `json:"correlatedCustomText10"`
	CorrelatedCustomText2      interface{}   `json:"correlatedCustomText2"`
	CorrelatedCustomText3      interface{}   `json:"correlatedCustomText3"`
	CorrelatedCustomText4      interface{}   `json:"correlatedCustomText4"`
	CorrelatedCustomText5      interface{}   `json:"correlatedCustomText5"`
	CorrelatedCustomText6      interface{}   `json:"correlatedCustomText6"`
	CorrelatedCustomText7      interface{}   `json:"correlatedCustomText7"`
	CorrelatedCustomText8      interface{}   `json:"correlatedCustomText8"`
	CorrelatedCustomText9      interface{}   `json:"correlatedCustomText9"`
	CorrelatedCustomTextBlock1 interface{}   `json:"correlatedCustomTextBlock1"`
	CorrelatedCustomTextBlock2 interface{}   `json:"correlatedCustomTextBlock2"`
	CorrelatedCustomTextBlock3 interface{}   `json:"correlatedCustomTextBlock3"`
	CostCenter                 interface{}   `json:"costCenter"`
	CustomDate1                interface{}   `json:"customDate1"`
	CustomDate2                interface{}   `json:"customDate2"`
	CustomDate3                interface{}   `json:"customDate3"`
	CustomFloat1               interface{}   `json:"customFloat1"`
	CustomFloat2               interface{}   `json:"customFloat2"`
	CustomFloat3               interface{}   `json:"customFloat3"`
	CustomInt1                 interface{}   `json:"customInt1"`
	CustomInt2                 interface{}   `json:"customInt2"`
	CustomInt3                 interface{}   `json:"customInt3"`
	CustomInt4                 interface{}   `json:"customInt4"`
	CustomInt5                 interface{}   `json:"customInt5"`
	CustomInt6                 interface{}   `json:"customInt6"`
	CustomInt7                 interface{}   `json:"customInt7"`
	CustomInt8                 interface{}   `json:"customInt8"`
	CustomText1                interface{}   `json:"customText1"`
	CustomText10               interface{}   `json:"customText10"`
	CustomText11               interface{}   `json:"customText11"`
	CustomText12               interface{}   `json:"customText12"`
	CustomText13               interface{}   `json:"customText13"`
	CustomText14               interface{}   `json:"customText14"`
	CustomText15               interface{}   `json:"customText15"`
	CustomText16               interface{}   `json:"customText16"`
	CustomText17               interface{}   `json:"customText17"`
	CustomText18               interface{}   `json:"customText18"`
	CustomText19               interface{}   `json:"customText19"`
	CustomText2                []interface{} `json:"customText2"`
	CustomText20               interface{}   `json:"customText20"`
	CustomText21               interface{}   `json:"customText21"`
	CustomText22               interface{}   `json:"customText22"`
	CustomText23               interface{}   `json:"customText23"`
	CustomText24               interface{}   `json:"customText24"`
	CustomText25               interface{}   `json:"customText25"`
	CustomText26               interface{}   `json:"customText26"`
	CustomText27               interface{}   `json:"customText27"`
	CustomText28               interface{}   `json:"customText28"`
	CustomText29               interface{}   `json:"customText29"`
	CustomText3                interface{}   `json:"customText3"`
	CustomText30               interface{}   `json:"customText30"`
	CustomText31               interface{}   `json:"customText31"`
	CustomText32               interface{}   `json:"customText32"`
	CustomText33               interface{}   `json:"customText33"`
	CustomText34               interface{}   `json:"customText34"`
	CustomText35               interface{}   `json:"customText35"`
	CustomText36               interface{}   `json:"customText36"`
	CustomText37               interface{}   `json:"customText37"`
	CustomText38               interface{}   `json:"customText38"`
	CustomText39               interface{}   `json:"customText39"`
	CustomText4                interface{}   `json:"customText4"`
	CustomText40               interface{}   `json:"customText40"`
	CustomText5                []interface{} `json:"customText5"`
	CustomText6                []interface{} `json:"customText6"`
	CustomText7                interface{}   `json:"customText7"`
	CustomText8                interface{}   `json:"customText8"`
	CustomText9                interface{}   `json:"customText9"`
	CustomTextBlock1           interface{}   `json:"customTextBlock1"`
	CustomTextBlock2           interface{}   `json:"customTextBlock2"`
	CustomTextBlock3           interface{}   `json:"customTextBlock3"`
	CustomTextBlock4           interface{}   `json:"customTextBlock4"`
	CustomTextBlock5           interface{}   `json:"customTextBlock5"`
	DateAdded                  int64         `json:"dateAdded"`
	DateClosed                 interface{}   `json:"dateClosed"`
	DateEnd                    int64         `json:"dateEnd"`
	DateLastExported           interface{}   `json:"dateLastExported"`
	DateLastModified           int64         `json:"dateLastModified"`
	DateLastPublished          interface{}   `json:"dateLastPublished"`
	DegreeList                 interface{}   `json:"degreeList"`
	Description                interface{}   `json:"description"`
	DurationWeeks              int           `json:"durationWeeks"`
	EducationDegree            interface{}   `json:"educationDegree"`
	EmploymentType             string        `json:"employmentType"`
	ExternalCategoryID         interface{}   `json:"externalCategoryID"`
	ExternalID                 interface{}   `json:"externalID"`
	FeeArrangement             interface{}   `json:"feeArrangement"`
	FileAttachments            struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"fileAttachments"`
	HoursOfOperation interface{} `json:"hoursOfOperation"`
	HoursPerWeek     int         `json:"hoursPerWeek"`
	Interviews       struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"interviews"`
	IsClientEditable    bool        `json:"isClientEditable"`
	IsDeleted           bool        `json:"isDeleted"`
	IsInterviewRequired bool        `json:"isInterviewRequired"`
	IsJobcastPublished  interface{} `json:"isJobcastPublished"`
	IsOpen              bool        `json:"isOpen"`
	IsPublic            int         `json:"isPublic"`
	IsWorkFromHome      interface{} `json:"isWorkFromHome"`
	JobBoardList        interface{} `json:"jobBoardList"`
	Location            interface{} `json:"location"`
	MarkUpPercentage    int         `json:"markUpPercentage"`
	Notes               struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"notes"`
	NumOpenings    int         `json:"numOpenings"`
	OnSite         string      `json:"onSite"`
	Opportunity    interface{} `json:"opportunity"`
	OptionsPackage interface{} `json:"optionsPackage"`
	Owner          struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"owner"`
	PayRate    int `json:"payRate"`
	Placements struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"placements"`
	PublicDescription     interface{} `json:"publicDescription"`
	PublishedCategory     interface{} `json:"publishedCategory"`
	PublishedZip          interface{} `json:"publishedZip"`
	ReasonClosed          interface{} `json:"reasonClosed"`
	ReportTo              interface{} `json:"reportTo"`
	ReportToClientContact interface{} `json:"reportToClientContact"`
	ResponseUser          interface{} `json:"responseUser"`
	Salary                int         `json:"salary"`
	SalaryUnit            string      `json:"salaryUnit"`
	Sendouts              struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"sendouts"`
	Shift  interface{} `json:"shift"`
	Shifts struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"shifts"`
	SkillList interface{} `json:"skillList"`
	Skills    struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"skills"`
	Source      interface{} `json:"source"`
	Specialties struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"specialties"`
	StartDate   int64  `json:"startDate"`
	Status      string `json:"status"`
	Submissions struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"submissions"`
	Tasks struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tasks"`
	TaxRate    int    `json:"taxRate"`
	TaxStatus  string `json:"taxStatus"`
	Tearsheets struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tearsheets"`
	TimeUnits struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"timeUnits"`
	Title              string      `json:"title"`
	TravelRequirements interface{} `json:"travelRequirements"`
	Type               int         `json:"type"`
	UsersAssigned      interface{} `json:"usersAssigned"`
	WebResponses       struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"webResponses"`
	WillRelocate    bool        `json:"willRelocate"`
	WillRelocateInt int         `json:"willRelocateInt"`
	WillSponsor     bool        `json:"willSponsor"`
	WorkersCompRate interface{} `json:"workersCompRate"`
	YearsRequired   int         `json:"yearsRequired"`
}

const JobSubmissionEntity = "JobSubmission"

type JobSubmission struct {
	Id           int `json:"id"`
	Appointments struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"appointments"`
	BillRate  interface{} `json:"billRate"`
	Branch    interface{} `json:"branch"`
	Candidate struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"candidate"`
	Comments         interface{} `json:"comments"`
	CustomDate1      interface{} `json:"customDate1"`
	CustomDate2      interface{} `json:"customDate2"`
	CustomDate3      interface{} `json:"customDate3"`
	CustomDate4      interface{} `json:"customDate4"`
	CustomDate5      interface{} `json:"customDate5"`
	CustomFloat1     interface{} `json:"customFloat1"`
	CustomFloat2     interface{} `json:"customFloat2"`
	CustomFloat3     interface{} `json:"customFloat3"`
	CustomFloat4     interface{} `json:"customFloat4"`
	CustomFloat5     interface{} `json:"customFloat5"`
	CustomInt1       interface{} `json:"customInt1"`
	CustomInt2       interface{} `json:"customInt2"`
	CustomInt3       interface{} `json:"customInt3"`
	CustomInt4       interface{} `json:"customInt4"`
	CustomInt5       interface{} `json:"customInt5"`
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
	CustomText21     interface{} `json:"customText21"`
	CustomText22     interface{} `json:"customText22"`
	CustomText23     interface{} `json:"customText23"`
	CustomText24     interface{} `json:"customText24"`
	CustomText25     interface{} `json:"customText25"`
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
	CustomTextBlock4 interface{} `json:"customTextBlock4"`
	CustomTextBlock5 interface{} `json:"customTextBlock5"`
	DateAdded        int64       `json:"dateAdded"`
	DateLastModified int64       `json:"dateLastModified"`
	DateWebResponse  int64       `json:"dateWebResponse"`
	EndDate          interface{} `json:"endDate"`
	IsDeleted        bool        `json:"isDeleted"`
	IsHidden         bool        `json:"isHidden"`
	JobOrder         struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"jobOrder"`
	JobSubmissionCertificationRequirements struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"jobSubmissionCertificationRequirements"`
	LatestAppointment interface{} `json:"latestAppointment"`
	MigrateGUID       interface{} `json:"migrateGUID"`
	Owners            struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"owners"`
	PayRate     interface{} `json:"payRate"`
	Salary      interface{} `json:"salary"`
	SendingUser struct {
		Id        int    `json:"id"`
		Subtype   string `json:"_subtype"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"sendingUser"`
	Source    string      `json:"source"`
	StartDate interface{} `json:"startDate"`
	Status    string      `json:"status"`
	Tasks     struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tasks"`
}

const PlacementEntity = "Placement"

type Placement struct {
	Id           int `json:"id"`
	Appointments struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"appointments"`
	ApprovedChangeRequests       int         `json:"approvedChangeRequests"`
	ApprovingClientContact       interface{} `json:"approvingClientContact"`
	BackupApprovingClientContact interface{} `json:"backupApprovingClientContact"`
	BenefitGroup                 interface{} `json:"benefitGroup"`
	BillingClientContact         interface{} `json:"billingClientContact"`
	BillingFrequency             interface{} `json:"billingFrequency"`
	BonusPackage                 interface{} `json:"bonusPackage"`
	Branch                       interface{} `json:"branch"`
	CanEnterTime                 bool        `json:"canEnterTime"`
	Candidate                    struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"candidate"`
	ChangeRequests struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"changeRequests"`
	ClientBillRate interface{} `json:"clientBillRate"`
	ClientContact  struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"clientContact"`
	ClientCorporation struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"clientCorporation"`
	ClientOvertimeRate interface{} `json:"clientOvertimeRate"`
	Comments           interface{} `json:"comments"`
	Commissions        struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"commissions"`
	CorrelatedCustomDate1      interface{} `json:"correlatedCustomDate1"`
	CorrelatedCustomDate2      interface{} `json:"correlatedCustomDate2"`
	CorrelatedCustomDate3      interface{} `json:"correlatedCustomDate3"`
	CorrelatedCustomFloat1     interface{} `json:"correlatedCustomFloat1"`
	CorrelatedCustomFloat2     interface{} `json:"correlatedCustomFloat2"`
	CorrelatedCustomFloat3     interface{} `json:"correlatedCustomFloat3"`
	CorrelatedCustomInt1       interface{} `json:"correlatedCustomInt1"`
	CorrelatedCustomInt2       interface{} `json:"correlatedCustomInt2"`
	CorrelatedCustomInt3       interface{} `json:"correlatedCustomInt3"`
	CorrelatedCustomText1      interface{} `json:"correlatedCustomText1"`
	CorrelatedCustomText10     interface{} `json:"correlatedCustomText10"`
	CorrelatedCustomText2      interface{} `json:"correlatedCustomText2"`
	CorrelatedCustomText3      interface{} `json:"correlatedCustomText3"`
	CorrelatedCustomText4      interface{} `json:"correlatedCustomText4"`
	CorrelatedCustomText5      interface{} `json:"correlatedCustomText5"`
	CorrelatedCustomText6      interface{} `json:"correlatedCustomText6"`
	CorrelatedCustomText7      interface{} `json:"correlatedCustomText7"`
	CorrelatedCustomText8      interface{} `json:"correlatedCustomText8"`
	CorrelatedCustomText9      interface{} `json:"correlatedCustomText9"`
	CorrelatedCustomTextBlock1 interface{} `json:"correlatedCustomTextBlock1"`
	CorrelatedCustomTextBlock2 interface{} `json:"correlatedCustomTextBlock2"`
	CorrelatedCustomTextBlock3 interface{} `json:"correlatedCustomTextBlock3"`
	CostCenter                 interface{} `json:"costCenter"`
	CustomBillRate1            interface{} `json:"customBillRate1"`
	CustomBillRate10           interface{} `json:"customBillRate10"`
	CustomBillRate2            interface{} `json:"customBillRate2"`
	CustomBillRate3            interface{} `json:"customBillRate3"`
	CustomBillRate4            interface{} `json:"customBillRate4"`
	CustomBillRate5            interface{} `json:"customBillRate5"`
	CustomBillRate6            interface{} `json:"customBillRate6"`
	CustomBillRate7            interface{} `json:"customBillRate7"`
	CustomBillRate8            interface{} `json:"customBillRate8"`
	CustomBillRate9            interface{} `json:"customBillRate9"`
	CustomDate1                interface{} `json:"customDate1"`
	CustomDate10               interface{} `json:"customDate10"`
	CustomDate11               interface{} `json:"customDate11"`
	CustomDate12               interface{} `json:"customDate12"`
	CustomDate13               interface{} `json:"customDate13"`
	CustomDate2                interface{} `json:"customDate2"`
	CustomDate3                interface{} `json:"customDate3"`
	CustomDate4                interface{} `json:"customDate4"`
	CustomDate5                interface{} `json:"customDate5"`
	CustomDate6                interface{} `json:"customDate6"`
	CustomDate7                interface{} `json:"customDate7"`
	CustomDate8                interface{} `json:"customDate8"`
	CustomDate9                interface{} `json:"customDate9"`
	CustomEncryptedText1       interface{} `json:"customEncryptedText1"`
	CustomEncryptedText10      interface{} `json:"customEncryptedText10"`
	CustomEncryptedText2       interface{} `json:"customEncryptedText2"`
	CustomEncryptedText3       interface{} `json:"customEncryptedText3"`
	CustomEncryptedText4       interface{} `json:"customEncryptedText4"`
	CustomEncryptedText5       interface{} `json:"customEncryptedText5"`
	CustomEncryptedText6       interface{} `json:"customEncryptedText6"`
	CustomEncryptedText7       interface{} `json:"customEncryptedText7"`
	CustomEncryptedText8       interface{} `json:"customEncryptedText8"`
	CustomEncryptedText9       interface{} `json:"customEncryptedText9"`
	CustomFloat1               interface{} `json:"customFloat1"`
	CustomFloat10              interface{} `json:"customFloat10"`
	CustomFloat11              interface{} `json:"customFloat11"`
	CustomFloat12              interface{} `json:"customFloat12"`
	CustomFloat13              interface{} `json:"customFloat13"`
	CustomFloat14              interface{} `json:"customFloat14"`
	CustomFloat15              interface{} `json:"customFloat15"`
	CustomFloat16              interface{} `json:"customFloat16"`
	CustomFloat17              interface{} `json:"customFloat17"`
	CustomFloat18              interface{} `json:"customFloat18"`
	CustomFloat19              interface{} `json:"customFloat19"`
	CustomFloat2               interface{} `json:"customFloat2"`
	CustomFloat20              interface{} `json:"customFloat20"`
	CustomFloat21              interface{} `json:"customFloat21"`
	CustomFloat22              interface{} `json:"customFloat22"`
	CustomFloat23              interface{} `json:"customFloat23"`
	CustomFloat3               interface{} `json:"customFloat3"`
	CustomFloat4               interface{} `json:"customFloat4"`
	CustomFloat5               interface{} `json:"customFloat5"`
	CustomFloat6               interface{} `json:"customFloat6"`
	CustomFloat7               interface{} `json:"customFloat7"`
	CustomFloat8               interface{} `json:"customFloat8"`
	CustomFloat9               interface{} `json:"customFloat9"`
	CustomInt1                 interface{} `json:"customInt1"`
	CustomInt10                interface{} `json:"customInt10"`
	CustomInt11                interface{} `json:"customInt11"`
	CustomInt12                interface{} `json:"customInt12"`
	CustomInt13                interface{} `json:"customInt13"`
	CustomInt14                interface{} `json:"customInt14"`
	CustomInt15                interface{} `json:"customInt15"`
	CustomInt16                interface{} `json:"customInt16"`
	CustomInt17                interface{} `json:"customInt17"`
	CustomInt18                interface{} `json:"customInt18"`
	CustomInt19                interface{} `json:"customInt19"`
	CustomInt2                 interface{} `json:"customInt2"`
	CustomInt20                interface{} `json:"customInt20"`
	CustomInt21                interface{} `json:"customInt21"`
	CustomInt22                interface{} `json:"customInt22"`
	CustomInt23                interface{} `json:"customInt23"`
	CustomInt3                 interface{} `json:"customInt3"`
	CustomInt4                 interface{} `json:"customInt4"`
	CustomInt5                 interface{} `json:"customInt5"`
	CustomInt6                 interface{} `json:"customInt6"`
	CustomInt7                 interface{} `json:"customInt7"`
	CustomInt8                 interface{} `json:"customInt8"`
	CustomInt9                 interface{} `json:"customInt9"`
	CustomPayRate1             interface{} `json:"customPayRate1"`
	CustomPayRate10            interface{} `json:"customPayRate10"`
	CustomPayRate2             interface{} `json:"customPayRate2"`
	CustomPayRate3             interface{} `json:"customPayRate3"`
	CustomPayRate4             interface{} `json:"customPayRate4"`
	CustomPayRate5             interface{} `json:"customPayRate5"`
	CustomPayRate6             interface{} `json:"customPayRate6"`
	CustomPayRate7             interface{} `json:"customPayRate7"`
	CustomPayRate8             interface{} `json:"customPayRate8"`
	CustomPayRate9             interface{} `json:"customPayRate9"`
	CustomText1                interface{} `json:"customText1"`
	CustomText10               interface{} `json:"customText10"`
	CustomText11               interface{} `json:"customText11"`
	CustomText12               interface{} `json:"customText12"`
	CustomText13               interface{} `json:"customText13"`
	CustomText14               interface{} `json:"customText14"`
	CustomText15               interface{} `json:"customText15"`
	CustomText16               interface{} `json:"customText16"`
	CustomText17               interface{} `json:"customText17"`
	CustomText18               interface{} `json:"customText18"`
	CustomText19               interface{} `json:"customText19"`
	CustomText2                interface{} `json:"customText2"`
	CustomText20               interface{} `json:"customText20"`
	CustomText21               interface{} `json:"customText21"`
	CustomText22               interface{} `json:"customText22"`
	CustomText23               interface{} `json:"customText23"`
	CustomText24               interface{} `json:"customText24"`
	CustomText25               interface{} `json:"customText25"`
	CustomText26               interface{} `json:"customText26"`
	CustomText27               interface{} `json:"customText27"`
	CustomText28               interface{} `json:"customText28"`
	CustomText29               interface{} `json:"customText29"`
	CustomText3                interface{} `json:"customText3"`
	CustomText30               interface{} `json:"customText30"`
	CustomText31               interface{} `json:"customText31"`
	CustomText32               interface{} `json:"customText32"`
	CustomText33               interface{} `json:"customText33"`
	CustomText34               interface{} `json:"customText34"`
	CustomText35               interface{} `json:"customText35"`
	CustomText36               interface{} `json:"customText36"`
	CustomText37               interface{} `json:"customText37"`
	CustomText38               interface{} `json:"customText38"`
	CustomText39               interface{} `json:"customText39"`
	CustomText4                interface{} `json:"customText4"`
	CustomText40               interface{} `json:"customText40"`
	CustomText41               interface{} `json:"customText41"`
	CustomText42               interface{} `json:"customText42"`
	CustomText43               interface{} `json:"customText43"`
	CustomText44               interface{} `json:"customText44"`
	CustomText45               interface{} `json:"customText45"`
	CustomText46               interface{} `json:"customText46"`
	CustomText47               interface{} `json:"customText47"`
	CustomText48               interface{} `json:"customText48"`
	CustomText49               interface{} `json:"customText49"`
	CustomText5                interface{} `json:"customText5"`
	CustomText50               interface{} `json:"customText50"`
	CustomText51               interface{} `json:"customText51"`
	CustomText52               interface{} `json:"customText52"`
	CustomText53               interface{} `json:"customText53"`
	CustomText54               interface{} `json:"customText54"`
	CustomText55               interface{} `json:"customText55"`
	CustomText56               interface{} `json:"customText56"`
	CustomText57               interface{} `json:"customText57"`
	CustomText58               interface{} `json:"customText58"`
	CustomText59               interface{} `json:"customText59"`
	CustomText6                interface{} `json:"customText6"`
	CustomText60               interface{} `json:"customText60"`
	CustomText7                interface{} `json:"customText7"`
	CustomText8                interface{} `json:"customText8"`
	CustomText9                interface{} `json:"customText9"`
	CustomTextBlock1           interface{} `json:"customTextBlock1"`
	CustomTextBlock10          interface{} `json:"customTextBlock10"`
	CustomTextBlock2           interface{} `json:"customTextBlock2"`
	CustomTextBlock3           interface{} `json:"customTextBlock3"`
	CustomTextBlock4           interface{} `json:"customTextBlock4"`
	CustomTextBlock5           interface{} `json:"customTextBlock5"`
	CustomTextBlock6           interface{} `json:"customTextBlock6"`
	CustomTextBlock7           interface{} `json:"customTextBlock7"`
	CustomTextBlock8           interface{} `json:"customTextBlock8"`
	CustomTextBlock9           interface{} `json:"customTextBlock9"`
	DateAdded                  int64       `json:"dateAdded"`
	DateBegin                  int64       `json:"dateBegin"`
	DateClientEffective        int64       `json:"dateClientEffective"`
	DateEffective              int64       `json:"dateEffective"`
	DateEnd                    interface{} `json:"dateEnd"`
	DateLastModified           int64       `json:"dateLastModified"`
	DaysGuaranteed             int         `json:"daysGuaranteed"`
	DaysProRated               int         `json:"daysProRated"`
	DurationWeeks              float64     `json:"durationWeeks"`
	EmployeeType               interface{} `json:"employeeType"`
	EmploymentStartDate        interface{} `json:"employmentStartDate"`
	EmploymentType             string      `json:"employmentType"`
	EstaffGUID                 interface{} `json:"estaffGUID"`
	ExpiringCredentials        int         `json:"expiringCredentials"`
	Fee                        float64     `json:"fee"`
	FileAttachments            struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"fileAttachments"`
	FlatFee          float64     `json:"flatFee"`
	HoursOfOperation interface{} `json:"hoursOfOperation"`
	HoursPerDay      float64     `json:"hoursPerDay"`
	HousingAmenities struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"housingAmenities"`
	HousingManagerID       interface{} `json:"housingManagerID"`
	HousingStatus          interface{} `json:"housingStatus"`
	IncompleteRequirements int         `json:"incompleteRequirements"`
	InvoiceGroupName       interface{} `json:"invoiceGroupName"`
	IsMultirate            bool        `json:"isMultirate"`
	IsWorkFromHome         interface{} `json:"isWorkFromHome"`
	JobOrder               struct {
		Id    int    `json:"id"`
		Title string `json:"title"`
	} `json:"jobOrder"`
	JobSubmission struct {
		Id int `json:"id"`
	} `json:"jobSubmission"`
	LastApprovedPlacementChangeRequest interface{} `json:"lastApprovedPlacementChangeRequest"`
	Location                           interface{} `json:"location"`
	MarkUpPercentage                   float64     `json:"markUpPercentage"`
	MigrateGUID                        interface{} `json:"migrateGUID"`
	Notes                              struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"notes"`
	OnboardingDocumentReceivedCount int `json:"onboardingDocumentReceivedCount"`
	OnboardingDocumentSentCount     int `json:"onboardingDocumentSentCount"`
	OnboardingPercentComplete       int `json:"onboardingPercentComplete"`
	OnboardingReceivedSent          struct {
		OnboardingDocumentReceivedCount int `json:"onboardingDocumentReceivedCount"`
		OnboardingDocumentSentCount     int `json:"onboardingDocumentSentCount"`
	} `json:"onboardingReceivedSent"`
	OnboardingStatus         interface{} `json:"onboardingStatus"`
	OptionsPackage           interface{} `json:"optionsPackage"`
	OtExemption              interface{} `json:"otExemption"`
	OtherHourlyFee           interface{} `json:"otherHourlyFee"`
	OtherHourlyFeeComments   interface{} `json:"otherHourlyFeeComments"`
	OvertimeMarkUpPercentage float64     `json:"overtimeMarkUpPercentage"`
	OvertimeRate             interface{} `json:"overtimeRate"`
	Owner                    struct {
		Id        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	} `json:"owner"`
	Owners struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"owners"`
	PayGroup                interface{} `json:"payGroup"`
	PayRate                 float64     `json:"payRate"`
	PayrollEmployeeType     interface{} `json:"payrollEmployeeType"`
	PayrollSyncStatus       interface{} `json:"payrollSyncStatus"`
	PendingChangeRequests   int         `json:"pendingChangeRequests"`
	PlacementCertifications struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"placementCertifications"`
	PositionCode                        interface{} `json:"positionCode"`
	ProjectCodeList                     interface{} `json:"projectCodeList"`
	QuitJob                             bool        `json:"quitJob"`
	RecruitingManagerPercentGrossMargin float64     `json:"recruitingManagerPercentGrossMargin"`
	ReferralFee                         float64     `json:"referralFee"`
	ReferralFeeType                     string      `json:"referralFeeType"`
	ReportTo                            interface{} `json:"reportTo"`
	ReportedMargin                      float64     `json:"reportedMargin"`
	Salary                              float64     `json:"salary"`
	SalaryUnit                          string      `json:"salaryUnit"`
	SalesManagerPercentGrossMargin      float64     `json:"salesManagerPercentGrossMargin"`
	Shift                               interface{} `json:"shift"`
	StatementClientContact              interface{} `json:"statementClientContact"`
	Status                              string      `json:"status"`
	Tasks                               struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"tasks"`
	TaxRate           float64     `json:"taxRate"`
	TaxState          interface{} `json:"taxState"`
	TerminationReason interface{} `json:"terminationReason"`
	TimeUnits         struct {
		Total int           `json:"total"`
		Data  []interface{} `json:"data"`
	} `json:"timeUnits"`
	VendorClientCorporation interface{} `json:"vendorClientCorporation"`
	WorkWeekStart           interface{} `json:"workWeekStart"`
	WorkersCompensationRate interface{} `json:"workersCompensationRate"`
}

type SubscribeEventResponse struct {
	LastRequestId  int    `json:"lastRequestId"`
	SubscriptionId string `json:"subscriptionId"`
	CreatedOn      int64  `json:"createdOn"`
	JmsSelector    string `json:"jmsSelector"`
}

type UnsubscribeEventResponse struct {
	Result bool `json:"result"`
}

type FetchEventResponse struct {
	RequestId int `json:"requestId"`
	Events    []struct {
		EventId        string `json:"eventId"`
		EventType      string `json:"eventType"`
		EventTimestamp int64  `json:"eventTimestamp"`
		EventMetadata  struct {
			PersonId      string `json:"PERSON_ID"`
			TransactionId string `json:"TRANSACTION_ID"`
		} `json:"eventMetadata"`
		EntityName        string   `json:"entityName"`
		EntityId          int      `json:"entityId"`
		EntityEventType   string   `json:"entityEventType"`
		UpdatedProperties []string `json:"updatedProperties"`
	} `json:"events"`
}
