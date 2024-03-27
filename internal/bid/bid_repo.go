package bid

import (
	"context"
	"sync"
)

type Repo interface {
	GetRunning(ctx context.Context) []Bid
	Get(ctx context.Context, id int64) Bid
	Add(ctx context.Context, b Bid) Bid
	UpdateHigh(ctx context.Context, id int64, h int) Bid
	UpdateLow(ctx context.Context, id int64, l int) Bid
}

type repo struct {
	idCnt int64
	lock  sync.RWMutex
	bids  map[int64]Bid
}

func (r *repo) GetRunning(ctx context.Context) (bids []Bid) {
	bids = []Bid{}
	r.lock.RLock()
	for _, b := range r.bids {
		if b.Status != Running {
			continue
		}
		bids = append(bids, b)
	}
	r.lock.RUnlock()
	return bids
}

func (r *repo) Get(id int64) (b Bid) {
	r.lock.RLock()
	b = r.bids[id]
	r.lock.RUnlock()
	return b
}

func (r *repo) Add(ctx context.Context, b Bid) (id int64) {
	r.lock.Lock()
	r.idCnt++
	b.Id = r.idCnt
	r.bids[r.idCnt] = b
	id = r.idCnt
	r.lock.Unlock()
	return
}

func (r *repo) UpdateHigh(ctx context.Context, id int64, h int) {
	r.lock.RLock()
	b, ok := r.bids[r.idCnt]
	if !ok {
		return
	}

	b.High = h
	r.bids[r.idCnt] = b
	r.lock.RUnlock()
}

func (r *repo) UpdateLow(ctx context.Context, id int64, l int) {
	r.lock.RLock()
	b, ok := r.bids[r.idCnt]
	if !ok {
		return
	}

	b.Low = l
	r.bids[r.idCnt] = b
	r.lock.RUnlock()
}
