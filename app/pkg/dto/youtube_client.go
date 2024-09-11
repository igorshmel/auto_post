package dto

// Video --
type Video struct {
	ID          string
	Title       string
	Description string
	ChannelID   string
}

// Comment --
type Comment struct {
	ID      string
	VideoID string
	Text    string
}

// Channel --
type Channel struct {
	ID              string
	Name            string
	Description     string
	SubscriberCount int
}
