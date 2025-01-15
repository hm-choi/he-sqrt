## Experimental Evaluation

There are five example codes for the performance evaluation of CryptoRoot.


### Experiment 1: Performance Comparison between Pivot-Tangent and CryptoInvSqrt.
- For the performance comparison, we evaluate the MAE and MRE between Pivot-Tangent and CryptoInvSqrt. 

#### Evaluation result samples of Experiment 1
|Algorithm|depth|∆|log$$PQ$$|MAE|MRE|Time (s)|
|------|---|---|---|---|---|---|
|$$Pivot-Tangent$$|47|40|2184|$$7.6752×10^{-4}$$|$$1.3135×10^{-3}$$|63.5327|
|$$Pivot-Tangent$$|47|42|2278|$$7.4206×10^{-4}$$|$$6.5749×10^{-4}$$|63.4373|
|$$Pivot-Tangent$$|47|46|2466|$$3.9691×10^{-3}$$|$$3.8667×10^{-3}$$|63.7479|
|$$Pivot-Tangent$$|47|50|2654|$$1.1102×10^{-3}$$|$$1.0090×10^{-3}$$|63.7831|

|Algorithm|depth|∆|log$$PQ$$|MAE|MRE|Time (s)|
|------|---|---|---|---|---|---|
|$$CryptoRoot$$|19|40|1064|$$4.6528×10^{-4}$$|$$4.6512×10^{-4}$$|18.7776|
|$$CryptoRoot$$|19|42|1102|$$1.2272×10^{-6}$$|$$1.3497×10^{-6}$$|18.7891|
|$$CryptoRoot$$|19|46|1178|$$1.0957×10^{-6}$$|$$1.0846×10^{-6}$$|18.6030|
|$$CryptoRoot$$|19|50|1254|$$5.0487×10^{-8}$$|$$3.9618×10^{-8}$$|18.8770|

### Experiment 2: Impact of Chebyshev Polynomial Degree on Newton’s Method Performance.
- In this experiment, we check how much the initial ciphertext generated according to the order of the Chebyshev polynomial affects Newton’s method.

#### Evaluation result samples of Experiment 2
MAE and MRE between $$Pivot-Tangent$$ and $$CryptoInvSqrt$$. Evaluation is conducted on the domain $$[0.001, 1,000]$$ and all experiment is conducted with the scale factor $$∆=40$$.
|Degree|depth|∆|log$$PQ$$|MAE|MRE|Time (s)|
|------|---|---|---|---|---|---|
|$$2^{7}-2$$|16|40|944|$$3.4751×10^{-3}$$|$$5.9317×10^{-4}$$|7.0801|
|$$2^{8}-2$$|17|40|984|$$5.3503×10^{-4}$$|$$4.6762×10^{-4}$$|8.8953|
|$$2^{9}-2$$|18|40|1024|$$4.6522×10^{-4}$$|$$4.6515×10^{-4}$$|14.2505|
|$$2^{10}-2$$|19|40|1064|$$4.6528×10^{-4}$$|$$4.6512×10^{-4}$$|18.7776|
|$$2^{11}-2$$|20|40|1104|$$4.6525×10^{-4}$$|$$4.6520 ×10^{-4}$$|32.8051|


### Experiment 3: Performance Evaluation of CryptoSqrt
- In this experiment, we evaluated the performance of two square root approximation methods, CryptoSqrt and a variation of CryptoInvSqrt, over the domain [0.001, 1000].

MAE and MRE between $$CryptoSqrt$$ and variance of $$CryptoInvSqrt$$ on the domain $$[0.001, 1000]$$ with running times (seconds). The experimental result were measured when the ring degree is $$N=2^{17}$$ and the scale factor $$∆=40$$.

|Degree|depth|∆|log$$PQ$$|MAE|MRE|Time (s)|
|------|---|---|---|---|---|---|
|$$CryptoSqrt$$|11|40|744|$$1.7779×10^{-4}$$|$$1.3008×10^{-4}$$|7.3074|
|$$CryptoSqrt$$|12|40|784|$$4.5984×10^{-5}$$|$$2.3033×10^{-5}$$|13.1067|
|$$CryptoSqrt$$|13|40|824|$$3.6760×10^{-5}$$|$$1.2880×10^{-5}$$|22.0763|
|$$CryptoSqrt$$|14|40|864|$$3.5828×10^{-6}$$|$$1.2505×10^{-5}$$|47.7569|

|Degree|depth|∆|log$$PQ$$|MAE|MRE|Time (s)|
|------|---|---|---|---|---|---|
|$$V of InvSqrt$$|19|40|1064|$$8.6697×10^{-3}$$|$$7.9648×10^{-4}$$|15.5055|
|$$V of InvSqrt$$|20|40|1104|$$8.6695×10^{-3}$$|$$7.9648×10^{-4}$$|21.1236|
|$$V of InvSqrt$$|21|40|1144|$$8.6697×10^{-3}$$|$$7.9649×10^{-4}$$|35.9899|
|$$V of InvSqrt$$|22|40|1184|$$8.6696×10^{-3}$$|$$7.9648×10^{-4}$$|54.5669|

### Experiment 4: Evaluation in Various Domains of CryptoInvSqrt
- To demonstrate that CryptoRoot guarantees consistent performance across various domains, we conducted experiments in three extreme domains: $$[10^{-4}, 10^{4}], [10^{−6}, 10^{2}]$$, and $$[10^{-2}, 10^{6}]$$.

Performance of $$CryptoInvSqrt$$ on various domain.
|Degree|depth|∆|log$$PQ$$|MAE|MRE|Time (s)|
|------|---|---|---|---|---|---|
|$$10^{-4}, 10^{4}$$|21|40|1144|$$4.6581×10^{-4}$$|$$4.6550×10^{-4}$$|50.5904|
|$$10^{-6},10^{2}$$|21|40|1144|$$5.3503×10^{-2}$$|$$4.6753×10^{-4}$$|49.7820|
|$$10^{-2},10^{6}$$|21|40|1144|$$4.6522×10^{-5}$$|$$2.0281×10^{-3}$$|49.1332|

### Experiment 5: Key and Ciphertext Storage Performance
- In this experiment, we measure the size of the set of keys and ciphertext with $$N=2^{17}$$, and $$∆ = 40$$.

Size of secret key (SK), public key (PK), Relinearlization Key (RLK), Galois Key (GLK), and Ciphertext size (CTXT) with $$N=2^{17}$$ and $$∆=40$$.

|Degree|depth|∆|log$$PQ$$|MAE|MRE|Time (s)|
|------|---|---|---|---|---|---|
|$$Pivot-Tangent$$|47|54.5264|109.0528 |1308.6334|1308.6334|100.6644|
|$$CryptoRoot$$|19|24.1174|48.2349|241.1746|241.1746|39.8465|

## Usage of Test
To run the n-th test, run the following command:
```
 go test experiment{n}_test.go -v
```
The -v option shows the detailed result of the experiment.

## 