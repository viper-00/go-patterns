package objectPool

import (
	"errors"
	"log"
	"math/rand"
	"sync"
	"time"
)

// wiki: https://en.wikipedia.org/wiki/Object_pool_pattern
//
// Instantiates and maintains a group of objects instances of the same type.

const getResMaxTime = 3 * time.Second

var (
	ErrPoolNotExist  = errors.New("pool not exist")
	ErrGetResTimeout = errors.New("get resource time out")
)

type Resource struct {
	resId int
}

func NewResource(id int) *Resource {
	time.Sleep(500 * time.Millisecond)
	return &Resource{resId: id}
}

func (r *Resource) Do(workId int) {
	time.Sleep(time.Duration(rand.Intn(5) * 100 * int(time.Millisecond)))
	log.Printf("using resource #%d finished work %d finish\n", r.resId, workId)
}

type Pool chan *Resource

func New(size int) Pool {
	p := make(Pool, size)
	wg := new(sync.WaitGroup)
	wg.Add(size)
	for i := 0; i < size; i++ {
		go func(resId int) {
			p <- NewResource(resId)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return p
}

func (p Pool) GetResource() (r *Resource, err error) {
	select {
	case r := <-p:
		return r, nil
	case <-time.After(getResMaxTime):
		return nil, ErrGetResTimeout
	}
}

func (p Pool) GiveBackResource(r *Resource) error {
	if p == nil {
		return ErrPoolNotExist
	}
	p <- r
	return nil
}

func main() {
	size := 5
	p := New(size)

	doWork := func(workId int, wg *sync.WaitGroup) {
		defer wg.Done()
		res, err := p.GetResource()
		if err != nil {
			log.Println(err)
			return
		}
		defer p.GiveBackResource(res)
		res.Do(workId)
	}

	num := 100
	wg := new(sync.WaitGroup)
	wg.Add(num)
	for i := 0; i < num; i++ {
		go doWork(i, wg)
	}
	wg.Wait()
}
