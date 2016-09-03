package strcat

import "expvar"

func init() {
	expvar.NewInt("count")
}
