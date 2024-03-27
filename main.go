package smartbidder

import "github.com/SurajKadam7/smart_bidder/internal/bid"

type SmartBidder interface {
	Bid() bid.Bid
}
