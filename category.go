package randomwords

import (
	"math/rand"
)

type Category struct {
	Name  string
	Words []string
}

func (c *Category) RandomWord() string {
	return c.Words[rand.Intn(len(c.Words)-1)]

}
