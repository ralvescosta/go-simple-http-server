package models

type (
	ReversalRequest struct {
		Mti            string `json:"mti" validate:"required"`
		ProcessingCode string `json:"processing_code" validate:"required"`
		Amount         string `json:"amount" validate:"required"`
		EntryMode      string `json:"entry_mode" validate:"required"`
		Track2         string `json:"track2" validate:"required"`
		TerminalID     string `json:"terminal_id" validate:"required"`
		MerchantID     string `json:"merchant_id" validate:"required"`
	}

	ReversalResponse struct {
		ResponseCode string `json:"response_code"`
	}
)
