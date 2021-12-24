package internal

import "time"

var Members = NewMemberList()

type MemberList struct {
	members []Member
}

func (m *MemberList) Add(member Member) {
	m.members = append(m.members, member)
}

func NewMemberList() *MemberList {
	return &MemberList{members: []Member{{
		Name:             "test",
		Email:            "test",
		RegistrationDate: time.Now().Format("2-01-2006"),
	}}}
}
