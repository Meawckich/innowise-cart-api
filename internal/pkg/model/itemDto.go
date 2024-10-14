package model

import (
	"encoding/json"
)

type ItemDto struct {
	Product  string `json:"product"`
	Quantity int    `json:"quantity"`
}

func (dto *ItemDto) UnmarshalJSON(data []byte) error {
	type Alias ItemDto // Create an alias to avoid recursion
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(dto),
	}

	// Try to unmarshal into the auxiliary struct
	if err := json.Unmarshal(data, &aux); err != nil {
		return NewInvalidBodyError("malformed JSON: "+err.Error(), string(data))
	}

	// Validate required fields
	if dto.Product == "" && dto.Quantity <= 0 {
		return NewInvalidBodyError("Product and quantity need to be provided", string(data))
	}

	return nil
}
