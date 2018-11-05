package main

import (
	"log"
	"time"

	"github.com/pkg/errors"
)

type autoexec struct {
	done   chan struct{}
	ticker *time.Ticker
	dr     *doer
}

// New returns a AutoExec instance.
func newAutoExec(periodSec int, dr *doer) (*autoexec, error) {
	a := &autoexec{
		ticker: time.NewTicker(time.Duration(periodSec) * time.Second),
		done:   make(chan struct{}),
		dr:     dr,
	}
	go a.start()

	return a, nil
}

func (a *autoexec) start() {
	for {
		select {
		case <-a.done:
			log.Println("autoexec exit")
			return
		case <-a.ticker.C:
			if err := a.exec(); err != nil {
				log.Printf("autoexec[start] exec failed:%v", err)
			}
		}
	}
}

func (a *autoexec) exec() error {
	return errors.Wrapf(a.dr.do(), "autoexec[exec] doer do failed")
}

func (a *autoexec) close() error {
	close(a.done)
	a.ticker.Stop()
	return nil
}
