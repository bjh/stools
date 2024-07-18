package cnvrt

import "strconv"

func ToI(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func ToF(s string) float32 {
	f, _ := strconv.ParseFloat(s, 32)
	return float32(f)
}
