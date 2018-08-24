# PCSA-TC - Probabilistic Counting with Stochastic Averaging with Tail Cutting

Inspired by ["Better with fewer bits: Improving the performance of cardinality estimation of large data streams - Qingjun Xiao, You Zhou, Shigang Chen"](http://cse.seu.edu.cn/PersonalPage/csqjxiao/csqjxiao_files/papers/INFOCOM17.pdf), I thought of applying the same idea to good old PCSA. Instead of using 32-bit or 64-bit registers I reduced it to 8-bit per register and an 8bit offset counter.
* The initial results are competitive and comparable to HyperLogLog. Throughout this research I will be comparing it with LogLog-Beta.
* Another observation is that the sum of `Tralining Ones` is always overestimated compared to classic PCSA. This is something that can be used.
* One can also start with a full fledged 32bit register version... instead of [k]int32 we can do [32]Bitstream where bitstream is k. Once we know all registers have at least trailing 1 we downsize the bitstream by 1 bit and increment the base... over time the PCSA will be shrinking.

## Run it
```bash
go run demo/main.go
```

## Example Results
Again this needs more testing and plotting and what not...

```
Exact Cardinality: 5000000      PCSA-TailCut (ratio): 0.9744    LogLogBeta (ratio): 0.9781      PCSA-TailCut win: false
Exact Cardinality: 10000000     PCSA-TailCut (ratio): 0.9860    LogLogBeta (ratio): 0.9861      PCSA-TailCut win: false
Exact Cardinality: 15000000     PCSA-TailCut (ratio): 0.9861    LogLogBeta (ratio): 0.9949      PCSA-TailCut win: false
Exact Cardinality: 20000000     PCSA-TailCut (ratio): 0.9943    LogLogBeta (ratio): 0.9911      PCSA-TailCut win: true
Exact Cardinality: 25000000     PCSA-TailCut (ratio): 0.9986    LogLogBeta (ratio): 0.9934      PCSA-TailCut win: true
Exact Cardinality: 30000000     PCSA-TailCut (ratio): 0.9936    LogLogBeta (ratio): 0.9943      PCSA-TailCut win: false
Exact Cardinality: 35000000     PCSA-TailCut (ratio): 1.0029    LogLogBeta (ratio): 1.0008      PCSA-TailCut win: false
Exact Cardinality: 40000000     PCSA-TailCut (ratio): 1.0019    LogLogBeta (ratio): 0.9966      PCSA-TailCut win: true
Exact Cardinality: 45000000     PCSA-TailCut (ratio): 1.0069    LogLogBeta (ratio): 0.9939      PCSA-TailCut win: false
Exact Cardinality: 50000000     PCSA-TailCut (ratio): 1.0013    LogLogBeta (ratio): 0.9942      PCSA-TailCut win: true
Exact Cardinality: 55000000     PCSA-TailCut (ratio): 0.9993    LogLogBeta (ratio): 0.9942      PCSA-TailCut win: true
Exact Cardinality: 60000000     PCSA-TailCut (ratio): 1.0013    LogLogBeta (ratio): 0.9939      PCSA-TailCut win: true
Exact Cardinality: 65000000     PCSA-TailCut (ratio): 0.9993    LogLogBeta (ratio): 0.9858      PCSA-TailCut win: true
Exact Cardinality: 70000000     PCSA-TailCut (ratio): 1.0036    LogLogBeta (ratio): 0.9860      PCSA-TailCut win: true
Exact Cardinality: 75000000     PCSA-TailCut (ratio): 1.0034    LogLogBeta (ratio): 0.9859      PCSA-TailCut win: true
Exact Cardinality: 80000000     PCSA-TailCut (ratio): 1.0029    LogLogBeta (ratio): 0.9856      PCSA-TailCut win: true
Exact Cardinality: 85000000     PCSA-TailCut (ratio): 1.0008    LogLogBeta (ratio): 0.9872      PCSA-TailCut win: true
Exact Cardinality: 90000000     PCSA-TailCut (ratio): 1.0019    LogLogBeta (ratio): 0.9916      PCSA-TailCut win: true
Exact Cardinality: 95000000     PCSA-TailCut (ratio): 1.0051    LogLogBeta (ratio): 0.9959      PCSA-TailCut win: false
wins: 12 loss: 7
```

```
Exact Cardinality: 5000000      PCSA-TailCut (ratio): 0.9719    LogLogBeta (ratio): 1.0063      PCSA-TailCut win: false
Exact Cardinality: 10000000     PCSA-TailCut (ratio): 0.9908    LogLogBeta (ratio): 1.0079      PCSA-TailCut win: false
Exact Cardinality: 15000000     PCSA-TailCut (ratio): 0.9882    LogLogBeta (ratio): 1.0109      PCSA-TailCut win: false
Exact Cardinality: 20000000     PCSA-TailCut (ratio): 0.9837    LogLogBeta (ratio): 1.0047      PCSA-TailCut win: false
Exact Cardinality: 25000000     PCSA-TailCut (ratio): 0.9894    LogLogBeta (ratio): 1.0100      PCSA-TailCut win: false
Exact Cardinality: 30000000     PCSA-TailCut (ratio): 0.9899    LogLogBeta (ratio): 1.0141      PCSA-TailCut win: true
Exact Cardinality: 35000000     PCSA-TailCut (ratio): 0.9883    LogLogBeta (ratio): 1.0121      PCSA-TailCut win: true
Exact Cardinality: 40000000     PCSA-TailCut (ratio): 0.9903    LogLogBeta (ratio): 1.0044      PCSA-TailCut win: false
Exact Cardinality: 45000000     PCSA-TailCut (ratio): 0.9948    LogLogBeta (ratio): 1.0046      PCSA-TailCut win: false
Exact Cardinality: 50000000     PCSA-TailCut (ratio): 0.9952    LogLogBeta (ratio): 1.0040      PCSA-TailCut win: false
Exact Cardinality: 55000000     PCSA-TailCut (ratio): 0.9972    LogLogBeta (ratio): 1.0013      PCSA-TailCut win: false
Exact Cardinality: 60000000     PCSA-TailCut (ratio): 1.0012    LogLogBeta (ratio): 0.9990      PCSA-TailCut win: false
Exact Cardinality: 65000000     PCSA-TailCut (ratio): 1.0020    LogLogBeta (ratio): 0.9999      PCSA-TailCut win: false
Exact Cardinality: 70000000     PCSA-TailCut (ratio): 1.0030    LogLogBeta (ratio): 1.0042      PCSA-TailCut win: true
Exact Cardinality: 75000000     PCSA-TailCut (ratio): 1.0039    LogLogBeta (ratio): 1.0066      PCSA-TailCut win: true
Exact Cardinality: 80000000     PCSA-TailCut (ratio): 1.0055    LogLogBeta (ratio): 1.0061      PCSA-TailCut win: true
Exact Cardinality: 85000000     PCSA-TailCut (ratio): 1.0037    LogLogBeta (ratio): 1.0068      PCSA-TailCut win: true
Exact Cardinality: 90000000     PCSA-TailCut (ratio): 1.0041    LogLogBeta (ratio): 1.0161      PCSA-TailCut win: true
Exact Cardinality: 95000000     PCSA-TailCut (ratio): 1.0038    LogLogBeta (ratio): 1.0188      PCSA-TailCut win: true
wins: 8 loss: 11
```

```
Exact Cardinality: 10000000     PCSA-TailCut (ratio): 0.9928    LogLogBeta (ratio): 0.9910      PCSA-TailCut win: true
Exact Cardinality: 15000000     PCSA-TailCut (ratio): 0.9942    LogLogBeta (ratio): 0.9948      PCSA-TailCut win: false
Exact Cardinality: 20000000     PCSA-TailCut (ratio): 0.9949    LogLogBeta (ratio): 0.9933      PCSA-TailCut win: true
Exact Cardinality: 25000000     PCSA-TailCut (ratio): 0.9886    LogLogBeta (ratio): 0.9898      PCSA-TailCut win: false
Exact Cardinality: 30000000     PCSA-TailCut (ratio): 0.9915    LogLogBeta (ratio): 0.9843      PCSA-TailCut win: true
Exact Cardinality: 35000000     PCSA-TailCut (ratio): 0.9886    LogLogBeta (ratio): 0.9919      PCSA-TailCut win: false
Exact Cardinality: 40000000     PCSA-TailCut (ratio): 0.9928    LogLogBeta (ratio): 0.9940      PCSA-TailCut win: false
Exact Cardinality: 45000000     PCSA-TailCut (ratio): 0.9964    LogLogBeta (ratio): 0.9932      PCSA-TailCut win: true
Exact Cardinality: 50000000     PCSA-TailCut (ratio): 0.9941    LogLogBeta (ratio): 0.9937      PCSA-TailCut win: true
Exact Cardinality: 55000000     PCSA-TailCut (ratio): 0.9911    LogLogBeta (ratio): 0.9923      PCSA-TailCut win: false
Exact Cardinality: 60000000     PCSA-TailCut (ratio): 0.9933    LogLogBeta (ratio): 0.9890      PCSA-TailCut win: true
Exact Cardinality: 65000000     PCSA-TailCut (ratio): 0.9951    LogLogBeta (ratio): 0.9845      PCSA-TailCut win: true
Exact Cardinality: 70000000     PCSA-TailCut (ratio): 0.9976    LogLogBeta (ratio): 0.9847      PCSA-TailCut win: true
Exact Cardinality: 75000000     PCSA-TailCut (ratio): 0.9970    LogLogBeta (ratio): 0.9858      PCSA-TailCut win: true
Exact Cardinality: 80000000     PCSA-TailCut (ratio): 0.9955    LogLogBeta (ratio): 0.9844      PCSA-TailCut win: true
Exact Cardinality: 85000000     PCSA-TailCut (ratio): 0.9967    LogLogBeta (ratio): 0.9844      PCSA-TailCut win: true
Exact Cardinality: 90000000     PCSA-TailCut (ratio): 0.9974    LogLogBeta (ratio): 0.9857      PCSA-TailCut win: true
Exact Cardinality: 95000000     PCSA-TailCut (ratio): 0.9970    LogLogBeta (ratio): 0.9874      PCSA-TailCut win: true
wins: 13 loss: 6
```
