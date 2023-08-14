package service

type tokenResponse struct {
	AccessToken string `json:"access_token"`
}

type getSettingSectionResponse struct {
	SectionName string `json:"sectionName"`
}
