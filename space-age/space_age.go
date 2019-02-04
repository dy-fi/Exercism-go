package space

// Age returns the age a person would be on a given planet, given the seconds they were alive but not accounting for relativity
func Age(seconds float64, planet string) float64 {
	switch planet {
	case "Earth":
		return seconds / 31557600.00
	case "Mercury":
		return seconds / 7600543.81992
	case "Venus":
		return seconds / 19414149.052176
	case "Mars":
		return seconds / 59354032.690079994
	case "Jupiter":
		return seconds / 374355659.124
	case "Saturn":
		return seconds / 929292362.8848
	case "Uranus":
		return seconds / 2651370019.3296
	case "Neptune":
		return seconds / 5200418560.032001
	default:
		return -1
	}
}
