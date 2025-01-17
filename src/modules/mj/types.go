package mj

const (
	ApplicationID string = "936929561302675456"
	SessionID     string = "ea8816d857ba9ae2f74c59ae1a953afe"
)

type InteractionsRequest struct {
	Type          int            `json:"type"`
	ApplicationID string         `json:"application_id"`
	MessageFlags  *int           `json:"message_flags,omitempty"`
	MessageID     *string        `json:"message_id,omitempty"`
	GuildID       string         `json:"guild_id"`
	ChannelID     string         `json:"channel_id"`
	SessionID     string         `json:"session_id"`
	Data          map[string]any `json:"data"`
}
