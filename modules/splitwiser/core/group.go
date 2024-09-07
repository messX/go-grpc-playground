package splitwiser

import "log"

type Group struct {
	members   []Member
	expenses  []Expense
	breakdown []Breakdown
}

func (g *Group) validateMemberAreadyExists(m Member) bool {
	for _, mem := range g.members {
		if m.id == mem.id {
			return true
		}
	}
	return false
}

func (g *Group) addMember(m Member) {
	if g.validateMemberAreadyExists(m) {
		log.Fatal("Member already exist")
	} else {
		g.members = append(g.members, m)
	}
}
