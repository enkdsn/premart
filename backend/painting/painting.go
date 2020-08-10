package painting

import "time"

type Painting struct {
	ObjectID          int           `json:"objectID"`
	IsHighlight       bool          `json:"isHighlight"`
	AccessionNumber   string        `json:"accessionNumber"`
	AccessionYear     string        `json:"accessionYear"`
	IsPublicDomain    bool          `json:"isPublicDomain"`
	PrimaryImage      string        `json:"primaryImage"`
	PrimaryImageSmall string        `json:"primaryImageSmall"`
	AdditionalImages  []interface{} `json:"additionalImages"`
	Constituents      []struct {
		Role                   string `json:"role"`
		Name                   string `json:"name"`
		ConstituentULANURL     string `json:"constituentULAN_URL"`
		ConstituentWikidataURL string `json:"constituentWikidata_URL"`
		Gender                 string `json:"gender"`
	} `json:"constituents"`
	Department            string      `json:"department"`
	ObjectName            string      `json:"objectName"`
	Title                 string      `json:"title"`
	Culture               string      `json:"culture"`
	Period                string      `json:"period"`
	Dynasty               string      `json:"dynasty"`
	Reign                 string      `json:"reign"`
	Portfolio             string      `json:"portfolio"`
	ArtistRole            string      `json:"artistRole"`
	ArtistPrefix          string      `json:"artistPrefix"`
	ArtistDisplayName     string      `json:"artistDisplayName"`
	ArtistDisplayBio      string      `json:"artistDisplayBio"`
	ArtistSuffix          string      `json:"artistSuffix"`
	ArtistAlphaSort       string      `json:"artistAlphaSort"`
	ArtistNationality     string      `json:"artistNationality"`
	ArtistBeginDate       string      `json:"artistBeginDate"`
	ArtistEndDate         string      `json:"artistEndDate"`
	ArtistGender          string      `json:"artistGender"`
	ArtistWikidataURL     string      `json:"artistWikidata_URL"`
	ArtistULANURL         string      `json:"artistULAN_URL"`
	ObjectDate            string      `json:"objectDate"`
	ObjectBeginDate       int         `json:"objectBeginDate"`
	ObjectEndDate         int         `json:"objectEndDate"`
	Medium                string      `json:"medium"`
	Dimensions            string      `json:"dimensions"`
	CreditLine            string      `json:"creditLine"`
	GeographyType         string      `json:"geographyType"`
	City                  string      `json:"city"`
	State                 string      `json:"state"`
	County                string      `json:"county"`
	Country               string      `json:"country"`
	Region                string      `json:"region"`
	Subregion             string      `json:"subregion"`
	Locale                string      `json:"locale"`
	Locus                 string      `json:"locus"`
	Excavation            string      `json:"excavation"`
	River                 string      `json:"river"`
	Classification        string      `json:"classification"`
	RightsAndReproduction string      `json:"rightsAndReproduction"`
	LinkResource          string      `json:"linkResource"`
	MetadataDate          time.Time   `json:"metadataDate"`
	Repository            string      `json:"repository"`
	ObjectURL             string      `json:"objectURL"`
	Tags                  interface{} `json:"tags"`
	ObjectWikidataURL     string      `json:"objectWikidata_URL"`
	IsTimelineWork        bool        `json:"isTimelineWork"`
}
