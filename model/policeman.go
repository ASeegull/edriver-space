package model

type Policeman struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	BadgeNum   string `json:"badge_num"`
	ExpireDate string `json:"expire_date"`
}

func MakePoliceman(name, surname, badgeNum, expireDate string) Policeman {
	return Policeman{Name: name, Surname: surname, BadgeNum: badgeNum, ExpireDate: expireDate}
}
