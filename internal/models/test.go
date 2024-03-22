package models

type TestCreateBody struct {
	Name int `json:"name"`
}

type TestFind struct {
	Id         int    `json:"id"`
	CreateDate string `json:"createDate"`
	CreateUser string `json:"createUser"`
	UpdateDate string `json:"updateDate"`
	UpdateUser string `json:"updateUser"`
	CreateApp  string `json:"createApp"`
	UpdateApp  string `json:"updateApp"`
}

type EntityAuthTypeData struct {
	Id     int    `json:"id"`
	NameRu string `json:"nameRu"`
	NameKz string `json:"nameKz"`
}
