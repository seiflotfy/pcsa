# PCSA-TC - Probabilistic Counting with Stochastic Averaging with Tail Cutting

Inspired by ["Better with fewer bits: Improving the performance of cardinality estimation of large data streams - Qingjun Xiao, You Zhou, Shigang Chen"](http://cse.seu.edu.cn/PersonalPage/csqjxiao/csqjxiao_files/papers/INFOCOM17.pdf), I thought of applying the same idea to good old PCSA. Instead of using 32-bit or 64-bit registers I reduced it to 8-bit per register and an 8bit offset counter.
* The initial results are competitive and comparable to HyperLogLog. Throught this research I will be comparing it with LogLog-Beta.
* Another observation is that the sum of `Tralining Ones` is always overestimated compared to classic PCSA

## Run it
```bash
go run demo/main.go
```

## Example Results
Again this needs more testing and plotting and what not...
```
Exact Cardinality: 10000000     PCSA-TailCut (ratio): 0.9923    LogLogBeta (ratio): 0.9964
Exact Cardinality: 20000000     PCSA-TailCut (ratio): 0.9992    LogLogBeta (ratio): 0.9982
Exact Cardinality: 30000000     PCSA-TailCut (ratio): 0.9946    LogLogBeta (ratio): 1.0039
Exact Cardinality: 40000000     PCSA-TailCut (ratio): 1.0018    LogLogBeta (ratio): 0.9983
Exact Cardinality: 50000000     PCSA-TailCut (ratio): 1.0080    LogLogBeta (ratio): 1.0024
Exact Cardinality: 60000000     PCSA-TailCut (ratio): 1.0064    LogLogBeta (ratio): 0.9999
Exact Cardinality: 70000000     PCSA-TailCut (ratio): 1.0063    LogLogBeta (ratio): 0.9978
Exact Cardinality: 80000000     PCSA-TailCut (ratio): 1.0118    LogLogBeta (ratio): 1.0015
Exact Cardinality: 90000000     PCSA-TailCut (ratio): 1.0092    LogLogBeta (ratio): 1.0021
```

```
Exact Cardinality: 10000000     PCSA-TailCut (ratio): 0.9875    LogLogBeta (ratio): 0.9942
Exact Cardinality: 20000000     PCSA-TailCut (ratio): 0.9951    LogLogBeta (ratio): 1.0073
Exact Cardinality: 30000000     PCSA-TailCut (ratio): 0.9965    LogLogBeta (ratio): 1.0128
Exact Cardinality: 40000000     PCSA-TailCut (ratio): 0.9937    LogLogBeta (ratio): 1.0140
Exact Cardinality: 50000000     PCSA-TailCut (ratio): 0.9982    LogLogBeta (ratio): 1.0046
Exact Cardinality: 60000000     PCSA-TailCut (ratio): 0.9990    LogLogBeta (ratio): 1.0107
Exact Cardinality: 70000000     PCSA-TailCut (ratio): 0.9974    LogLogBeta (ratio): 1.0084
Exact Cardinality: 80000000     PCSA-TailCut (ratio): 0.9975    LogLogBeta (ratio): 1.0033
Exact Cardinality: 90000000     PCSA-TailCut (ratio): 0.9991    LogLogBeta (ratio): 1.0041
```

```
Exact Cardinality: 10000000     PCSA-TailCut (ratio): 1.0024    LogLogBeta (ratio): 1.0083
Exact Cardinality: 20000000     PCSA-TailCut (ratio): 1.0102    LogLogBeta (ratio): 0.9918
Exact Cardinality: 30000000     PCSA-TailCut (ratio): 1.0117    LogLogBeta (ratio): 0.9898
Exact Cardinality: 40000000     PCSA-TailCut (ratio): 1.0093    LogLogBeta (ratio): 0.9925
Exact Cardinality: 50000000     PCSA-TailCut (ratio): 1.0050    LogLogBeta (ratio): 0.9925
Exact Cardinality: 60000000     PCSA-TailCut (ratio): 1.0023    LogLogBeta (ratio): 1.0012
Exact Cardinality: 70000000     PCSA-TailCut (ratio): 1.0035    LogLogBeta (ratio): 0.9968
Exact Cardinality: 80000000     PCSA-TailCut (ratio): 1.0043    LogLogBeta (ratio): 1.0021
Exact Cardinality: 90000000     PCSA-TailCut (ratio): 1.0017    LogLogBeta (ratio): 1.0057
```