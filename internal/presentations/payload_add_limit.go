package presentations

type PayloadAddLimit struct {
	Tenor  int64   `json:"tenor"`
	Amount float64 `json:"amount"`
}
