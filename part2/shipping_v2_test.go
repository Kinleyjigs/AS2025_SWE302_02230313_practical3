package shipping

import "testing"

func TestCalculateShippingFee_V2(t *testing.T) {
	testCases := []struct {
		name        string
		weight      float64
		zone        string
		insured     bool
		expectedFee float64
		expectError bool
	}{
		// ---- Invalid Weights (P1, P4) ----
		{"Invalid: weight = 0", 0, "Domestic", false, 0, true},
		{"Invalid: weight < 0", -2, "Domestic", true, 0, true},
		{"Invalid: weight > 50", 60, "Express", false, 0, true},

		// ---- Invalid Zones (P6) ----
		{"Invalid: unknown zone", 10, "Local", false, 0, true},
		{"Invalid: empty zone", 10, "", true, 0, true},

		// ---- Standard Tier (0 < w ≤ 10) (P2) ----
		{"Standard Domestic uninsured", 5, "Domestic", false, 5.00, false},
		{"Standard Domestic insured", 5, "Domestic", true, 5.00 * 1.015, false},

		{"Standard International uninsured", 10, "International", false, 20.00, false},
		{"Standard International insured", 10, "International", true, 20.00 * 1.015, false},

		{"Standard Express uninsured", 8, "Express", false, 30.00, false},
		{"Standard Express insured", 8, "Express", true, 30.00 * 1.015, false},

		// ---- Heavy Tier (10 < w ≤ 50) (P3) ----
		{"Heavy Domestic uninsured", 20, "Domestic", false, 5.00 + 7.50, false},
		{"Heavy Domestic insured", 20, "Domestic", true, (5.00 + 7.50) * 1.015, false},

		{"Heavy International uninsured", 50, "International", false, 20.00 + 7.50, false},
		{"Heavy International insured", 50, "International", true, (20.00 + 7.50) * 1.015, false},

		{"Heavy Express uninsured", 12, "Express", false, 30.00 + 7.50, false},
		{"Heavy Express insured", 12, "Express", true, (30.00 + 7.50) * 1.015, false},

		// ---- Boundary Values ----
		{"Boundary: 0 (invalid)", 0, "Domestic", false, 0, true},
		{"Boundary: 0.1 (first valid)", 0.1, "Domestic", false, 5.0, false},
		{"Boundary: 10 (Standard limit)", 10, "Domestic", false, 5.0, false},
		{"Boundary: 10.1 (Heavy start)", 10.1, "Domestic", false, 5.0 + 7.5, false},
		{"Boundary: 50 (max valid)", 50, "Domestic", false, 5.0 + 7.5, false},
		{"Boundary: 50.1 (invalid)", 50.1, "Domestic", false, 0, true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := CalculateShippingFee(tc.weight, tc.zone, tc.insured)

			if tc.expectError {
				if err == nil {
					t.Errorf("Expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if got != tc.expectedFee {
				t.Errorf("Expected %.4f, got %.4f", tc.expectedFee, got)
			}
		})
	}
}
