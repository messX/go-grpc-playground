package splitwiser

import "log"

type Group struct {
	members   []Member
	expenses  []Expense
	breakdown []Breakdown
}

func NewGroup() *Group {
	return &Group{
		members:   make([]Member, 0),
		expenses:  make([]Expense, 0),
		breakdown: make([]Breakdown, 0),
	}
}

func (g *Group) validateMemberAreadyExists(m Member) bool {
	for _, mem := range g.members {
		if m.Id == mem.Id {
			return true
		}
	}
	return false
}

func (g *Group) AddMember(m Member) {
	if g.validateMemberAreadyExists(m) {
		log.Fatal("Member already exist")
	} else {
		g.members = append(g.members, m)
	}
}

func (g *Group) AddExpense(exp Expense) {
	g.expenses = append(g.expenses, exp)
}
