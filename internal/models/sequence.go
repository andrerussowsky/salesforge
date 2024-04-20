package models

// Sequence represents a sequence of emails
type Sequence struct {
    ID                   int    `json:"id"`
    Name                 string `json:"name"`
    OpenTrackingEnabled  bool   `json:"open_tracking_enabled"`
    ClickTrackingEnabled bool   `json:"click_tracking_enabled"`
    Steps                []Step `json:"steps"`
}

// Step represents an email step in a sequence
type Step struct {
    ID           int    `json:"id"`
    SequenceID   int    `json:"sequence_id"`
    EmailSubject string `json:"email_subject"`
    EmailContent string `json:"email_content"`
}

// Tracking represents the tracking settings for a sequence
type Tracking struct {
    OpenTracking  bool `json:"open_tracking"`
    ClickTracking bool `json:"click_tracking"`
}
