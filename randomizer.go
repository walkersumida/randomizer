package randomizer

import (
	"math/rand"
	"strings"

	"github.com/walkersumida/randomizer/seed"
)

func Run(opt string) string {
	opts := strings.Split(opt, ",")
	rand.Seed(seed.Int64())

	selectedOpt := opts[rand.Intn(len(opts))]

	return selectedOpt
}
