package randomwords

type Generator struct {
	Categories            map[string]*Category
	Words                 []string
	DefaultCategoryLayout []string // EG: {"adj","adj","noun"}
	AllowDupes            bool
}

func (g *Generator) Random() string {
	return g.RandomFromCategories(g.DefaultCategoryLayout...)
}

func (g *Generator) RandomFromCategories(categories ...string) string {
	words := make(map[string]bool, len(categories))
	result := ""

	for _, name := range categories {
		//fmt.Println("name", name)
		cat := g.Categories[name]
		word := ""
		var ok bool
		for ok || word == "" {
			word = cat.RandomWord()
			_, ok = words[word]
		}
		words[word] = true
		result += word
	}
	return result
}

// This is only an estimate because it doesn't take into account the AllowDupes setting at all
func (g *Generator) EstimatedNumCombinationsFromLayout(categories ...string) int {
	num := 0
	for _, v := range categories {
		cat := g.Categories[v]
		if num == 0 {
			num = len(cat.Words)
		} else {
			num = num * len(cat.Words)
		}
	}
	return num
}

func (g *Generator) EstimatedNumCombinations() int {
	return g.EstimatedNumCombinationsFromLayout(g.DefaultCategoryLayout...)
}

func CombineCategories(name string, cats ...*Category) *Category {
	wordCheck := make(map[string]bool)
	for _, cat := range cats {
		for _, word := range cat.Words {
			if _, ok := wordCheck[word]; !ok {
				wordCheck[word] = true
			}
		}
	}

	words := make([]string, len(wordCheck))
	i := 0
	for word, _ := range wordCheck {
		words[i] = word
		i++
	}

	return &Category{
		Name:  name,
		Words: words,
	}
}

func DefaultGenerator() *Generator {
	return &Generator{
		Categories: map[string]*Category{
			"all": CombineCategories("all",
				Armor,
				Colors,
				CosmicThings,
				Defenders,
				Fortresses,
				GoodQualities,
				PrettyAdjs,
				RPGClasses,
				Sizes,
				Stories,
				Titles,
			),
		},
		DefaultCategoryLayout: []string{"all", "all", "all"},
	}
}

func AdjNounGenerator() *Generator {
	var adjs = CombineCategories("adj", Colors, Sizes, GoodQualities, PrettyAdjs, Moods)
	var nouns = CombineCategories("nouns", Titles, CosmicThings, Fortresses, Defenders, Armor)

	return &Generator{
		Categories: map[string]*Category{
			adjs.Name:  adjs,
			nouns.Name: nouns,
		},
		DefaultCategoryLayout: []string{"adj", "adj", "nouns"},
	}
}
