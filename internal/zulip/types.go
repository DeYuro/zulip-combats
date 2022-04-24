package zulip

type verb = string

const (
	GET  verb = "GET"
	POST verb = "POST"
)

type QueueEventType string

const (
	MessageQueueType   QueueEventType = "message"
	HeartbeatQueueType QueueEventType = "heartbeat"
)

type RegisterEventType string

const (
	Messages      RegisterEventType = "messages"
	Subscriptions RegisterEventType = "subscriptions"
	RealmUser     RegisterEventType = "realm_user"
	Pointer       RegisterEventType = "pointer"
	Heartbeat     RegisterEventType = "heartbeat"
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

type Event struct {
	ID      int
	Type    QueueEventType
	Message EventMessage
}

type EventMessage struct {
	AvatarURL       string        `json:"avatar_url,omitempty"`
	Client          string        `json:"client,omitempty"`
	Content         string        `json:"content,omitempty"`
	ContentType     string        `json:"content_type,omitempty"`
	GravatarHash    string        `json:"gravatar_hash,omitempty"`
	ID              int           `json:"id"`
	RecipientID     int           `json:"recipient_id,omitempty"`
	SenderDomain    string        `json:"sender_domain,omitempty"`
	SenderEmail     string        `json:"sender_email,omitempty"`
	SenderFullName  string        `json:"sender_full_name,omitempty"`
	SenderID        int           `json:"sender_id,omitempty"`
	SenderShortName string        `json:"sender_short_name,omitempty"`
	Subject         string        `json:"subject,omitempty"`
	SubjectLinks    []interface{} `json:"subject_links,omitempty"`
	Timestamp       int           `json:"timestamp,omitempty"`
	Type            string        `json:"type"`
}
