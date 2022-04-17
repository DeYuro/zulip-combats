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

func (b *Bot) RegisterEventQueue() (*RegisterEventQueueResponse, error) {
	resp, err := b.registerEventQueue()
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response *RegisterEventQueueResponse

	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
