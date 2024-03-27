package bid

type StatergyType int

const (
	IHighestUnique StatergyType = iota //defaut
	ILowestUnique
)

func getStatergy(st StatergyType) Statergy {
	switch st {
	case IHighestUnique:
		return HighestUnique
	case ILowestUnique:
		return LowestUnique
	default:
		return nil
	}
}

type Statergy func(buyers []Buyer) (b Buyer, err error)

var HighestUnique = func(buyers []Buyer) (b Buyer, err error) {
	return
}

var LowestUnique = func(buyers []Buyer) (b Buyer, err error) {
	return
}
