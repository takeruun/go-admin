package domain

import (
	"fmt"
	"strings"
	"time"
)

type Model struct {
	ID        uint       `gorm:"primary_key" json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `sql:"index"`
}

func (m *Model) ParseTommdd(t time.Time) string {
	return t.Format("01月02日")
}

func (m *Model) ParseTommddHM(t time.Time) string {
	return t.Format("01月02日 15時04分")
}

func (m *Model) ConvertPrice(p int) string {
	return "¥" + convert(p)
}

func convert(integer int) string {
	arr := strings.Split(fmt.Sprintf("%d", integer), "")
	cnt := len(arr) - 1
	res := ""
	i2 := 0
	for i := cnt; i >= 0; i-- {
		if i2 > 2 && i2%3 == 0 {
			res = fmt.Sprintf(",%s", res)
		}
		res = fmt.Sprintf("%s%s", arr[i], res)
		i2++
	}
	return res
}
