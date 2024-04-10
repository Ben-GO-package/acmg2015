package evidence

// ture	:	"1"
// flase:	"0"
func CheckBP4(item map[string]string) string {
	if !(item["GERP++_RS_pred"] == "不保守" &&
		item["PhyloP Vertebrates Pred"] == "不保守" &&
		item["PhyloP Placental Mammals Pred"] == "不保守") {
		return "0"
	}
	if isSpliceIntron.MatchString(item["Function"]) {
		if isSplice.MatchString(item["Function"]) && get_distance(item["cHGVS"]) == 20 {
			if isD.MatchString(item["dbscSNV_RF_pred"]) ||
				isD.MatchString(item["dbscSNV_ADA_pred"]) ||
				isD.MatchString(item["SpliceAI Pred"]) {
				return "0"
			}
			if isP.MatchString(item["SpliceAI Pred"]) {
				return "1"
			}
			return "0"
		}
	} else {
		if isP.MatchString(item["SIFT Pred"]) &&
			isP.MatchString(item["Polyphen2 HVAR Pred"]) &&
			isP.MatchString(item["MutationTaster Pred"]) &&
			isNeutral.MatchString(item["Ens Condel Pred"]) {
			return "1"
		} else {
			return "0"
		}
	}
	return "0"
}

func CompareBP4(item map[string]string, lostOnly bool) {
	rule := "BP4"
	val := CheckBP4(item)
	if val != item[rule] {
		if lostOnly && val == "0" {
			return
		}
		PrintConflict(
			item,
			rule,
			val,
			"GERP++_RS_pred",
			"PhyloP Vertebrates Pred",
			"PhyloP Placental Mammals Pred",
			"Function",
			"dbscSNV_RF_pred",
			"dbscSNV_ADA_pred",
			"SIFT Pred",
			"Polyphen2 HVAR Pred",
			"MutationTaster Pred",
			"Ens Condel Pred",
			"SpliceAI Pred",
		)
	}
}
