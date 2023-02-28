package schema

type Asset struct {
	ID      int64  `gorm:"primary_key;auto_increment" json:"id"`
	EccKey  string `json:"eccKey"`
	Address string `json:"address"`
	Balance string `json:"balance"`
}
