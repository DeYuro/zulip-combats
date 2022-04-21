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

func (b *Bot) GetEvents() ([]EventMessage, error) {
	resp, err := b.rawGetEvents()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	msgs, err := b.parseEventMessages(body)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (b *Bot) parseEventMessages(rawEventResponse []byte) ([]EventMessage, error) {
	rawResponse := map[string]json.RawMessage{}
	err := json.Unmarshal(rawEventResponse, &rawResponse)
	if err != nil {
		return nil, err
	}

	events := []map[string]json.RawMessage{}
	err = json.Unmarshal(rawResponse["events"], &events)
	if err != nil {
		return nil, err
	}

	messages := []EventMessage{}
	newLastEventID := 0
	for _, event := range events {
		// Todo parse all Event then other
		var id int
		json.Unmarshal(event["id"], &id)
		if id > newLastEventID {
			newLastEventID = id
		}

		var msg EventMessage
		err = json.Unmarshal(event["message"], &msg)
		if err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}

	b.queue.LastEventID = newLastEventID

	return messages, nil
}
