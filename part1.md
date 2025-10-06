# Part 1: Test Case Design (Analysis)

This section analyzes the specification for the updated `CalculateShippingFee` function using **Equivalence Partitioning (EP)** and **Boundary Value Analysis (BVA)**.

---

## Function Overview

**Signature:**  
`CalculateShippingFee(weight float64, zone string, insured bool) (float64, error)`

**Specification Summary:**
- `weight`: must be > 0 and ≤ 50  
  - (0, 10] → Standard tier  
  - (10, 50] → Heavy tier (+ $7.50 surcharge)
- `zone`: must be one of `"Domestic"`, `"International"`, or `"Express"`
- `insured`: if true, adds 1.5% of (base fee + surcharge)

---

##  1. Equivalence Partitioning (EP)

Identify sets (partitions) of input values expected to behave the same.

###  Input 1: `weight`

| Partition | Description | Example |
|------------|--------------|----------|
| **P1** | Invalid — weight ≤ 0 | `0`, `-1` |
| **P2** | Valid — Standard tier (0 < weight ≤ 10) | `5`, `10` |
| **P3** | Valid — Heavy tier (10 < weight ≤ 50) | `20`, `50` |
| **P4** | Invalid — weight > 50 | `60` |

---

### Input 2: `zone`

| Partition | Description | Example |
|------------|--------------|----------|
| **P5** | Valid zone | `"Domestic"`, `"International"`, `"Express"` |
| **P6** | Invalid zone | `"Local"`, `"domestic"`, `""` |

---

### Input 3: `insured`

| Partition | Description | Example |
|------------|--------------|----------|
| **P7** | Not insured | `false` |
| **P8** | Insured | `true` |

---

##  2. Boundary Value Analysis (BVA)

We test the **edges** between valid and invalid inputs (only for numeric fields).

### Input: `weight`

| Boundary | Description | Test Values |
|-----------|--------------|--------------|
| **Lower Boundary** | Transition from invalid to valid | `0` (invalid), `0.1` (valid) |
| **Mid Boundary** | Transition between Standard ↔ Heavy tiers | `10` (Standard), `10.1` (Heavy) |
| **Upper Boundary** | Transition from valid to invalid | `50` (valid), `50.1` (invalid) |

---

**Summary:**
- **Weight**: 4 partitions, 3 boundaries  
- **Zone**: 2 partitions  
- **Insured**:  2 partitions  
- Total coverage ensures every input type and edge case is represented.

