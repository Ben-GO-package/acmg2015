package evidence

import (
	"regexp"
)

var (
	isSplice      = regexp.MustCompile(`splice`)
	isSplice20    = regexp.MustCompile(`splice[+-]20`)
	isD           = regexp.MustCompile(`D`)
	isDeleterious = regexp.MustCompile(`deleterious`)
)

// ture	:	"1"
// flase:	"0"
// nil	:	""
func CheckPP3(item map[string]string) string {
	if item["GERP++_RS_pred"] == "保守" &&
		item["PhyloP Vertebrates Pred"] == "保守" &&
		item["PhyloP Placental Mammals Pred"] == "保守" {
	} else {
		return "0"
	}
	if isSplice.MatchString(item["Function"]) && !isSplice20.MatchString(item["Function"]) {
		if item["PVS1"] == "1" || item["PVS1"] == "5" {
			return "0"
		} else {
			if isD.MatchString(item["dbscSNV_RF_pred"]) &&
				isD.MatchString(item["dbscSNV_ADA_pred"]) {
				return "1"
			} else {
				return "0"
			}
		}
	} else {
		if isD.MatchString(item["SIFT Pred"]) &&
			isD.MatchString(item["Polyphen2 HVAR Pred"]) &&
			isD.MatchString(item["MutationTaster Pred"]) &&
			isDeleterious.MatchString(item["Ens Condel Pred"]) {
			return "1"
		} else {
			return "0"
		}
	}
	return ""
}

func ComparePP3(item map[string]string, lostOnly bool) {
	rule := "PP3"
	val := CheckPP3(item)
	if val != item[rule] {
		if item[rule] == "0" && val == "" {
		} else {
			if lostOnly && val != "1" {
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
				"PVS1",
				"dbscSNV_RF_pred",
				"dbscSNV_ADA_pred",
				"SIFT Pred",
				"Polyphen2 HVAR Pred",
				"MutationTaster Pred",
				"Ens Condel Pred",
			)
		}
	}
}
