package evidence

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/brentp/bix"
	"github.com/brentp/irelate/interfaces"
	"github.com/liserjrqlxue/goUtil/fmtUtil"
	"github.com/liserjrqlxue/goUtil/osUtil"
	"github.com/liserjrqlxue/goUtil/simpleUtil"
	"github.com/liserjrqlxue/goUtil/stringsUtil"
	"github.com/liserjrqlxue/goUtil/textUtil"
)

type Region struct {
	Seqid string
	//Source string
	Type       string
	Chromosome string
	Start      uint64
	End        uint64
	Strand     string
	Gene       string
}

var (
	PM1Function    = regexp.MustCompile(`missense|cds-indel`)
	isBP3Func      = regexp.MustCompile(`cds-del|cds-ins|cds-indel|inframe_deletion|inframe_insertion|protein_altering_variant`)
	isPM4Func      = regexp.MustCompile(`cds-del|cds-ins|cds-indel|stop-loss|inframe_deletion|inframe_insertion|stop_lost|protein_altering_variant`)
	getAAPos       = regexp.MustCompile(`^p\.[A-Z]\d+`)
	IsClinVarPLP   = regexp.MustCompile(`Pathogenic|Likely_pathogenic`)
	IsHgmdDM       = regexp.MustCompile(`DM$|DM\|`)
	isARDRXLPRDDNA = regexp.MustCompile(`AR|DR|XL|PR|DD|NA|UNK`)
	isADPDYL       = regexp.MustCompile(`AD|PD|YL`)
	isSplice       = regexp.MustCompile(`splice`)
	ismissense     = regexp.MustCompile(`missense`)
	isSpliceAccDon = regexp.MustCompile(`splice_acceptor_variant|splice_donor_variant`)
	isSpliceIntron = regexp.MustCompile(`splice|intron`)
	isIntron       = regexp.MustCompile(`intron`)
	isSplice20     = regexp.MustCompile(`splice[+-]20`)
	isP            = regexp.MustCompile(`P`)
	isD            = regexp.MustCompile(`D`)
	isI            = regexp.MustCompile(`I`)
	isNeutral      = regexp.MustCompile(`neutral`)
	isDeleterious  = regexp.MustCompile(`deleterious`)
	repeatSeq      = regexp.MustCompile(`c\..*\[(\d+)>\d+]`)
)

var chgvsReg = regexp.MustCompile(`c\.\d+([+-])(\d+)`)

func get_abs(input string) int {
	if input == "-" {
		return 0
	}
	input_int, _ := strconv.Atoi(input)
	if input_int < 0 {
		return -input_int
	} else {
		return input_int
	}
}

// Tier1 >1
// LoF 3
var FuncInfo = map[string]int{
	"splice-3":     3,
	"splice-5":     3,
	"init-loss":    3,
	"alt-start":    3,
	"frameshift":   3,
	"nonsense":     3,
	"stop-gain":    3,
	"span":         3,
	"missense":     2,
	"cds-del":      2,
	"cds-indel":    2,
	"cds-ins":      2,
	"splice-10":    2,
	"splice+10":    2,
	"coding-synon": 1,
	"splice-20":    1,
	"splice+20":    1,
}

func checkFuncInfo(Function string) int {
	var isFuncInfo = regexp.MustCompile(`splice-3|splice_acceptor_variant|splice-5|splice_donor_variant|init-loss|start_lost|alt-start|frameshift|nonsense|stop-gain|span`)
	if isFuncInfo.MatchString(Function) {
		return 3
	} else {
		return 1
	}
}
func CheckAFAllLowThen(item map[string]string, AFlist []string, threshold float64, includeEqual bool) bool {
	for _, key := range AFlist {
		af := item[key]
		if af == "" || af == "." || af == "0" || af == "-" {
			continue
		}
		AF, err := strconv.ParseFloat(af, 64)
		simpleUtil.CheckErr(err)
		if includeEqual {
			if AF > threshold {
				return false
			}
		} else {
			if AF >= threshold {
				return false
			}
		}
	}
	return true
}

func PrintConflict(item map[string]string, rule, val string, keys ...string) {
	fmtUtil.Fprintf(
		os.Stderr,
		"Conflict %s:[%s] vs [%s]\t%s[%s]\n",
		rule,
		val,
		item[rule],
		"MutationName",
		item["MutationName"],
	)
	for _, key := range keys {
		fmtUtil.Fprintf(os.Stderr, "\t%30s:[%s]\n", key, item[key])
	}
}

var (
	hgvsCount         = make(map[string]int)
	phgvsCount        = make(map[string]int)
	aaPostCount       = make(map[string]int)
	pm1PfamId         = make(map[string]bool)
	pm1InterproDomain = make(map[string]bool)
	bp1GeneList       = make(map[string]bool)
	bs2GeneList       = make(map[string]bool)
	ba1Exception      = make(map[string]bool)
	pp2GeneList       = make(map[string]bool)
	LOFGeneList       = make(map[string]int)
)

func LoadPS1PM5(hgvs, pHgvs, aaPos string) {
	hgvsCount = tsv2mapStringInt(hgvs)
	phgvsCount = tsv2mapStringInt(pHgvs)
	aaPostCount = tsv2mapStringInt(aaPos)
}

func tsv2mapStringInt(tsv string) map[string]int {
	var db = make(map[string]int)

	var file = osUtil.Open(tsv)
	defer simpleUtil.DeferClose(file)

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		var array = strings.Split(scanner.Text(), "\t")
		array = append(array, "NA", "NA")
		if array[1] == "NA" {
			array[1] = "0"
		}
		var v, ok = db[array[0]]
		if ok {
			var vStr = strconv.Itoa(v)
			if array[1] != vStr {
				panic("dup key[" + array[0] + "],different value:[" + vStr + "]vs[" + array[1] + "]")
			}
		} else {
			db[array[0]] = stringsUtil.Atoi(array[1])
		}
	}
	simpleUtil.CheckErr(scanner.Err())
	return db
}

func LoadPM1(pfamId, interproDomain string) {
	var pfamIdArray = textUtil.File2Array(pfamId)
	var interproDomainArray = textUtil.File2Array(interproDomain)
	for _, key := range pfamIdArray {
		pm1PfamId[key] = true
	}
	for _, key := range interproDomainArray {
		pm1InterproDomain[key] = true
	}
}

func countBix(tbx *bix.Bix, chr string, start, end int) (n int) {
	rdr, err := tbx.Query(interfaces.AsIPosition(chr, start, end))
	simpleUtil.CheckErr(err)
	defer simpleUtil.DeferClose(rdr)
	for {
		_, err := rdr.Next()
		if err == io.EOF {
			break
		}
		simpleUtil.CheckErr(err)
		n++
	}
	return n
}

func LoadLOF(LOFGeneListFile string) map[string]int {
	for _, line := range textUtil.File2Array(LOFGeneListFile) {
		array := strings.Split(line, "\t")
		LOFGeneList[array[1]] = stringsUtil.Atoi(array[1])
	}
	return LOFGeneList
}

func LoadBP1(bp1geneList string) {
	var genes = textUtil.File2Array(bp1geneList)
	bp1GeneList = make(map[string]bool)
	for _, line := range genes {
		array := strings.Split(line, "\t")
		bp1GeneList[array[1]] = true
	}
}

func LoadBS2(fileName string) {
	for _, line := range textUtil.File2Array(fileName) {
		array := strings.Split(line, "\t")
		bs2GeneList[array[1]] = true
	}
}

func LoadPP2(fileName string) {
	for _, line := range textUtil.File2Array(fileName) {
		array := strings.Split(line, "\t")
		pp2GeneList[array[1]] = true
	}
}

func LoadBA1(fileName string) {
	for _, key := range textUtil.File2Array(fileName) {
		ba1Exception[key] = true
	}
}
