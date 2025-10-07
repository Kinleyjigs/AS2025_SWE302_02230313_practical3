# Practical 3 Advanced Shipping Fee Calculator 

## Goal

Implement and test an updated `CalculateShippingFee` function in Go. The exercise focuses on deriving test cases using Equivalence Partitioning (EP) and Boundary Value Analysis (BVA), then translating that design into robust, table-driven tests.

## Requirements (spec summary)

- Function signature: `CalculateShippingFee(weight float64, zone string, insured bool) (float64, error)`
- `weight` must be > 0 and ≤ 50
  - Standard tier: 0 < weight ≤ 10 (no surcharge)
  - Heavy tier: 10 < weight ≤ 50 (add $7.50 surcharge)
- `zone` must be one of: `"Domestic"`, `"International"`, `"Express"`
- `insured`: if true, add 1.5% of (base fee + surcharge)

Base fees (inferred from implementation):
- Domestic: $5.00
- International: $20.00
- Express: $30.00

Errors should be returned for invalid weights and invalid zones.

### Equivalence Partitioning (EP)

We partition inputs into equivalence classes that should behave similarly.

- Weight (float):
  - P1: Invalid - weight ≤ 0 (examples: 0, -1)
  - P2: Valid- Standard tier (0 < weight ≤ 10) (examples: 5, 10)
  - P3: Valid- Heavy tier (10 < weight ≤ 50) (examples: 20, 50)
  - P4: Invalid- weight > 50 (example: 60)

- Zone (string):
  - P5: Valid- `"Domestic"`, `"International"`, `"Express"`
  - P6: Invalid- anything else (example: `"Local"`, `""`)

- Insured (bool):
  - P7: Not insured- `false`
  - P8: Insured- `true` (adds 1.5% of subtotal)

### Boundary Value Analysis (BVA)

We select tests at edges of numeric ranges:

- Lower boundary: `0` (invalid) and `0.1` (first valid)
- Mid boundary: `10` (Standard upper) and `10.1` (Heavy lower)
- Upper boundary: `50` (valid) and `50.1` (invalid)

### Implementation

- `part2/shipping_v2.go` implements `CalculateShippingFee` with the rules above. Key behaviors:
  - Validates weight: returns error if weight ≤ 0 or weight > 50
  - Validates zone: returns error if not one of the three allowed strings
  - Applies base fee per zone
  - Applies heavy surcharge ($7.50) when weight > 10
  - Applies insurance as 1.5% of subtotal when `insured == true`

Source: `part2/shipping_v2.go`

### Test design and coverage

- `part2/shipping_v2_test.go` contains a single table-driven test: `TestCalculateShippingFee_V2`.
- Test table includes cases for:
  - Invalid weights (0, negative, >50)
  - Invalid zones (unknown, empty)
  - Standard tier (various zones), insured and uninsured
  - Heavy tier (various zones), insured and uninsured
  - Boundary values: 0, 0.1, 10, 10.1, 50, 50.1

Each case includes: name, weight, zone, insured, expectedFee, expectError.

Notes: current tests compare floating values with direct equality. For this problem the arithmetic is simple and stable; however using an epsilon (e.g., 1e-9) is more robust for future changes.

### Test run (local)

I ran the test suite in `part2` locally. All tests passed.

![alt text](<images/Screenshot 2025-10-07 at 1.20.26 AM.png>)

Sample run command :

```
go test -v
```


## Conclusion

In this practical, I learned how to design and implement tests based only on a system’s specification using black-box testing techniques.
By applying Equivalence Partitioning and Boundary Value Analysis, I identified valid and invalid input ranges and key edge cases.
Then, I created a table-driven test suite in Go to cover all partitions and boundaries.
All tests passed successfully, proving that the implementation meets the given business rules.

