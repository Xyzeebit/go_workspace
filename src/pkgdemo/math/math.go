package math

// returns the average of a slice of float64 numbers
//
//     x := []float64{1,2,3}
//
//     y := math.Average(x)
//
//     fmt.Println(y) //2
func Average(numbers []float64) float64 {
	total := float64(0);
	for _, i := range numbers {
		total += i;
	}
	return total / float64(len(numbers));
}

// Finds the minimum number in a slice of numbers
//
//     x := []float64{1, 3, 2}
//
//     y := math.Min(x)
//
//     fmt.Println(y) // 1
//
func Min(numbers []float64) float64 {
	var min float64 = numbers[0];
	for _, i := range numbers {
		if i < min {
			min = i;
		}
	}
	return min
}

// Finds the maximum number in a slice of numbers
//
//     x := []float64{1, 3, 2}
//
//     y := math.Max(x)
//
//     fmt.Println(y) // 3
//
func Max(numbers []float64) float64 {
	var max float64 = numbers[0];
	for _, i := range numbers {
		if max < i {
			max = i;
		}
	}
	return max;
}