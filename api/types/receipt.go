package types

import (
	"github.com/google/uuid"
)

type Item struct {

	// The Short Product Description for the item.
	ShortDescription string `json:"shortDescription" validate:"regexp=^[\\w\\s\\-]+$"`

	// The total price payed for this item.
	Price string `json:"price" validate:"regexp=^\\d+\\.\\d{2}$"`
}

type Receipt struct {
	// Generated uuid
	Id uuid.UUID `json:"-"`

	// Calculated Points
	Points int64 `json:"-"`

	// The name of the retailer or store the receipt is from.
	Retailer string `json:"retailer" validate:"regexp=^[\\w\\s\\-&]+$"`

	// The date of the purchase printed on the receipt.
	PurchaseDate string `json:"purchaseDate"`

	// The time of the purchase printed on the receipt. 24-hour time expected.
	PurchaseTime string `json:"purchaseTime" validate:"regexp=^([01]\d|2[0-3]):([0-5]\d)$"`

	Items []Item `json:"items"`

	// The total amount paid on the receipt.
	Total string `json:"total" validate:"regexp=^\\d+\\.\\d{2}$"`
}
