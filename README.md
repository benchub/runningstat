# RunningStat
RunningStat is a bit of code that lets you keep track of the mean, standard deviation, and varience for a series of numbers. Push more samples, and the metrics are recomputed as you go. 

That's neat and all, but what's maybe less obvious (and certainly more handy) is that you can save your current stats to disk/database, reload them at a later date, and keep using the running tally.

You can also merge running stat variables.

```
rs := RunningStat{}

rs.Push(1)
rs.Push(2)
rs.RunningStatCount() // 2
rs.RunningStatMean() // 1.5
rs.RunningStatDeviation() // 0.70711
rs.RunningStatVariation() // 0.5

rs.Push(3)
rs.RunningStatCount() // 3
rs.RunningStatMean() // 2
rs.RunningStatDeviation() // 1
rs.RunningStatVariation() // 1

rs.Reset()
rs.Init(3,2,1)
rs.Push(4)
rs.RunningStatCount() // 4
rs.RunningStatMean() // 2.5
rs.RunningStatDeviation() // 1.29099
rs.RunningStatVariation() // 1.66667

// You can also merge things
rs.Reset()
rs.Push(5)
rs.Push(5)
rs.Push(5)
rs.RunningStatCount() // 3
rs.RunningStatMean() // 5
rs.RunningStatDeviation() // 0
rs.RunningStatVariation() // 0

rs2 := RunningStat{}
rs2.Push(1)
rs2.Push(2)
rs2.Push(6)
rs2.RunningStatCount() // 3
rs2.RunningStatMean() // 3
rs2.RunningStatDeviation() // 2.6458
rs2.RunningStatVariation() // 7

rs.Merge(rs2)
rs.RunningStatCount() // 6
rs.RunningStatMean() // 4
rs.RunningStatDeviation() // 2
rs.RunningStatVariation() // 4
```

RunningStat is adapted from https://gist.github.com/turnersr/11390535, which in turn credits:
 1. Numerically Stable, Single-Pass, Parallel Statistics Algorithms - http://www.janinebennett.org/index_files/ParallelStatisticsAlgorithms.pdf
 2. Accurately computing running variance - http://www.johndcook.com/standard_deviation.html
