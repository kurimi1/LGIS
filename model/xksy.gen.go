// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameXksy = "xksy"

// Xksy mapped from table <xksy>
type Xksy struct {
	KCAAA  string  `gorm:"column:KCAAA" json:"KCAAA"`
	XYAD   string  `gorm:"column:XYAD" json:"XYAD"`
	XYAABA string  `gorm:"column:XYAABA" json:"XYAABA"`
	XYABA  string  `gorm:"column:XYABA" json:"XYABA"`
	XYACF  float32 `gorm:"column:XYACF" json:"XYACF"`
	XYACB  string  `gorm:"column:XYACB" json:"XYACB"`
	XYACA  string  `gorm:"column:XYACA" json:"XYACA"`
	PKGLC  string  `gorm:"column:PKGLC" json:"PKGLC"`
	XYACJ  float32 `gorm:"column:XYACJ" json:"XYACJ"`
	XYACD  string  `gorm:"column:XYACD" json:"XYACD"`
	PKGKPE string  `gorm:"column:PKGKPE" json:"PKGKPE"`
	PKGKQ  string  `gorm:"column:PKGKQ" json:"PKGKQ"`
	PKHFE  string  `gorm:"column:PKHFE" json:"PKHFE"`
	数据     string  `gorm:"column:数据" json:"数据"`
}

// TableName Xksy's table name
func (*Xksy) TableName() string {
	return TableNameXksy
}
