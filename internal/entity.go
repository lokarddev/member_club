package internal

import "time"

type Member struct {
	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
}

type MemberDto struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func NewMember(dto MemberDto) *Member {
	return &Member{
		Name:             dto.Name,
		Email:            dto.Email,
		RegistrationDate: time.Now().Format("202-01-2006"),
	}
}
