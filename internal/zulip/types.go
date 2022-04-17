package zulip

type verb = string

const (
	GET  verb = "GET"
	POST verb = "POST"
)

type EventType string

const (
	Messages      EventType = "messages"
	Subscriptions EventType = "subscriptions"
	RealmUser     EventType = "realm_user"
	Pointer       EventType = "pointer"
)

type Narrow string

const (
	NarrowPrivate Narrow = `[["is", "private"]]`
	NarrowAt      Narrow = `[["is", "mentioned"]]`
)

type GetAllStreamsResponse struct {
	Msg     string `json:"msg"`
	Streams []struct {
		StreamID    int    `json:"stream_id"`
		InviteOnly  bool   `json:"invite_only"`
		Description string `json:"description"`
		Name        string `json:"name"`
	} `json:"streams"`
	Result string `json:"result"`
}

type RegisterEventQueueResponse struct {
	LastEventID       int    `json:"last_event_id"`
	Msg               string `json:"msg"`
	QueueID           string `json:"queue_id"`
	Result            string `json:"result"`
	ZulipFeatureLevel int    `json:"zulip_feature_level"`
	ZulipMergeBase    string `json:"zulip_merge_base"`
	ZulipVersion      string `json:"zulip_version"`
}
