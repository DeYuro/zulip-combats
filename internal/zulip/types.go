package zulip

type verb = string

const (
	GET  verb = "GET"
	POST verb = "POST"
)

type StreamJSON struct {
	Msg     string `json:"msg"`
	Streams []struct {
		StreamID    int    `json:"stream_id"`
		InviteOnly  bool   `json:"invite_only"`
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"streams"`
	Result string `json:"result"`
}
