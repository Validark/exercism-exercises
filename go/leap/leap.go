package leap

func main() {

}
func IsLeapYear(year int) bool {
	switch {
	case year%4 == 0 && year%100 != 0 || year%400 == 0:
		return true
	default:
		return false
	}
}

/**
```plain
on every year that is evenly divisible by 4  (y%4 == 0)
  except every year that is evenly divisible by 100 (y%100 != 0)
    unless the year is also evenly divisible by 400 (y%)
```
**/