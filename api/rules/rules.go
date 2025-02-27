package rules

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/fritzyl/receipt-processor-challenge/api/types"
	"github.com/fritzyl/receipt-processor-challenge/api/utilities"
)

type Rule interface {
	Evaluate()
}

type RetailerAlphaNumericRule struct {
	Receipt *types.Receipt
	Name    string
	Points  int64
}

func (rule RetailerAlphaNumericRule) Evaluate() {
	alphaNumericCount := utilities.AlphaNumericCount(rule.Receipt.Retailer)
	awardedPoints := alphaNumericCount * rule.Points
	IncrementPoints(rule.Receipt, awardedPoints, rule.Name)
}

type TotalModuloRule struct {
	Receipt *types.Receipt
	Name    string
	Points  int64
	Divisor float64
}

func (rule TotalModuloRule) Evaluate() {
	var awardedPoints int64 = 0
	floatTotal, _ := strconv.ParseFloat(rule.Receipt.Total, 64)
	if math.Mod(floatTotal, rule.Divisor) == 0 {
		awardedPoints = rule.Points
	}
	IncrementPoints(rule.Receipt, awardedPoints, rule.Name)
}

type ItemCountRule struct {
	Receipt *types.Receipt
	Name    string
	Points  int64
}

func (rule ItemCountRule) Evaluate() {
	itemCount := float64(len(rule.Receipt.Items))
	awardedPoints := int64(math.Floor(itemCount/2) * float64(rule.Points))
	IncrementPoints(rule.Receipt, awardedPoints, rule.Name)
}

type ItemDescriptionLengthRule struct {
	Receipt *types.Receipt
	Name    string
	Points  float64
}

func (rule ItemDescriptionLengthRule) Evaluate() {
	var awardedPoints int64 = 0

	for _, item := range rule.Receipt.Items {
		itemPoints := 0.0
		descLength := len(strings.Trim(item.ShortDescription, " "))

		if descLength%3 == 0 {
			floatPrice, _ := strconv.ParseFloat(item.Price, 64)
			itemPoints += math.Ceil(floatPrice * rule.Points)
		}
		awardedPoints += int64(itemPoints)
	}

	IncrementPoints(rule.Receipt, awardedPoints, rule.Name)
}

type TimeCompareRule struct {
	Receipt    *types.Receipt
	Name       string
	Points     int64
	LowerBound string
	UpperBound string
}

func (rule TimeCompareRule) Evaluate() {
	var awardedPoints int64 = 0
	lowerBound := utilities.GetTime(rule.LowerBound)
	upperBound := utilities.GetTime(rule.UpperBound)

	var evalLowerBound bool = utilities.CompareTime(utilities.GetTime(rule.Receipt.PurchaseTime), ">", lowerBound)
	var evalUpperBound bool = utilities.CompareTime(utilities.GetTime(rule.Receipt.PurchaseTime), "<", upperBound)
	if evalLowerBound && evalUpperBound {
		awardedPoints = rule.Points
	}
	IncrementPoints(rule.Receipt, awardedPoints, rule.Name)
}

type OddDateRule struct {
	Receipt *types.Receipt
	Name    string
	Points  int64
}

func (rule OddDateRule) Evaluate() {
	var awardedPoints int64 = 0
	day := strings.Split(rule.Receipt.PurchaseDate, "-")[2]
	if d, _ := strconv.Atoi(day); d%2 != 0 {
		awardedPoints = rule.Points
	}
	IncrementPoints(rule.Receipt, awardedPoints, rule.Name)
}

type LLMGeneratedRule struct {
	Receipt     *types.Receipt
	Name        string
	Points      int64
	IsGenerated bool
}

func (rule LLMGeneratedRule) Evaluate() {
	var awardedPoints int64 = 0
	floatTotal, _ := strconv.ParseFloat(rule.Receipt.Total, 64)
	if rule.IsGenerated && floatTotal > 10.00 {
		awardedPoints = rule.Points
	}
	IncrementPoints(rule.Receipt, awardedPoints, rule.Name)
}

func IncrementPoints(receipt *types.Receipt, points int64, note string) {
	fmt.Println(note, ":", points)
	receipt.Points += points
}

func CalculatePoints(receipt *types.Receipt) (int64, error) {
	var ruleSet []Rule
	// One point for every alphanumeric character in the retailer name.
	ruleSet = append(ruleSet, RetailerAlphaNumericRule{Receipt: receipt, Points: 1, Name: "Retailer Name Rule"})
	// 50 points if the total is a round dollar amount with no cents.
	ruleSet = append(ruleSet, TotalModuloRule{Receipt: receipt, Divisor: 1.00, Points: 50, Name: "Total Round Dollar Rule"})
	// 25 points if the total is a multiple of `0.25`.
	ruleSet = append(ruleSet, TotalModuloRule{Receipt: receipt, Divisor: 0.25, Points: 25, Name: "Total Multiple of 0.25 Rule"})
	// 5 points for every two items on the receipt.
	ruleSet = append(ruleSet, ItemCountRule{Receipt: receipt, Points: 5, Name: "Item Count Rule"})
	// If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. The result is the number of points earned.
	ruleSet = append(ruleSet, ItemDescriptionLengthRule{Receipt: receipt, Points: 0.2, Name: "Item Description Rule"})
	// If and only if this program is generated using a large language model, 5 points if the total is greater than 10.00.
	ruleSet = append(ruleSet, LLMGeneratedRule{Receipt: receipt, IsGenerated: "Yep" == "Nope", Points: 5, Name: "Large Language Model Gen Rule"})
	// 6 points if the day in the purchase date is odd.
	ruleSet = append(ruleSet, OddDateRule{Receipt: receipt, Points: 6, Name: "Odd Date Rule"})
	// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	ruleSet = append(ruleSet, TimeCompareRule{Receipt: receipt, LowerBound: "14:00", UpperBound: "16:00", Points: 10, Name: "Time of Day Rule"})

	for _, rule := range ruleSet {
		rule.Evaluate()
	}
	fmt.Println("----------------------\nCalculated Points:", receipt.Points)
	return receipt.Points, nil
}
