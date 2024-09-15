package main

import (
	"log"

	core "github.com/messx/go-grpc-playground/modules/splitwiser/core"
)

func main() {
	log.Print("creating member")
	m1 := core.Member{
		Name: "Test1",
		Id:   "m1",
	}
	m2 := core.Member{
		Name: "Test2",
		Id:   "m2",
	}
	m3 := core.Member{
		Name: "Test3",
		Id:   "m3",
	}
	m4 := core.Member{
		Name: "Test4",
		Id:   "m4",
	}
	grp := core.NewGroup()
	grp.AddMember(m1)
	grp.AddMember(m2)
	grp.AddMember(m3)
	grp.AddMember(m4)
	sp1 := []core.Member{m2, m3, m4}
	sp2 := []core.Member{m1, m3}
	sp3 := []core.Member{m2, m3, m1}
	exp1 := core.Expense{
		AddedBy:      m1,
		Amount:       300.0,
		SplitMembers: sp1,
	}
	exp2 := core.Expense{
		AddedBy:      m2,
		Amount:       1000.0,
		SplitMembers: sp2,
	}
	exp3 := core.Expense{
		AddedBy:      m4,
		Amount:       1300.0,
		SplitMembers: sp3,
	}
	grp.AddExpense(exp1)
	grp.AddExpense(exp2)
	grp.AddExpense(exp3)
	expManager := core.ExpenseManager{EGroup: *grp}
	breakdown, _ := expManager.GetBreakDown()
	log.Print("Expenses for group", breakdown)

}
