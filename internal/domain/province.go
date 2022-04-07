package domain

type Province struct {
	ID           int    `json:"id"`
	ProvinceName string `json:"province"`
	CountryId    string `json:"country"`
}
