package oracles

import (
	"encoding/json"
	"time"
)

type PriceInfo struct {
	Price, Volume float64
}

type PriceStamp struct {
	Created time.Time
	PriceInfo
}

type window struct {
	buffer []*PriceStamp
	start  int
	num    int
	keep   time.Duration
}

func NewWindow(keep time.Duration) *window {
	return &window{
		keep: keep,
	}
}

func (w *window) Clear() {
	w.start = 0
	w.num = 0
}

func (w *window) Trim() {
	now := time.Now()
	n := len(w.buffer)
	for w.num > 0 {
		v := w.buffer[w.start]
		if now.Sub(v.Created) < w.keep {
			w.start = (w.start + 1) % n
			w.num--
		} else {
			break
		}
	}
}

func (w *window) CloseTwo(at time.Time) (before, after *PriceStamp) {
	if w.num == 0 {
		return
	}
	i := 0
	n := len(w.buffer)
	for i < w.num {
		c := w.buffer[(w.start+i)%n]
		if c.Created.After(at) {
			after = c
			break
		}
		before = c
		i++
	}
	return
}

func (w *window) WithinRange(at time.Time, delta time.Duration) []*PriceStamp {
	var items []*PriceStamp
	i := 0
	for i < w.num {
		c := w.buffer[(w.start+i)%len(w.buffer)]
		d := at.Sub(c.Created)
		if d < 0 {
			break
		}
		if d <= delta {
			items = append(items, c)
		}
		i++
	}
	return items
}

func (w *window) Insert(x *PriceStamp) {
	now := time.Now()
	t := x.Created
	latest := w.Latest()
	//ignore if too old already or if older than current newest
	if now.Sub(t) > w.keep || (latest != nil && t.Sub(latest.Created) < 0) {
		return
	}
	w.Trim()
	n := len(w.buffer)
	if w.num == n {
		newLen := 2 * n
		if newLen == 0 {
			newLen = 1
		}
		bigger := make([]*PriceStamp, newLen)
		for i := 0; i < w.num; i++ {
			bigger[i] = w.buffer[(w.start+i)%n]
		}
		w.start = 0
		w.buffer = bigger
		n = newLen
	}
	w.buffer[(w.start+w.num)%n] = x
	w.num++
}

func (w *window) Len() int {
	w.Trim()
	return w.num
}

func (w *window) Latest() *PriceStamp {
	w.Trim()
	if w.num == 0 {
		return nil
	}

	n := len(w.buffer)
	return w.buffer[(w.start+w.num-1)%n]
}

func (w *window) UnmarshalJSON(b []byte) error {
	var v []*PriceStamp
	err := json.Unmarshal(b, &v)
	if err != nil {
		return err
	}
	w.Clear()
	if w.keep == 0 {
		w.keep = 7 * 24 * time.Hour
	}
	for i := range v {
		w.Insert(v[i])
	}
	return nil
}

func (w *window) MarshalJSON() ([]byte, error) {
	w.Trim()
	a := make([]*PriceStamp, w.num)
	n := len(w.buffer)
	for i := 0; i < w.num; i++ {
		a[i] = w.buffer[(w.start+i)%n]
	}
	return json.Marshal(a)
}
