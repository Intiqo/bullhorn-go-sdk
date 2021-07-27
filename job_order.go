package bullhorn

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
