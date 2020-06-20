package collector

func dollars(munits int64) float64 {
	return float64(munits) / 1000
}

func bool2str(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
