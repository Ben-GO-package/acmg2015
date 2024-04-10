# acmg2015
acmg2015 pred

# 需求字段及各字段影响证据项目

| 解析字段 | BA1 | BP1 | BP3 | BP4 | BP7 | BS1 | BS2 | PM1 | PM2 | PM4 | PM5 | PP2 | PP3 | PS1 | PS4 | PVS1 |
| -------- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | ---- |
| Function | /   | √   | √   | √   | √   | /   | /   | √   | /   | √   | √   | √   | √   | √   | /   | √    |



# changelog
## 2024.4.9
适配vep注释软件对acmg中的逻辑进行了更新,主要适配vep注释的Consequence结果作为功能字段输入
#### 功能字段的具体调整内容
1. 所有功能识别，均改为了基于正则进行匹配，所以map格式的匹配统一取消
2. 匹配 `cds-del、cds-ins、cds-indel` 三个字段的位置统一增加了 `inframe_deletion 和inframe_insertion`
3. 匹配 `splice-3,splice-5,init-loss,alt-start,frameshift,nonsense,stop-gain,span` 的位置，增加 `splice_acceptor_variant、splice_donor_variant、start_lost`
4. 匹配`coding-synon` 的位置，增加`synonymous`
5. 上游结果中的`splice±20`标记删除，改为Function中查找splice，通过cHGNVS 判断是否±20标记（±10~20范围）
6. 匹配 `stop-loss`的字段同步增加`stop_lost`

#### 其他调整
数据库信息、repeattag等字段，原本识别的空字符串为 "" 和 "." ，新版本再原有基础上增加了 "-"。

#### 由于注释调整以后，转录本更换，需要同步更换涉及转录本的相关数据库信息
在这里记录一下前期比较发现的一些非代码结构导致的判定证据项不一致的问题，后续更新需要注意迭代相关数据库配置，确保数据库配置和注释结果匹配。
######  BA1 不一致部分
BA1.exception 库文件更新

######  PM2 不一致部分
ModeInheritance 字段结果不一致。

######  PM5 不一致部分
相关数据库需要更新为新转录本对应库
- PS1PM5.AApos.count
```
NM_000014.4:p.C972      1
NM_000014.4:p.R1297     1
NM_000016.4:p.A157      2
NM_000016.4:p.A165      1
NM_000016.4:p.A170      1
NM_000016.4:p.A176      2
NM_000016.4:p.A205      1
NM_000016.4:p.A218      1
NM_000016.4:p.A303      1
NM_000016.4:p.A318      1
```

- PS1PM5.MutationName.count
```
NM_000014.4(A2M): c.2915G>A (p.Cys972Tyr)       1
NM_000014.4(A2M): c.3889C>T (p.Arg1297Cys)      1
NM_000016.4(ACADM): c.1001G>A (p.Arg334Lys)     1
NM_000016.4(ACADM): c.1007G>A (p.Ser336Asn)     1
NM_000016.4(ACADM): c.1008T>A (p.Ser336Arg)     1
NM_000016.4(ACADM): c.1010A>C (p.Tyr337Ser)     1
NM_000016.4(ACADM): c.1012C>G (p.Gln338Glu)     1
NM_000016.4(ACADM): c.1019C>T (p.Ala340Val)     1
NM_000016.4(ACADM): c.1022C>T (p.Ala341Val)     1
NM_000016.4(ACADM): c.1033G>T (p.Asp345Tyr)     1
```

- PS1PM5.pHGVS1.count
```
NM_000014.4:p.C972Y     1
NM_000014.4:p.R1297C    1
NM_000016.4:p.A157T     1
NM_000016.4:p.A157V     1
NM_000016.4:p.A165T     1
NM_000016.4:p.A170S     1
NM_000016.4:p.A176S     1
NM_000016.4:p.A176T     1
NM_000016.4:p.A205P     1
NM_000016.4:p.A218G     1
```

###### BP4、PP3不一致
dbnsfp 版本差异导致的 `SIFT Pred``Polyphen2 HVAR Pred``MutationTaster Pred`

###### PVS1 不一致
旧版本中部分突变缺失autoPVS结果