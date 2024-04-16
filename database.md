# 需求数据库
本模块（acmg2015）使用需要导入数据库通过 acmg2015.Init(cfg map[string]string) 完成初始化
其中cfg是一个map文件，配置文件示例如下:

```shell
#文件关键字 文件目录
PS1PM5.MutationName.count	acmg/PS1PM5.MutationName.count
PS1PM5.pHGVS1.count	acmg/PS1PM5.pHGVS1.count
PS1PM5.AApos.count	acmg/PS1PM5.AApos.count
PM1InterproDomain	acmg/PM1.InterproDomain.count
PM1PfamIdDomain	acmg/PM1.pfamId.count
LOFList	acmg/PVS1GeneList.json
transcriptInfo	acmg/transcript.info.json
PathogenicLite	acmg/PathogenicLite.bed.gz
PP2GeneList	acmg/PP2.gene.list
LateOnset	acmg/BS2.gene.list
BP1GeneList	acmg/BP1.gene.list
BA1ExceptionList	acmg/BA1.exception.list
```

## PS1PM5.MutationName.count 
ClinVar或HGMD中收录为pathogenic、likely pathogenic、Uncertain_significance、 DM、pathogenic/likely pathogenic、Conflicting_interpretations_of_pathogenicity（不包含 B、LB）的missense变异
```
ATP6: c.14T>C (p.Leu5Pro)    1
```

## PS1PM5.pHGVS1.count
PS1PM5.MutationName.count 仅包含含计算变化，不记录变异位点（cHGVS）
```
ATP6:p.L5P	1
```

## PS1PM5.AApos.count
PS1PM5.pHGVS1.count 中的含计算变化位置，只记录原氨基酸和位置，不记录突变后的氨基酸
```
ATP6:p.L5	1
```

## PM1InterproDomain  &  PM1PfamIdDomain
将HGMD和ClinVar数据库的missense和仅导致相同个数氨基酸替换的cds-indel变异通过dbNSFP和pfam蛋白预测区段注释得到对应的蛋白结构域信息

## LOFList
基因功能缺失（LOF）是致病机制的基因list清单。
数据库配置未见格式如下，流程使用文件的**第二列** `entrez_id` 进行处理
```
HGNC	entrez_id
AAGAB	79719
ABCA7	10347
ADAM9	8754
```

## transcriptInfo
记录了转录本的染色体起始终止，和正负链信息，用来确认PathogenicLite记录的相关变异和目标变异的关系
```JSON
{"NM_000015.2": [
    {"Seqid": "",
    "Type": "",
    "Chromosome": "8",
    "Start": 18241254,
    "End": 18258723,
    "Strand": "+",
    "Gene": "NAT2"}
    ]
}
```

## PathogenicLite
前面是染色体的位置和ref/alt 信息，后面分别为HGMD的致病性和clinVar的致病性。
ClinVar收录为pathogenic、likely pathogenic、pathogenic/likely pathogenic但HGMD非DP、FP、DFP的变异，或Clinvar无致病性收录但HGMD收录为DM的变异
```
1	877522	877523	C	G	DM?	Uncertain_significance
```

## PP2GeneList
基因集A ： 基因在GnomAD数据库中的mis_z>3.09 的基因。
基因集B ：根据HGMD/ClinVar统计：
```
a. 将ClinVar和HGMD的所有变异合并后去重。
b. 对于某一个基因，统计致病变异数和良性变异数。
    说明1：致病变异：指ClinVar收录为pathogenic、likely pathogenic、pathogenic/likely pathogenic但HGMD非DP、FP、DFP的变异，或Clinvar无致病性收录但HGMD收录为DM的变异；
    说明2：良性变异：指ClinVar收录为benign、likely benign、likely benign但HGMD非DM的变异，或ClinVar无致病性收录但HGMD收录为DP、FP、DFP的变异。
c. 若该基因的致病变异集（致病变异数≥10）中≥80%是missense；且该基因的良性变异集（不做数目要求）中≤10%的是missense，则生成一个基因集。
```
将基因集A和基因集B合并（并集）得到最终的PP2基因集("PP2GeneList")。
数据库配置未见格式如下，流程使用文件的**第二列** `entrez_id` 进行处理
```
HGNC	entrez_id
AAGAB	79719
ABCA7	10347
ADAM9	8754
```
## LateOnset
参考OMIM等总结外显不全或晚发疾病集，确定相关基因集。
```
A2M
AAGAB
AARS
```
数据库配置未见格式如下，流程使用文件的**第二列** `entrez_id` 进行处理
```
HGNC	entrez_id
AAGAB	79719
ABCA7	10347
ADAM9	8754
```
## BP1GeneList
1. 将ClinVar和HGMD的所有变异合并后去重；
2. 对于某一个基因，统计致病变异数；
`说明1：致病变异：指ClinVar收录为pathogenic、likely pathogenic、pathogenic/likely pathogenic但HGMD无DP、FP、DFP的变异，或Clinvar无致病性收录但HGMD收录为DM的变异；`
3. 若该基因的致病变异集（致病变异数≥10）中≥80%是LOF（splice-3, splice-5, init-loss, alt-start, frameshift, nonsense, stop-gain，span）；则生成一个基因集 `BP1GeneList`。
数据库配置未见格式如下，流程使用文件的**第二列** `entrez_id` 进行处理
```
HGNC	entrez_id
AAGAB	79719
ABCA7	10347
ADAM9	8754
```
## BA1ExceptionList
不适用BA1规则的突变列表。
```
NM_000410.3:c.187C>G
NM_014049.4:c.-44_-41dupTAAG
```


# 注意事项
- 部分库文件解析是基于转录本，需要确认建库所用的转录本和数据库记录的转录本一致，比如 `PS1PM5相关的三个数据库` 和 `BA1.exception`
