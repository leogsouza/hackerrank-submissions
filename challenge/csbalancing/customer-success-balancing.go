package csbalancing

import "fmt"

// Complete the csRush function below.
func csRush(n int32, m int32, css [][]int32, customers [][]int32, vacant_css []int32) int32 {

	// Creating map for customers
	mapCss := make(map[int32]int32)
	for _, cs := range css {
		mapCss[cs[0]] = cs[1]
	}

	// Remove customers on vacation
	for _, vc := range vacant_css {
		delete(mapCss, vc)
	}

	mapCustomers := make(map[int32]int32)

	for _, customer := range customers {
		mapCustomers[customer[0]] = customer[1]
	}

	mapCssAttend := make(map[int32]int32)
	mapCustomerAttended := make(map[int32]bool)

	for k, csCapacity := range mapCss {
		for kc, customerSize := range mapCustomers {
			_, ok := mapCustomerAttended[kc]
			if csCapacity >= customerSize && !ok {
				mapCssAttend[k]++
				mapCustomerAttended[kc] = true
			}
		}
	}

	var max, customerId int32
	fmt.Println(mapCssAttend)
	for k, totalAttend := range mapCssAttend {

		if totalAttend > max {
			max = totalAttend
			customerId = k
		}
	}

	return customerId

}
