package domain

import (
	"strconv"
	"time"
)

type User struct {
	Model
	FamilyName         string       `form:"familyName" binding:"required" json:"familyName"`
	GivenName          string       `form:"givenName" binding:"required" json:"givenName"`
	FamilyNameKana     string       `form:"familyNameKana" binding:"required"`
	GivenNameKana      string       `form:"givenNameKana" binding:"required"`
	PostalCode         string       `form:"postalCode" binding:"required"`
	PrefectureId       Prefecture   `form:"prefectureId" binding:"required"`
	Address1           string       `form:"address1" binding:"required"`
	Address2           string       `form:"address2" binding:"required"`
	Address3           *string      `form:"address3" binding:"required"`
	PhoneNumber        string       `form:"phoneNumber" binding:"required"`
	HomePhoneNumber    *string      `form:"homePhoneNumber"`
	Email              string       `form:"email" binding:"required"`
	Gender             Gender       `form:"gender" binding:"required"`
	Birthday           time.Time    `form:"birthday" binding:"required" time_format:"2006-01-02"`
	Occupation         Occupation   `form:"occupation"`
	FirstVisitDate     *time.Time   `form:"firstVistDate" time_format:"2006-01-02 15:04"`
	DmForwardingFlg    *string      `form:"dmForwardingFlg"`
	Memo               *string      `form:"memo"`
	ReasonForComing    *string      `form:"reasonForComing"`
	FamilyUserId       *uint        `form:"familyUserId"`
	FamilyRelationship Relationship `form:"familyRelationship"`
	Orders             []Order
	NextVisitDate      *time.Time `json:"nextVisitDate"`
	PreviousVisitDate  *time.Time `json:"previousVisitDate"`
	LastVistDates      *int       `json:"lastVistDates"`
}

type Gender int
type Occupation int
type Prefecture int
type Relationship int

const (
	Mela Gender = iota + 1
	Femela
	NoAnswer
)

func (c Gender) String() string {
	switch c {
	case Mela:
		return "男性"
	case Femela:
		return "女性"
	case NoAnswer:
		return "選択なし"
	default:
		return "Unknown"
	}
}

const (
	ITEngineer Occupation = iota + 1
)

func (occ Occupation) String() string {
	switch occ {
	case ITEngineer:
		return "ITエンジニア"
	default:
		return "Unkown"
	}
}

const (
	hokkaido Prefecture = iota + 1
	aomori
	iwate
	miyagi
	akita
	yamagata
	fukushima
	ibaraki
	tochigi
	gunma
	saitama
	chiba
	tokyo
	kanagawa
	niigata
	toyama
	ishikawa
	fukui
	yamanashi
	nagano
	gifu
	shizuoka
	aichi
	mie
	shiga
	kyoto
	osaka
	hyogo
	nara
	wakayama
	tottori
	shimane
	okayama
	hiroshima
	yamaguchi
	tokushima
	kagawa
	ehime
	kochi
	fukuoka
	saga
	nagasaki
	kumamoto
	oita
	miyazaki
	kagoshima
	okinawa
)

func (pref Prefecture) String() string {
	switch pref {
	case hokkaido:
		return "北海道"
	case aomori:
		return "青森県"
	case iwate:
		return "岩手県"
	case miyagi:
		return "宮城県"
	case akita:
		return "秋田県"
	case yamagata:
		return "山形県"
	case fukushima:
		return "福島県"
	case ibaraki:
		return "茨城県"
	case tochigi:
		return "栃木県"
	case gunma:
		return "群馬県"
	case saitama:
		return "埼玉県"
	case chiba:
		return "千葉県"
	case tokyo:
		return "東京都"
	case kanagawa:
		return "神奈川県"
	case niigata:
		return "新潟県"
	case toyama:
		return "富山県"
	case ishikawa:
		return "石川県"
	case fukui:
		return "福井県"
	case yamanashi:
		return "山梨県"
	case nagano:
		return "長野県"
	case gifu:
		return "岐阜県"
	case shizuoka:
		return "静岡県"
	case aichi:
		return "愛知県"
	case mie:
		return "三重県"
	case shiga:
		return "滋賀県"
	case kyoto:
		return "京都府"
	case osaka:
		return "大阪府"
	case hyogo:
		return "兵庫県"
	case nara:
		return "奈良県"
	case wakayama:
		return "和歌山県"
	case tottori:
		return "鳥取県"
	case shimane:
		return "島根県"
	case okayama:
		return "岡山県"
	case hiroshima:
		return "広島県"
	case yamaguchi:
		return "山口県"
	case tokushima:
		return "徳島県"
	case kagawa:
		return "香川県"
	case ehime:
		return "愛媛県"
	case kochi:
		return "高知県"
	case fukuoka:
		return "福岡県"
	case saga:
		return "佐賀県"
	case nagasaki:
		return "長崎県"
	case kumamoto:
		return "熊本県"
	case oita:
		return "大分県"
	case miyazaki:
		return "宮崎県"
	case kagoshima:
		return "鹿児島県"
	case okinawa:
		return "沖縄県"
	default:
		return "Unknow"
	}
}

const (
	Mother Relationship = iota + 1
	Father
	Brother
	Sister
	Children
)

func (rela Relationship) String() string {
	switch rela {
	case Mother:
		return "母"
	case Father:
		return "父"
	case Brother:
		return "兄弟"
	case Sister:
		return "姉妹"
	case Children:
		return "子ども"
	default:
		return "Unkwon"
	}
}

func Genders() map[int]map[string]string {
	gs := make(map[int]map[string]string)

	for i, v := range []Gender{Mela, Femela, NoAnswer} {
		gs[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}

	return gs
}

func Occupations() map[int]map[string]string {
	occs := make(map[int]map[string]string)

	for i, v := range []Occupation{ITEngineer} {
		occs[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}
	return occs
}

func Prefectures() map[int]map[string]string {
	prefs := make(map[int]map[string]string)

	for i, v := range []Prefecture{hokkaido, aomori, iwate, miyagi, akita, yamagata, fukushima, ibaraki, tochigi, gunma, saitama, chiba, tokyo, kanagawa, niigata, toyama, ishikawa, fukui, yamanashi, nagano, gifu, shizuoka, aichi, mie, shiga, kyoto, osaka, hyogo, nara, wakayama, tottori, shimane, okayama, hiroshima, yamaguchi, tokushima, kagawa, ehime, kochi, fukuoka, saga, nagasaki, kumamoto, oita, miyazaki, kagoshima, okinawa} {
		prefs[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}

	return prefs
}

func Relationships() map[int]map[string]string {
	rela := make(map[int]map[string]string)

	for i, v := range []Relationship{Mother, Father, Brother, Sister, Children} {
		rela[i] = map[string]string{"ja": v.String(), "value": strconv.Itoa(int(v))}
	}

	return rela
}
