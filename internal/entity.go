package internal

import (
	"errors"
	"net/mail"
	"regexp"
	"time"
)

type Member struct {
	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	RegistrationDate string `json:"registration_date,omitempty"`
}

type MemberDto struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func (m *MemberDto) IsValid() error {
	if !regexp.MustCompile(`^[a-zA-Z .]+$`).MatchString(m.Name) {
		return errors.New("invalid name")
	}
	_, err := mail.ParseAddress(m.Email)
	if err != nil {
		return errors.New("invalid email")
	}
	_, ok := Members.members[m.Email]
	if !ok {
		return nil
	}
	return errors.New("email already exists")
}

func NewMember(dto MemberDto) *Member {
	return &Member{
		Name:             dto.Name,
		Email:            dto.Email,
		RegistrationDate: time.Now().Format("2-01-2006"),
	}
}
