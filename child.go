package child

import (
	v3 "github.com/Stebalien/go-mod-play-dep/v3"
)

var Thing = "2.1.0"

func init() {
	v3.Thing += Thing
}
