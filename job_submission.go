package bullhorn

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
