package rules

import (
	"testing"

	"github.com/fritzyl/receipt-processor-challenge/api/types"
)

type TestReceipt struct {
	Receipt        types.Receipt
	ExpectedPoints int64
}

var testReceipts []TestReceipt = []TestReceipt{
	{
		Receipt: types.Receipt{
			Retailer:     "Target",     // 6
			PurchaseDate: "2025-01-03", // 6
			PurchaseTime: "12:00",      // 0
			Items: []types.Item{
				{
					ShortDescription: "Nabisco Cookies - 8oz", // 1
					Price:            "4.29",
				},
				{
					ShortDescription: "Ice Mountain Water - 24 pack", // 0
					Price:            "5.71",
				}, // 5 (2 items)
			},
			Total: "10.00", // 75 (50 + 25)
		},
		ExpectedPoints: 93,
	},
	{
		Receipt: types.Receipt{
			Retailer:     "Walmart Superstore - Plover WI",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []types.Item{
				{
					ShortDescription: "Mountain Dew 12PK",
					Price:            "6.49",
				},
				{
					ShortDescription: "Emils Cheese Pizza",
					Price:            "12.25",
				},
				{
					ShortDescription: "Knorr Creamy Chicken",
					Price:            "1.26",
				},
				{
					ShortDescription: "Doritos Nacho Cheese",
					Price:            "3.35",
				},
				{
					ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ",
					Price:            "12.00",
				},
			},
			Total: "35.35",
		},
		ExpectedPoints: 47,
	},
	{
		Receipt: types.Receipt{
			Retailer:     "M&M Corner Market",
			PurchaseDate: "2022-03-20",
			PurchaseTime: "14:33",
			Items: []types.Item{
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
				{
					ShortDescription: "Gatorade",
					Price:            "2.25",
				},
			},
			Total: "9.00",
		},
		ExpectedPoints: 109,
	},
}

func TestRules(t *testing.T) {
	for idx, test := range testReceipts {
		actual, _ := CalculatePoints(&test.Receipt)
		expected := test.ExpectedPoints
		if actual != expected {
			t.Errorf("Test Receipt: %d. Got %d, Want %d", idx, actual, expected)
		}
	}
}
