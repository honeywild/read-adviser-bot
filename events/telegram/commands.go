package telegram

import (
	"log"
	"net/url"
	"strings"
)

const (
	RndCmd   = "/rnd"
	HelpCmd  = "/help"
	StartCmd = "/start"
)

//methods:

func (p *Proccessor) doCmd(text, chatID, username) error {
	text = strings.Trimspace(text)
	log.Printf("got new command '%s' from '%s'", text, username)

	// add page: http://
	// rnd page: /rnd
	// help: /help
	// start: /start: hi + help

	//router
	if isAddCmd(text) {
		return p.savePage(chatID, text, username)
	}

	switch text {
	case RndCmd:
		return p.sendRandom(chatID, username)
	case HelpCmd:
		return p.sendHelp(chatID)
	case StartCmd:
		return p.sendHello(chatID)
	default:
		return p.tg.SendMessage(chatID, msgUnknownCommand)
	}
}

func isAddCmd(text) bool {
	return isUrl(text)
}

func isUrl(text) bool {
	//TODO: links like 'ya.ru'
	u, err := url.Parse(text)
	return err == nil && u.Host != ""

}

//TODO: implement closure

func (p *Processor) savePage(chatID, text, username) err {
	defer func() { err = e.WrapIfErr("can't do command: save page", err) }()

	page := &Storage.Page{
		URL:      pageURL,
		UserName: username,
	}

	isExists, err := p.storage.IsExists(page)
	if err != nil {
		return err
	}

	if isExists {
		return p.tg.SendMessage(chatID, msgAlreadyExists)
	}

	if err := p.storage.Save(page); err != nil {
		return err
	}

	if err := p.tg.SendMessage(chatID, msgSaved); err != nil {
		return err
	}

	return nil
}

func (p *Processor) sendRandom(chatID, username) (err error) {
	defer func() { err = e.WrapIfErr("can't do command: can't send random", err) }()

	page, err := p.storage.PickRandom(username)
	if err != nil && !errors.Is(err, storage.ErrNoSavedPages) {
		return err
	}
	if errors.Is(err, storage.ErrNoSavedPages) {
		return p.tg.SendMessage(chatID, msgNoSavedPages)
	}

	if err := p.tg.SendMessage(chatID, page.URL); err != nil {
		return err
	}

	//delete a like when found

	return p.storage.Remove(page)

}

func (p *Proccessor) sendHelp(chatID) error {
	return p.tg.SendMessage(chatID, msgHelp)
}

func (p *Proccessor) sendHello(chatID) error {
	return p.tg.SendMessage(chatID, msgHello)
}
