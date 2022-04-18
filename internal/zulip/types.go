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
	Msg     string   `json:"msg"`
	Streams []Stream `json:"streams"`
	Result  string   `json:"result"`
}

type Stream struct {
	StreamID    int    `json:"stream_id"`
	InviteOnly  bool   `json:"invite_only"`
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Queue struct {
	Result            string `json:"result"`
	Msg               string `json:"msg"`
	QueueID           string `json:"queue_id"`
	ZulipVersion      string `json:"zulip_version"`
	ZulipFeatureLevel int    `json:"zulip_feature_level"`
	MaxMessageID      int    `json:"max_message_id"`
	LastEventID       int    `json:"last_event_id"`
}

type EventMessage struct {
	AvatarURL       string        `json:"avatar_url"`
	Client          string        `json:"client"`
	Content         string        `json:"content"`
	ContentType     string        `json:"content_type"`
	GravatarHash    string        `json:"gravatar_hash"`
	ID              int           `json:"id"`
	RecipientID     int           `json:"recipient_id"`
	SenderDomain    string        `json:"sender_domain"`
	SenderEmail     string        `json:"sender_email"`
	SenderFullName  string        `json:"sender_full_name"`
	SenderID        int           `json:"sender_id"`
	SenderShortName string        `json:"sender_short_name"`
	Subject         string        `json:"subject"`
	SubjectLinks    []interface{} `json:"subject_links"`
	Timestamp       int           `json:"timestamp"`
	Type            string        `json:"type"`
}
