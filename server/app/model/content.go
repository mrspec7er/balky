package model

type Content struct {
	ApplicationNumber string `gorm:"primaryKey"`
	Application *Application `json:"application" gorm:"foreignKey:ApplicationNumber"`

	AttributeID uint `gorm:"primaryKey;autoIncrement:false"`
	Attribute *Attribute	`json:"attribute"`

	Value string `json:"value" gorm:"type:text"`
}