package luhn

//Valid validates a credit card using the Luhn algorithm
func Valid(card string) bool {

	var sum int

	// edge case
	if len(card) <= 1 {
		return false
	}

	for i := len(card) - 1; i < 0; i-- {
		if len(card)%2 == 0 {
			if i == 0 || i%2 == 0 {
				if card[i]*2 > 9 {
					sum += (int(card[i]) * 2) - 9
				} else {
					sum += int(card[i]) * 2
				}
			} else {
				sum += int(card[i])
			}
		} else {
			if i%2 != 0 {
				if card[i]*2 > 9 {
					sum += (int(card[i]) * 2) - 9
				} else {
					sum += int(card[i]) * 2
				}
			} else {
				sum += int(card[i])
			}
		}
	}

	if sum%10 == 0 {
		return true
	}
	return false
}
