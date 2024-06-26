package acmg2015

import (
	"github.com/Ben-GO-package/acmg2015/evidence"
	"github.com/brentp/bix"
	"github.com/liserjrqlxue/goUtil/jsonUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
)

var (
	tbx            *bix.Bix
	LOFGeneList    map[string]int
	transcriptInfo map[string][]evidence.Region
)

func Init(cfg map[string]string, AutoPVS1 bool, runPM1 bool) {
	evidence.LoadPS1PM5(cfg["PS1PM5.MutationName.count"], cfg["PS1PM5.pHGVS1.count"], cfg["PS1PM5.AApos.count"])
	if runPM1 {
		evidence.LoadPM1(cfg["PM1InterproDomain"], cfg["PM1PfamIdDomain"])
	}
	if !AutoPVS1 {
		LOFGeneList = evidence.LoadLOF(cfg["LOFList"])
		jsonUtil.JsonFile2Data(cfg["transcriptInfo"], &transcriptInfo)
	}
	tbx = simpleUtil.HandleError(bix.New(cfg["PathogenicLite"])).(*bix.Bix)
	evidence.LoadPP2(cfg["PP2GeneList"])
	evidence.LoadBS2(cfg["LateOnset"])
	evidence.LoadBP1(cfg["BP1GeneList"])
	evidence.LoadBA1(cfg["BA1ExceptionList"])
}

func AddEvidences(item map[string]string, AutoPVS1 bool, runPM1 bool) {
	if !AutoPVS1 {
		item["PVS1"] = evidence.CheckPVS1(item, LOFGeneList, transcriptInfo, tbx)
	}
	item["PS1"] = evidence.CheckPS1(item)
	item["PM5"] = evidence.CheckPM5(item)
	item["PS4"] = evidence.CheckPS4(item)
	if runPM1 {
		item["PM1"] = evidence.CheckPM1(item, tbx)
	}
	item["PM2"] = evidence.CheckPM2(item)
	item["PM4"] = evidence.CheckPM4(item)
	item["PP2"] = evidence.CheckPP2(item)
	item["PP3"] = evidence.CheckPP3(item, AutoPVS1)
	item["BA1"] = evidence.CheckBA1(item)
	item["BS1"] = evidence.CheckBS1(item)
	item["BS2"] = evidence.CheckBS2(item)
	item["BP1"] = evidence.CheckBP1(item)
	item["BP3"] = evidence.CheckBP3(item)
	item["BP4"] = evidence.CheckBP4(item)
	item["BP7"] = evidence.CheckBP7(item)
}
