package memo5



type Memo struct {
	requests chan request
}

type request struct {
	key string
	response chan<- result
}

type result struct {
	value interface{}
	err error
}

type Func func(name string) (interface{}, error)

type entry struct {
	ready chan struct{}
	res result
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.serve(f)
	return memo
}

func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <- response
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) serve(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)

	close(e.ready)
}

func (e *entry) deliver(res chan<- result){
	<- e.ready

	res <- e.res
}