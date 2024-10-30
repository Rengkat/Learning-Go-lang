package calcfunc

func Calculate(amount float64, year int, rate float64 ) float64  {
	answer := (amount * rate * float64(year))/1000
	return answer
}