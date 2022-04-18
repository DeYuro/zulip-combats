package zulip

import (
	"encoding/json"
	"io/ioutil"
)

func (b *Bot) GetStreams() ([]string, error) {
	resp, err := b.getStreamList()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var sj GetAllStreamsResponse
	err = json.Unmarshal(body, &sj)
	if err != nil {
		return nil, err
	}

	var streams []string
	for _, s := range sj.Streams {
		streams = append(streams, s.Name)
	}

	return streams, nil
}
func (b *Bot) RegisterEventQueuePrivate() (*Queue, error) {
	return b.RegisterEventQueue(nil, NarrowPrivate)
}

func (b *Bot) RegisterEventQueue(eventList []EventType, narrow Narrow) (*Queue, error) {
	resp, err := b.registerEventQueue(eventList, narrow)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var queue Queue

	err = json.Unmarshal(body, &queue)
	if err != nil {
		return nil, err
	}

	return &queue, nil
}

func (b *Bot) GetEventChan() {

}
