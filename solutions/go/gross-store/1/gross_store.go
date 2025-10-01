package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {

	for s := range units {
		if s == unit {
			billValue := units[unit]
			bill[item] += billValue
			return true
		}
	}
	return false

	// billItem, unitsUnit := units[unit]

	// if !unitsUnit {
	// 	return false
	// }

	// bill[item] += billItem
	// return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	removeBillItem, removeUnitsUnit := units[unit]
	billItem, unitsUnit := bill[item]

	if removeBillItem > billItem || !removeUnitsUnit || !unitsUnit {
		return false
	} else if removeBillItem == billItem {
		delete(bill, item)
	} else {
		bill[item] -= removeBillItem
	}
	return true
}


// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	for searchedItem, quantity := range bill {
		switch {
		case searchedItem == item:
			return quantity, true
		}
	}
	return 0, false
}