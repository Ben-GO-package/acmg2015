## v.1.1.21
适配vep注释软件对acmg中的逻辑进行了更新,主要适配vep注释的Consequence结果作为功能字段输入

#### Function兼容vep注释consequence结果作为功能字段输入
1. 所有功能识别，均改为了基于正则进行匹配，所以map格式的匹配统一取消
2. 匹配 `cds-del、cds-ins、cds-indel` 三个字段的位置统一增加了 `inframe_deletion 和inframe_insertion、 protein_altering_variant`
3. 匹配 `splice-3,splice-5,init-loss,alt-start,frameshift,nonsense,stop-gain,span` 的位置，增加 `splice_acceptor_variant、splice_donor_variant、start_lost`
4. 匹配`coding-synon` 的位置，增加`synonymous`
5. 上游结果中的`splice±20`标记删除，改为Function中查找splice，通过cHGNVS 判断是否±20标记（±10~20范围）
6. 匹配 `stop-loss`的字段同步增加`stop_lost`
#### 其他调整
1. 数据库信息的 `repeattag` 等字段，原本识别的空字符串为 "" 和 "." ，新版本再原有基础上增加了 "-"。
2. 变异距离外显子的距离，从使用 `cHGVS` 进行正则匹配获取，改为使用vep插件计算的 `HGVS_IntronStartOffset` 和 `HGVS_IntronEndOffset`

## v1.2.1
1. 4个基因list库（`LOFList、PP2GeneList、LateOnset、BP1GeneList`)匹配从使用 `symble` 转成 `entrez_id` 。对应数据库格式调整为如下格式：
首列为基因名称，第二列为 `entrez_id` 。
```
HGNC	entrez_id
AAGAB	79719
ABCA7	10347
ADAM9	8754
```
2. 默认不在进行`PM1`证据项的处理。

