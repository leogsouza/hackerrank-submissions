package main

/*
 * Complete the 'writeIn' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING_ARRAY ballot as parameter.
 */

func writeIn(ballot []string) string {
	candidates := make(map[string]int)
	var winner string
	var total int

	for _, candidate := range ballot {
		candidates[candidate]++
	}

	for name, votes := range candidates {
		if votes > total {
			total = votes
			winner = name
		} else if votes == total {
			if name > winner {
				winner = name
			}
		}
	}

	return winner
}
