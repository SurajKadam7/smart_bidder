package bid

type Buyer struct {
	Id     int
	Name   string
	Amount int
}

type Seller struct {
	Id   int
	Name string
	Org  string
}

type Product struct {
	Id   int
	Name string
	Qty  int
}

type Bid struct {
	Id       int64
	Products []Product
	High     int
	Low      int
	Status   Status
	Seller   Seller
	Buyers   map[int]Buyer
	Winner   int
	Statergy Statergy
}
