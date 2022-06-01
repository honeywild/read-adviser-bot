package event_consumer

import (
	"log"
	"read-adviser-bot/events"
	"time"
)

type Consumer struct {
	fether    events.Fetcher
	processor events.Processor
	batchSize int
}

func New(fetcher, processor, batchSize) Consumer {
	return Consumer{
		fetcher:   fetcher,
		processor: processor,
		batchSize: batchSize,
	}
}

func (c Consumer) Start() error {
	for {
		//TODO: fetcher RETRY
		gotEvents, err := c.fetcher.Fetch(c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())

			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		if err := c.handleEvents(gotEvents); err != nil {
			log.Print(err)
			continue
		}

	}
}

/* 1. event losses: retry, backup, fallback,
	   store in memory,
	   fetcher's ack.

    2. when handle all batch

	3.  sync.WaitGroup{}
*/

func (c *Consumer) handleEvents(events) error {
	for _, event := range events {
		log.Printf("got new event %s", event.Text)

		if ee := c.processor.Proccess(event); err != nil {
			//TODO: retry, backup
			log.Printf("can't handle event: %s", err.Error())

			continue
		}

	}

}
