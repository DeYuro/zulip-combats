package zulip

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

func (b *Bot) getStreamList() (*http.Response, error) {
	req, err := b.request(GET, "streams", "")
	if err != nil {
		return nil, err
	}

	return b.client.Do(req)
}

func (b *Bot) registerEventQueue(eventTypes []RegisterEventType, narrow Narrow) (*http.Response, error) {
	query := `event_types=["message"]`

	if len(eventTypes) != 0 {
		query = `event_types=["`
		for i, s := range eventTypes {
			query += fmt.Sprintf("%s", s)
			if i != len(eventTypes)-1 {
				query += `", "`
			}
		}
		query += `"]`
	}

	if narrow != "" {
		query += fmt.Sprintf("&narrow=%s", narrow)
	}

	req, err := b.request(POST, "register", query)

	if err != nil {
		return nil, err
	}

	return b.client.Do(req)
}

func (b *Bot) rawGetEvents() (*http.Response, error) {
	values := url.Values{}
	values.Set("queue_id", b.queue.QueueID)
	values.Set("last_event_id", strconv.Itoa(b.queue.LastEventID))

	url := "events?" + values.Encode()

	req, err := b.request(GET, url, "")
	if err != nil {
		return nil, err
	}

	return b.client.Do(req)
}
