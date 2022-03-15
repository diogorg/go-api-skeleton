package cast

import (
	"api/support"
	"strconv"
)

func StringToInt(v string) int {
	changed, err := strconv.Atoi(v)
	support.Panic(err)

	return changed
}
