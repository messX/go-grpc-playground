package splitwiser

type Expense struct {
	addedBy Member
	// we can make split type and split the expenses as per % and absolute values
	amount       float32
	splitMembers []Member
}
