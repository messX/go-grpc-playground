package splitwiser

type Expense struct {
	AddedBy Member
	// we can make split type and split the expenses as per % and absolute values
	Amount       float32
	SplitMembers []Member
}
