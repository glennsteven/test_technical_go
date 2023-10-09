package presentations

type MessageErrorValidation struct {
	Field    string   `json:"field"`
	Messages []string `json:"messages"`
}
