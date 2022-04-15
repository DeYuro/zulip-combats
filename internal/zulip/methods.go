package zulip

import "net/http"

func (b *Bot) getStreamList() (*http.Response, error) {
	req, err := b.request(GET, "streams", "")
	if err != nil {
		return nil, err
	}
	//client := http.Client{}
	a, e := b.client.Do(req)

	return a, e
}
