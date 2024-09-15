package splitwiser

import "fmt"

type Split struct {
	By     Member
	To     Member
	Amount float32
}

func (splt *Split) PrintBeautifully() string {
	return fmt.Sprintf("%s is to be paid an equivalent of amount: %f by %s", splt.To.Name, splt.Amount, splt.By.Name)
}
