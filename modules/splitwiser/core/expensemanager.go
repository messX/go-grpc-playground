package splitwiser

import (
	"log"
	"math"
)

type ExpenseManager struct {
	EGroup Group
}

// Method to create a breakdown of balances for each member
func (expm *ExpenseManager) GetBreakDown() ([]Breakdown, error) {
	log.Printf("Getting group expenses")
	expenseMap := make(map[string]float32)
	for _, exp := range expm.EGroup.expenses {
		expOwner := exp.AddedBy.Id
		expAmount := exp.Amount
		_, exists := expenseMap[expOwner]
		if exists {
			expenseMap[expOwner] += expAmount
		} else {
			expenseMap[expOwner] = expAmount
		}
		splitAmount := expAmount / float32((len(exp.SplitMembers)))
		for _, expMem := range exp.SplitMembers {
			_, existMem := expenseMap[expMem.Id]
			if existMem {
				expenseMap[expMem.Id] -= splitAmount
			} else {
				expenseMap[expMem.Id] = -1 * splitAmount
			}
		}
	}
	expm.EGroup.breakdown = make([]Breakdown, 0)
	for memId, amt := range expenseMap {
		var mem Member
		for _, grpMem := range expm.EGroup.members {
			if grpMem.Id == memId {
				mem = grpMem
			}
		}
		memBreakdown := Breakdown{mem, amt}
		expm.EGroup.breakdown = append(expm.EGroup.breakdown, memBreakdown)

	}
	return expm.EGroup.breakdown, nil
}

// Recursive Method to create split balances for between the members
func getRecSplitBalance(bdList []Breakdown, prevSplit []Split) ([]Split, error) {
	allZero := true

	bdListUpdate := func(spltInner Split, toAmt float32, byAmt float32) {
		for i, bdInner := range bdList {
			if bdInner.member.Id == spltInner.By.Id {
				bdList[i].amount = byAmt
			}
			if bdInner.member.Id == spltInner.To.Id {
				bdList[i].amount = toAmt
			}
		}
	}

	for _, bd := range bdList {
		if float32(math.Round(math.Abs(float64(bd.amount)))) != float32(0) {
			allZero = false
		}
	}

	if allZero {
		return prevSplit, nil
	} else {
		least, most := findLeastAndMost(bdList)
		if most.amount > -1*least.amount {
			splt := Split{least.member, most.member, -1 * least.amount}
			prevSplit = append(prevSplit, splt)
			bdListUpdate(splt, most.amount+least.amount, 0)
		} else {
			splt := Split{least.member, most.member, most.amount}
			prevSplit = append(prevSplit, splt)
			bdListUpdate(splt, 0, least.amount+most.amount)
		}
	}
	return getRecSplitBalance(bdList, prevSplit)
}

func (expm *ExpenseManager) SplitBalance() ([]Split, error) {
	breakdownList := deepCopySlice(expm.EGroup.breakdown)
	splitList := []Split{}
	return getRecSplitBalance(breakdownList, splitList)
}

func deepCopySlice(original []Breakdown) []Breakdown {
	// Create a new slice of the same length and capacity
	copy := make([]Breakdown, len(original))

	// Copy each element to the new slice
	for i, val := range original {
		copy[i] = val
	}
	return copy
}

// Function to find the minimum value in a slice of integers
func findMin(nums []Breakdown) Breakdown {
	// Initialize min with the maximum possible integer value
	min := nums[0]

	// Loop through the list and find the minimum value
	for _, num := range nums[1:] {
		if num.amount < min.amount {
			min = num
		}
	}

	return min
}

// Function to find the minimum value in a slice of integers
func findMax(nums []Breakdown) Breakdown {
	// Initialize min with the maximum possible integer value
	max := nums[0]

	// Loop through the list and find the minimum value
	for _, num := range nums[1:] {
		if num.amount > max.amount {
			max = num
		}
	}

	return max
}

func findLeastAndMost(breakdown []Breakdown) (Breakdown, Breakdown) {
	min := findMin(breakdown)
	max := findMax(breakdown)
	return min, max
}
