package presentations

type PayloadAddTransaction struct {
	IDConsumer        string  `json:"id_consumer"`
	ContractNumber    string  `json:"contract_number"`
	OTR               float64 `json:"otr"`
	FeeAdmin          float64 `json:"fee_admin"`
	InstallmentAmount int     `json:"installment_amount"`
	TotalInterest     int     `json:"total_interest"`
	AssetName         string  `json:"asset_name"`
}
