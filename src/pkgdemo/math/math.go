package math

func Average(xs []float64) float64 {
	total := float64(0);
	for _, i := range xs {
		total += i;
	}
	return total / float64(len(xs));
}