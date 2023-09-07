package animal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnimal(t *testing.T) {

	dog := "dog"
	cat := "cat"
	hamster := "hamster"
	tarantula := "tarantula"

	// Dog
	animalDog, err := Animal(dog)
	if err != "" {
		t.Fatal("unexpected error:", err)
	}

	animalDogResult := animalDog(10)
	dogResultExpected := 100.0

	assert.Equal(t, dogResultExpected, animalDogResult, err)

	//Cat
	animalCat, err := Animal(cat)
	if err != "" {
		t.Fatal("unexpected error:", err)
	}

	animalCatResult := animalCat(10)
	catResultExpected := 50.0
	assert.Equal(t, catResultExpected, animalCatResult, err)

	//Hamster
	animalHamster, err := Animal(hamster)
	if err != "" {
		t.Fatal("unexpected error:", err)
	}

	animalHamsterResult := animalHamster(10)
	hamsterResultExpected := 2.5
	assert.Equal(t, hamsterResultExpected, animalHamsterResult, err)

	//Tarantula
	animalTarantula, err := Animal(tarantula)
	if err != "" {
		t.Fatal("unexpected error:", err)
	}

	animalTarantulaResult := animalTarantula(10)
	tarantulaResultExpected := 1.5
	assert.Equal(t, tarantulaResultExpected, animalTarantulaResult, err)
}
