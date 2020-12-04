package dates
const DaysInWeek int=7//상수선언
func WeeksToDays(weeks int) int{
	return weeks*7
}
func DaysToWeeks(days int) float64{
	return float64(days) / float64(7)
}