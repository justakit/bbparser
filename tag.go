package bbparser

// Tag - structure contains info about bbcode tag
type Tag struct {
	Name       string
	Closing    bool
	Attributes map[string]string
}
