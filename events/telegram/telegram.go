package telegram

import (
	"read-adviser-bot/clients/telegram"
	"read-adviser-bot/storage"
)

//meta for Telegram

type Meta struct {
	ChatID   int
	Username string
}

//Collector
type Processor struct {
	tg     *telegram.Client
	offset int
	//storage
	storage storage.Storage
}

func New(client, storage) *Processor {

	return &Processor{
		tg:      client,
		storage: storage,
	}

}

func (p *Processor) Fetch(limit) ([]events.Event, error) {
	update, err := p.tg.Updates(p.offset, limit)
	if err != nil {
		return nil, e.Wrap("can't get events", err)
	}

	if len(update) == 0 {
		return nil, nil
	}

	//allocate memory

	res := make([]events.Event, 0, len(update))

	for _, u := range update {
		res = append(res, Event(u))
	}

	//important:

	p.offset = updates[len(updates)-1].ID + 1

	return res, nil

}

func Event(update) events.Event {
	updateType := fetchType(update)

	res := events.Event{
		Type: updateType,
		Text: fetchText(update),
	}

	//chatID username - for telegram only

	if updateType == events.Message {
		res.Meta = Meta{
			ChatID:   update.Message.Chat.ID,
			Username: update.Message.From.Username,
		}
	}

	return res
}

func fetchText(update) string {
	if update.Message == nil {
		return ""
	}
	return update.Message.Text
}

func fetchType(update) events.Type {
	if update.Message == nil {
		return events.Unknown
	}

	return events.Message
}
