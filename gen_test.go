package randomwords

import "testing"

func TestGenerate(t *testing.T) {
	t.Log("Ten Default")
	gen := DefaultGenerator()
	for i := 0; i <= 10; i++ {
		rand := gen.Random()
		t.Log(rand)
	}
	t.Logf("Num Combinations: %d", gen.EstimatedNumCombinations())

	t.Log("------")
	t.Log("Ten AdjAdjNoun")
	an := AdjNounGenerator()
	for i := 0; i <= 10; i++ {
		rand := an.Random()
		t.Log(rand)
	}
	t.Logf("Num Combinations: %d", an.EstimatedNumCombinations())

	custom := &Generator{
		Categories: map[string]*Category{
			CosmicThings.Name: CosmicThings,
			Fortresses.Name:   Fortresses,
			"cities": &Category{
				Name: "cities",
				Words: []string{
					"LosAngeles",
					"TelAviv",
					"Cologne",
					"Vancouver",
					"Dallas"},
			},
		},
		DefaultCategoryLayout: []string{CosmicThings.Name, "cities", Fortresses.Name},
	}
	t.Log("------")
	t.Log("Ten Custom")
	for i := 0; i <= 10; i++ {
		rand := custom.Random()
		t.Log(rand)
	}
	t.Logf("Num Combinations: %d", custom.EstimatedNumCombinations())

}
