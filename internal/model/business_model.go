package model

type BusinesseResponse struct {
	ID                  string `json:"id"`
	Location            string `json:"location"`
	Latitude            int64  `json:"latitude"`
	Longitude           int64  `json:"longitude"`
	Term                string `json:"term"`
	Radius              int64  `json:"radius"`
	Categories          string `json:"categories"`
	Locale              string `json:"locale"`
	Price               string `json:"price"`
	OpenNow             bool   `json:"open_now"`
	OpenAt              int64  `json:"open_at"`
	Attributes          string `json:"attributes"`
	SortBy              string `json:"sort_by"`
	DevicePlatform      string `json:"device_platform"`
	ReservationDate     string `json:"reservation_date"`
	ReservationTime     string `json:"reservation_time"`
	ReservationCover    int64  `json:"reservation_covers"`
	MatchPartySizeParam bool   `json:"matches_party_size_param"`
	Limit               int64  `json:"limit"`
	Offset              int64  `json:"offset"`
}

type BusinesseUpdateRequest struct {
	ID                  string   `json:"id"`
	Location            string   `json:"location"`
	Latitude            int64    `json:"latitude"`
	Longitude           int64    `json:"longitude"`
	Term                string   `json:"term"`
	Radius              int64    `json:"radius"`
	Categories          []string `json:"categories"`
	Locale              string   `json:"locale"`
	Price               []int64  `json:"price"`
	OpenNow             bool     `json:"open_now"`
	OpenAt              int64    `json:"open_at"`
	Attributes          []string `json:"attributes"`
	SortBy              string   `json:"sort_by"`
	DevicePlatform      string   `json:"device_platform"`
	ReservationDate     string   `json:"reservation_date"`
	ReservationTime     string   `json:"reservation_time"`
	ReservationCover    int64    `json:"reservation_covers"`
	MatchPartySizeParam bool     `json:"matches_party_size_param"`
	Limit               int64    `json:"limit"`
	Offset              int64    `json:"offset"`
}

type BusinesseCreateRequest struct {
	Location            string   `json:"location"`
	Latitude            int64    `json:"latitude"`
	Longitude           int64    `json:"longitude"`
	Term                string   `json:"term"`
	Radius              int64    `json:"radius"`
	Categories          []string `json:"categories" gorm:"type:json"`
	Locale              string   `json:"locale"`
	Price               []int64  `json:"price"`
	OpenNow             bool     `json:"open_now"`
	OpenAt              int64    `json:"open_at"`
	Attributes          []string `json:"attributes" gorm:"type:json"`
	SortBy              string   `json:"sort_by"`
	DevicePlatform      string   `json:"device_platform"`
	ReservationDate     string   `json:"reservation_date"`
	ReservationTime     string   `json:"reservation_time"`
	ReservationCover    int64    `json:"reservation_covers"`
	MatchPartySizeParam bool     `json:"matches_party_size_param"`
	Limit               int64    `json:"limit"`
	Offset              int64    `json:"offset"`
}

type SearchBusinesseRequest struct {
	Term      string `json:"term" `
	Location  string `json:"location" `
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Radius    string `json:"radius"`
	Page      int    `json:"page"`
	Size      int    `json:"size"`
}

type DeleteBusinesseRequest struct {
	ID string `json:"-" validate:"required,max=100,uuid"`
}
