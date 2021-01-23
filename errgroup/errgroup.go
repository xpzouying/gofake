package errgroup

import "sync"

type group struct {
	err     error
	errOnce sync.Once

	wg sync.WaitGroup
}

func New() *group {
	return new(group)
}

func (g *group) Go(fn func() error) {
	g.wg.Add(1)

	go func() {
		defer g.wg.Done()

		if err := fn(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
			})
		}
	}()
}

func (g *group) Wait() error {
	g.wg.Wait()

	return g.err
}
