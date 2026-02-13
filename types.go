package main

// ToWhom matches the JSON response from GET /api/v1/towhom.
type ToWhom struct {
	TotalCountPubkeys            int64 `json:"total_count_pubkeys"`
	TotalCountLightningAddresses int64 `json:"total_count_lightning_addresses"`
	TotalCount                   int64 `json:"total_count"`
	TotalSatogramsSent           int64 `json:"total_satograms_sent"`
	TotalSignedUp                int64 `json:"total_signed_up"`
}

// SatogramPayload is the request body for POST /api/v1/satogram.
type SatogramPayload struct {
	TotalCost          int64   `json:"total_cost"`
	AmtPerSatogram     *int64  `json:"amt_per_satogram,omitempty"`
	MaxFees            *int64  `json:"max_fees,omitempty"`
	Message            string  `json:"message"`
	SenderAddress      *string `json:"sender_address,omitempty"`
	RecipientSelection *string `json:"recipient_selection,omitempty"`
}

// SatogramReturn is the response from POST /api/v1/satogram.
type SatogramReturn struct {
	PaymentRequest string `json:"payment_request"`
}

// PaymentStatus is the response from GET /api/v1/invoice/status/{payment_request}.
// The upstream struct has NO json tags, so fields serialize as PascalCase.
type PaymentStatus struct {
	PaymentRequest string `json:"PaymentRequest"`
	Status         string `json:"Status"`
}

// SatogramProgress holds delivery progress details.
type SatogramProgress struct {
	TargetMarket  string `json:"target_market"`
	TotalSuccess  int64  `json:"success_count"`
	TotalFailure  int64  `json:"failure_count"`
	TotalKickoff  int64  `json:"kickoff_count"`
	TotalSkipped  int64  `json:"skipped_count"`
	TotalExpired  int64  `json:"expired_count"`
	SatSpentSoFar int64  `json:"sats_spent_so_far"`
	Status        int    `json:"satogram_status"`
}

// SatogramStored is the response from GET /api/v1/satogram/status/{payment_request}.
type SatogramStored struct {
	PaymentRequest   string           `json:"payment_request"`
	SatogramPayload  SatogramPayload  `json:"satogram_payload"`
	SatogramProgress SatogramProgress `json:"satogram_progress"`
	InvoiceStates    int32            `json:"invoice_states"`
	CreationTime     string           `json:"creation_time"`
	StartProcessTime *string          `json:"start_process_time"`
	EndProcessTime   *string          `json:"end_process_time"`
}
