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
Exact Cardinality: 100000000    PCSA-TailCut (%err): 0.8013     LogLogBeta (%err): -0.6308
Exact Cardinality: 200000000    PCSA-TailCut (%err): 1.5675     LogLogBeta (%err): 0.5411
Exact Cardinality: 300000000    PCSA-TailCut (%err): 0.6733     LogLogBeta (%err): 0.2857
Exact Cardinality: 400000000    PCSA-TailCut (%err): 0.4522     LogLogBeta (%err): 0.3363
Exact Cardinality: 500000000    PCSA-TailCut (%err): 0.0727     LogLogBeta (%err): 0.4330
Exact Cardinality: 600000000    PCSA-TailCut (%err): 0.6009     LogLogBeta (%err): 0.2855
Exact Cardinality: 700000000    PCSA-TailCut (%err): 0.1649     LogLogBeta (%err): 0.7454
Exact Cardinality: 800000000    PCSA-TailCut (%err): 0.3970     LogLogBeta (%err): 0.7733
Exact Cardinality: 900000000    PCSA-TailCut (%err): 0.8352     LogLogBeta (%err): 0.9800
```

```
Exact Cardinality: 100000000    PCSA-TailCut (%err): 0.2060     LogLogBeta (%err): 0.7875
Exact Cardinality: 200000000    PCSA-TailCut (%err): 0.2230     LogLogBeta (%err): -0.1703
Exact Cardinality: 300000000    PCSA-TailCut (%err): -0.5290    LogLogBeta (%err): -0.9651
Exact Cardinality: 400000000    PCSA-TailCut (%err): -0.3014    LogLogBeta (%err): -0.9124
Exact Cardinality: 500000000    PCSA-TailCut (%err): -0.4046    LogLogBeta (%err): -1.5113
Exact Cardinality: 600000000    PCSA-TailCut (%err): 0.0957     LogLogBeta (%err): -1.6920
Exact Cardinality: 700000000    PCSA-TailCut (%err): -0.0975    LogLogBeta (%err): -2.1187
Exact Cardinality: 800000000    PCSA-TailCut (%err): 0.1340     LogLogBeta (%err): -1.6607
Exact Cardinality: 900000000    PCSA-TailCut (%err): 0.0237     LogLogBeta (%err): -1.6979
```