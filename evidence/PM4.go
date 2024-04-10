package evidence

// ture	:	"1"
// flase:	"0"
// nil	:	""
func CheckPM4(item map[string]string) string {
	if isPM4Func.MatchString(item["Function"]) {
		if item["RepeatTag"] == "" || item["RepeatTag"] == "." || item["RepeatTag"] == "-" {
			return "1"
		} else {
			return "0"
		}
	} else {
		return "0"
	}
}

func ComparePM4(item map[string]string) {
	rule := "PM4"
	val := CheckPM4(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, "Function", "RepeatTag")
	}
}
