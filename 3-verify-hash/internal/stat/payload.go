package stat

type ClicksByPeriod struct {
	Period string
	Clicks int64
}

type GetStatResponse struct {
	Data []ClicksByPeriod `json:"data"`
}
