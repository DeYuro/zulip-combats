package zulip

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"sync/atomic"
	"time"
)

func (q *Queue) EventsChan() (chan EventMessage, func()) {
	end := false
	endFunc := func() {
		end = true
	}

	out := make(chan EventMessage)
	go func() {
		defer close(out)
		for {
			backoffTime := time.Now().Add(q.Bot.Backoff * time.Duration(math.Pow10(int(atomic.LoadInt64(&q.Bot.Retries)))))
			minTime := time.Now().Add(q.Bot.Backoff)
			if end {
				return
			}
			ems, err := q.GetEvents()
			switch {
			case err == HeartbeatError:
				time.Sleep(time.Until(minTime))
				continue
			case err == BackoffError:
				time.Sleep(time.Until(backoffTime))
				atomic.AddInt64(&q.Bot.Retries, 1)
				continue
			case err == UnauthorizedError:
				// TODO? have error channel when ending the continuously running process?
				return
			default:
				atomic.StoreInt64(&q.Bot.Retries, 0)
			}
			if err != nil {
				// TODO: handle unknown error
				// For now, handle this like an UnauthorizedError and end the func.
				return
			}
			for _, em := range ems {
				out <- em
			}
			// Always make sure we wait the minimum time before asking again.
			time.Sleep(time.Until(minTime))
		}
	}()

	return out, endFunc
}

func (q *Queue) GetEvents() ([]EventMessage, error) {
	resp, err := q.RawGetEvents()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	switch {
	case resp.StatusCode == 429:
		return nil, BackoffError
	case resp.StatusCode == 403:
		return nil, UnauthorizedError
	case resp.StatusCode >= 400:
		return nil, UnknownError
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	msgs, err := q.ParseEventMessages(body)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

// RawGetEvents is a blocking call that receives a response containing a list
// of events (a.k.a. received messages) since the last message id in the queue.
func (q *Queue) RawGetEvents() (*http.Response, error) {
	values := url.Values{}
	values.Set("queue_id", q.ID)
	values.Set("last_event_id", strconv.Itoa(q.LastEventID))

	url := "events?" + values.Encode()

	req, err := q.Bot.constructRequest("GET", url, "")
	if err != nil {
		return nil, err
	}

	return q.Bot.Client.Do(req)
}

func (q *Queue) ParseEventMessages(rawEventResponse []byte) ([]EventMessage, error) {
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
		// Update the lastEventID
		var id int
		json.Unmarshal(event["id"], &id)
		if id > newLastEventID {
			newLastEventID = id
		}

		// If the event is a heartbeat, there won't be any more events.
		// So update the last event id and return a special error.
		if string(event["type"]) == `"heartbeat"` {
			q.LastEventID = newLastEventID
			return nil, HeartbeatError
		}
		var msg EventMessage
		err = json.Unmarshal(event["message"], &msg)
		// TODO? should this check be here
		if err != nil {
			return nil, err
		}
		msg.Queue = q
		messages = append(messages, msg)
	}

	q.LastEventID = newLastEventID

	return messages, nil
}
