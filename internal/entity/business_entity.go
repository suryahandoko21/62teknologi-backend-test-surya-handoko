package entity

type Businesses struct {
	ID                  string `gorm:"column:id;primaryKey"`
	Location            string `gorm:"column:location"`
	Latitude            int64  `gorm:"column:latitude"`
	Longitude           int64  `gorm:"column:longitude"`
	Term                string `gorm:"column:term"`
	Radius              int64  `gorm:"column:radius"`
	Categories          string `gorm:"column:categories"`
	Locale              string `gorm:"column:locale"`
	Price               string `gorm:"column:price"`
	OpenNow             bool   `gorm:"column:open_now"`
	OpenAt              int64  `gorm:"column:open_at"`
	Attributes          string `gorm:"column:attributes"`
	SortBy              string `gorm:"column:sort_by"`
	DevicePlatform      string `gorm:"column:device_platform"`
	ReservationDate     string `gorm:"column:reservation_date"`
	ReservationTime     string `gorm:"column:reservation_time"`
	ReservationCover    int64  `gorm:"column:reservation_covers"`
	MatchPartySizeParam bool   `gorm:"column:matches_party_size_param"`
	Limit               int64  `gorm:"column:limit"`
	Offset              int64  `gorm:"column:offset"`
}

func (c *Businesses) TableName() string {
	return "businesses"
}
