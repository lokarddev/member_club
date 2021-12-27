package internal

var Members = NewMemberList()

type MemberList struct {
	members map[string]Member
}

func (m *MemberList) Add(member Member) {
	m.members[member.Email] = member
}

func NewMemberList() *MemberList {
	members := make(map[string]Member)
	return &MemberList{members: members}
}
