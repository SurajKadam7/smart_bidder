package bid

import "context"

type Status int

const (
	Running Status = iota + 1
	Completed
	Canceled
)

// if bid has the control then it can check whether the given user can do perticular thing or not
// if caller has access then it is completely decoupled

type Common interface {
	GetRunningBids(ctx context.Context) []Bid
	GetWinner(ctx context.Context, id int64)
}

type BuyerService interface {
	Withdraw(ctx context.Context, buyerId int64, id int64)
	Bid(ctx context.Context, userId int64, id int64) Bid
	Common
}

type SellerService interface {
	AddToAuction(ctx context.Context, o Opt) Bid
	WithDrawFromAuction(ctx context.Context, sellerId int64, id int64)
	SellToBidder(ctx context.Context, sellerId int64, id int64) // this will contain the logic of choosing the statergy
	UpdateBid(ctx context.Context, sellerId int64, b Bid)       // not workable
	Common
}

type Opt struct {
	Products []Product
	High     int
	Low      int
	Seller   Seller
	Statergy Statergy // if not pass setting defualt
}

func (o Opt) isValid() bool {
	if len(o.Products) == 0 {
		return false
	}
	if o.Low >= o.High {
		return false
	}

	if o.Seller.Id == 0 {
		return false
	}
	return true
}

type service struct {
	BuyerService
	SellerService
}

func New(o Opt) (b BuyerService, s SellerService) {
	if !o.isValid() {
		return
	}
	return
}
