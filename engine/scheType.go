package engine

type Scheduler interface {
	Submit(Request)
	ConfigureChan(chan Request)
}

type SimpleScheduler struct {
	scheduler chan Request
}

func (s *SimpleScheduler) Submit(r Request) {
	go func() {
		s.scheduler <- r
	}()
}

func (s *SimpleScheduler) ConfigureChan(c chan Request) {
	s.scheduler = c
}
