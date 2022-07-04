package memo3

import "sync"



type Memo struct {
	mu sync.Mutex
	f Func
	cache map[string]result
}

type Func func (name string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f:f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()

	if !ok {
		res.value, res.err = memo.f(key)

		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}

	return res.value, res.err
}