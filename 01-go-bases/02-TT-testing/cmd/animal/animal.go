package animal

func Animal(animalType string) (func(int) float64, string) {
	switch animalType {
	case "dog":
		return AnimalDog, ""
	case "cat":
		return AnimalCat, ""
	case "hamster":
		return AnimalHamster, ""
	case "tarantula":
		return AnimalTarantula, ""
	default:
		return nil, "animal not available. the options are dog, cat, hamster or tarantula"
	}
}

func AnimalDog(amount int) float64 {
	return float64(amount) * 10.0
}

func AnimalCat(amount int) float64 {
	return float64(amount) * 5.0
}

func AnimalHamster(amount int) float64 {
	return float64(amount) * 0.250
}

func AnimalTarantula(amount int) float64 {
	return float64(amount) * 0.150
}
