package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	proverb := []string{}

	for i := range(rhyme) {
		if len(rhyme) == 0 {
			return []string{}
		}
		if len(rhyme) == 1 || i + 2 > len(rhyme) {
			proverb = append(proverb, "And all for the want of a " + rhyme[0] + ".")
			return proverb
		}
		
		proverb = append(proverb, "For want of a " + rhyme[i] + " the " + rhyme[i + 1] + " was lost.")
	}
	return []string{}
}
