package receipts

import (
	"time"

	"github.com/fritzyl/receipt-processor-challenge/api/rules"
	"github.com/fritzyl/receipt-processor-challenge/api/types"
	"github.com/fritzyl/receipt-processor-challenge/api/utilities"
	"github.com/google/uuid"
	validator "gopkg.in/validator.v2"
)

var InMemoryDS map[uuid.UUID]int64 = map[uuid.UUID]int64{}

func Validate(receipt *types.Receipt) error {
	if errs := validator.Validate(receipt); errs != nil {
		return errs
	}
	// Validate Date
	_, dErr := time.Parse(utilities.DateFormat, receipt.PurchaseDate)
	if dErr != nil {
		return dErr
	}
	// Validate Time
	_, tErr := time.Parse(utilities.TimeFormat, receipt.PurchaseTime)
	if tErr != nil {
		return tErr
	}
	return nil
}

func Process(receipt *types.Receipt) (uuid.UUID, error) {
	// Gen a new uuid and attach to receipt
	generatedUuid := uuid.New()
	receipt.Id = generatedUuid
	points, err := rules.CalculatePoints(receipt)

	if err != nil {
		return generatedUuid, err
	}

	storePoints(receipt.Id, points)
	return generatedUuid, nil
}

func Lookup(lookupId string) (int64, error) {
	id, err := uuid.Parse(lookupId)
	if err != nil {
		return -1, err
	}

	found, ok := InMemoryDS[id]

	if ok {
		return found, nil
	}

	return -1, nil
}

func storePoints(id uuid.UUID, points int64) {
	InMemoryDS[id] = points
}
