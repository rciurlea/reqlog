package main

import (
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type reqentry struct {
	When    time.Time
	URL     *url.URL
	Headers http.Header
	Method  string
	Body    []byte
}

type reqlist struct {
	requests []*reqentry
	m        sync.Mutex
}

func newReqlist() *reqlist {
	return &reqlist{
		requests: make([]*reqentry, 0),
	}
}

func (rl *reqlist) add(req *reqentry) {
	rl.m.Lock()
	defer rl.m.Unlock()
	rl.requests = append(rl.requests, req)
	spew.Dump(rl.requests)
}

func (rl *reqlist) getAll() []*reqentry {
	return nil
}
