package evidence

// PP2
func CheckPP2(item map[string]string) string {
	if !ismissense.MatchString(item["Function"]) {
		return "0"
	}
	if pp2GeneList[item["entrez_id"]] {
		return "1"
	} else {
		return "0"
	}
}

func ComparePP2(item map[string]string) {
	rule := "PP2"
	val := CheckPP2(item)
	if val != item[rule] {
		PrintConflict(item, rule, val, "Function", "entrez_id")
	}
}
