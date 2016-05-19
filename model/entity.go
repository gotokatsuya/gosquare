package model

type Venue struct {
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Contact      Contact      `json:"contact"`
	Location     Location     `json:"location"`
	Categories   []Category   `json:"categories"`
	Verified     bool         `json:"verified"`
	Stats        Stats        `json:"stats"`
	URL          string       `json:"url"`
	CanonicalURL string       `json:"canonicalUrl"`
	ShortURL     string       `json:"shortUrl"`
	Likes        Likes        `json:"likes"`
	Rating       float32      `json:"rating"`
	Like         bool         `json:"like"`
	Dislike      bool         `json:"dislike"`
	TimeZone     string       `json:"timeZone"`
	FriendVisits FriendVisits `json:"friendVisits"`
	BeenHere     BeenHere     `json:"beenHere"`
	Specials     Specials     `json:"specials"`
	HereNow      HereNow      `json:"hereNow"`
	Tags         []string     `json:"tags"`
	Photos       Photos       `json:"photos"`
	BestPhoto    PhotoItem    `json:"bestPhoto"`
	Reasons      Reasons      `json:"reasons"`
	Tips         Tips         `json:"tips"`
	Popular      Popular      `json:"popular"`
}

type Contact struct {
	Phone string `json:"phone"`
}

type Location struct {
	Address     string  `json:"address"`
	CrossStreet string  `json:"crossStreet"`
	City        string  `json:"city"`
	State       string  `json:"state"`
	PostalCode  string  `json:"postalCode"`
	Country     string  `json:"country"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
	Distance    float64 `json:"distance"`
}

type Category struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	PluralName string `json:"pluralName"`
	ShortName  string `json:"shortName"`
	Icon       Icon   `json:"icon"`
	Primary    bool   `json:"primary"`
}

type Icon struct {
	Prefix string `json:"prefix"`
	Suffix string `json:"suffix"`
}

type Stats struct {
	CheckinsCount int `json:"checkinsCount"`
	UsersCount    int `json:"usersCount"`
	TipCount      int `json:"tipCount"`
	VisitsCount   int `json:"visitsCount"`
}

type HereNow struct {
	Count int `json:"count"`
}

type Likes struct {
	Count int `json:"count"`
}

type FriendVisits struct {
	Count int `json:"count"`
}

type BeenHere struct {
	Count int `json:"count"`
}

type Specials struct {
	Count int `json:"count"`
}

type Photos struct {
	Count  int            `json:"count"`
	Groups []PhotosGroups `json:"groups"`
}

type PhotosGroups struct {
	Items []PhotoItem `json:"items"`
}

type PhotoItem struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Source     PhotoSource `json:"source"`
	Prefix     string      `json:"prefix"`
	Suffix     string      `json:"suffix"`
	Width      int         `json:"width"`
	Height     int         `json:"height"`
	Visibility string      `json:"visibility"`
}

type PhotoSource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Reasons struct {
	Count int `json:"count"`
}

type Tips struct {
	Count int `json:"count"`
}

type Popular struct {
	Status string `json:"status"`
	IsOpen bool   `json:"isOpen"`
}
