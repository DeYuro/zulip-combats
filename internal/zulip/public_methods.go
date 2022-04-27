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

func (b *Bot) RegisterEventQueue(eventList []RegisterEventType, narrow Narrow) (*Queue, error) {
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

func (b *Bot) GetEvents() ([]Event, error) {
	resp, err := b.rawGetEvents()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	events, err := b.parseEventMessages(body)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (b *Bot) SendPrivateMessage(content, email string) {
	m := Message{
		Emails:  []string{email},
		Content: content,
	}

	b.sendPrivateMessage(m)
}

func (b *Bot) parseEventMessages(body []byte) ([]Event, error) {
	rawResponse := map[string]json.RawMessage{}
	err := json.Unmarshal(body, &rawResponse)
	if err != nil {
		return nil, err
	}

	rawEvents := []map[string]json.RawMessage{}
	err = json.Unmarshal(rawResponse["events"], &rawEvents)
	if err != nil {
		return nil, err
	}

	var eventType string

	err = json.Unmarshal(rawEvents[0]["type"], &eventType)
	if err != nil {
		return nil, err
	}

	var events []Event

	newLastEventID := 0
	for _, rawEvent := range rawEvents {

		var eventType QueueEventType

		json.Unmarshal(rawEvent["type"], &eventType)
		var id int
		json.Unmarshal(rawEvent["id"], &id)

		event := Event{
			Type: eventType,
			ID:   id,
		}

		if event.ID > newLastEventID {
			newLastEventID = event.ID
		}
		if event.Type != MessageQueueType {
			events = append(events, event)
			continue
		}

		var msg EventMessage

		err = json.Unmarshal(rawEvent["message"], &msg)
		if err != nil {
			return nil, err
		}
		event.Message = msg
		events = append(events, event)
	}

	b.queue.LastEventID = newLastEventID

	return events, nil
}
