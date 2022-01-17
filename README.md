# Color Thief

Fork https://github.com/legion-zver/color-thief

A high performance Golang module for grabbing color palettes from images. Instead of 
implementing conventional Modified Median Color Quantization (MMCQ), it implements Xiaolin Wu's Color Quantizer[[1]](#1) as well as
Weighted Sort-Means + Wu algorithm[[2]](#2). They both yield 
much better color quantization result from the evaluation.[[2]](#2).

### performance:
#### Wu's Color Quantizer
 ```
goos: darwin
goarch: amd64
pkg: color-thief/wu
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkQuantWu
BenchmarkQuantWu-12    	    1315	    914729 ns/op
PASS
```

#### WSM-WU Color Quantizer
```
goos: darwin
goarch: amd64
pkg: color-thief/wsm
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkWSM
BenchmarkWSM-12    	     244	   4648723 ns/op
PASS
```
## Reference
 - <a id="1">[1]</a>
   X. Wu, Graphics Gems Volume II, Academic Press, 1991, Ch. Efficient Statistical Computations for Optimal Color Quantization, pp. 126–133.
 - <a id="2">[2]</a>
   Celebi, M. Emre (2011).
   Improving the performance of k-means for color quantization.
   Image and Vision Computing 29, 260–271.
   
 
