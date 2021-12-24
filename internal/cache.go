package internal

import (
	"errors"
)

var Members = NewMemberList()

type MemberList struct {
	members map[string]Member
}

func (m *MemberList) IsValid(memberDto MemberDto) error {
	_, ok := Members.members[memberDto.Email]
	if !ok {
		return nil
	}
	return errors.New("email already exists")
}

func (m *MemberList) Add(member Member) {
	m.members[member.Email] = member
}

func NewMemberList() *MemberList {
	members := make(map[string]Member)
	return &MemberList{members: members}
}
