package splitwiser

import (
	"log"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func CreateBasicExpenses() ExpenseManager {
	m1 := Member{
		Name: "Test1",
		Id:   "m1",
	}
	m2 := Member{
		Name: "Test2",
		Id:   "m2",
	}
	m3 := Member{
		Name: "Test3",
		Id:   "m3",
	}
	m4 := Member{
		Name: "Test4",
		Id:   "m4",
	}
	grp := NewGroup()
	grp.AddMember(m1)
	grp.AddMember(m2)
	grp.AddMember(m3)
	grp.AddMember(m4)
	sp1 := []Member{m2, m3, m4}
	sp2 := []Member{m1, m3}
	sp3 := []Member{m2, m3, m1}
	exp1 := Expense{
		AddedBy:      m1,
		Amount:       300.0,
		SplitMembers: sp1,
	}
	exp2 := Expense{
		AddedBy:      m2,
		Amount:       1000.0,
		SplitMembers: sp2,
	}
	exp3 := Expense{
		AddedBy:      m4,
		Amount:       1300.0,
		SplitMembers: sp3,
	}
	grp.AddExpense(exp1)
	grp.AddExpense(exp2)
	grp.AddExpense(exp3)
	expManager := ExpenseManager{EGroup: *grp}
	return expManager
}

// to check if the group balance is correct
func TestGroupBreakup(t *testing.T) {

	log.Println("Checking the group balance")
	expManager := CreateBasicExpenses()
	breakdown, err := expManager.GetBreakDown()
	assert.Nil(t, err)
	var want float64
	var got float32
	want = 0
	got = 0
	for _, bd := range breakdown {
		log.Printf("Mem %s amount is %f", bd.member.Name, bd.amount)
		got += bd.amount
	}
	assert.Equal(t, want, math.Round(math.Abs(float64(got))))

}

func TestGroupSplit(t *testing.T) {

	log.Println("Checking the group split")
	expManager := CreateBasicExpenses()
	breakdown, err := expManager.GetBreakDown()
	assert.Nil(t, err)
	log.Print("breakdown :", breakdown)
	spltList, err := expManager.SplitBalance()

	for _, splt := range spltList {
		log.Println(splt.PrintBeautifully())
	}
	assert.Nil(t, err)

}
