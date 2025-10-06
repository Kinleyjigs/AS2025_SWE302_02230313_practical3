// shipping_v2.go
package shipping

import (
	"errors"
	"fmt"
)

// CalculateShippingFee calculates the fee based on the updated business rules.
func CalculateShippingFee(weight float64, zone string, insured bool) (float64, error) {
	// Validate weight
	if weight <= 0 || weight > 50 {
		return 0, errors.New("invalid weight")
	}

	// Determine base fee based on zone
	var baseFee float64
	switch zone {
	case "Domestic":
		baseFee = 5.0
	case "International":
		baseFee = 20.0
	case "Express":
		baseFee = 30.0
	default:
		return 0, fmt.Errorf("invalid zone: %s", zone)
	}

	// Determine heavy surcharge
	var heavySurcharge float64
	if weight > 10 {
		heavySurcharge = 7.50
	}

	// Subtotal before insurance
	subTotal := baseFee + heavySurcharge

	// Apply insurance cost if applicable
	var insuranceCost float64
	if insured {
		insuranceCost = subTotal * 0.015
	}

	finalTotal := subTotal + insuranceCost

	return finalTotal, nil
}
