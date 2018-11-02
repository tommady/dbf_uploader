package main

import (
	"log"
	"time"
)

type autoexec struct {
	done   chan struct{}
	ticker *time.Ticker
}

// New returns a AutoExec instance.
func newAutoExec(periodSec int) (*autoexec, error) {
	a := &autoexec{
		ticker: time.NewTicker(time.Duration(periodSec) * time.Second),
		done:   make(chan struct{}),
	}
	go a.start()

	return a, nil
}

func (a *autoexec) start() {
	for {
		select {
		case <-a.done:
			log.Println("autoexec: exit")
			return
		case <-a.ticker.C:
			if err := a.exec(); err != nil {
				log.Println("autoexec: [start] exec failed:%v", err)
			}
		}
	}
}

func (a *autoexec) exec() error {
	// keys, err := a.service.GetExternalKeys()
	// if err != nil {
	// 	return errors.Wrapf(err, "autoexec: [exec] GetExternalKeys failed")
	// }

	// keyMap := make(map[string]struct{})
	// for _, key := range keys {
	// 	if _, exist := keyMap[key]; !exist {
	// 		err = a.service.PushJob(&model.Job{
	// 			Type:  model.ExternalKeyJob,
	// 			Value: key,
	// 		})
	// 		if err != nil {
	// 			return errors.Wrapf(err, "autoexec: [exec] push external key:%s job failed", key)
	// 		}

	// 		keyMap[key] = struct{}{}
	// 	}
	// }

	return nil
}

func (a *autoexec) close() error {
	close(a.done)
	a.ticker.Stop()
	return nil
}
