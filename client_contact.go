package bullhorn

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
