## Experimental Evaluation

There are five example codes for the performance evaluation of CryptoRoot.


### Example 1: Experiment 1: Performance Comparison between Pivot-Tangent and CryptoInvSqrt.
- For the performance comparison, we evaluate the MAE and MRE between Pivot-Tangent and CryptoInvSqrt. 

### Experiment 2: Impact of Chebyshev Polynomial Degree on Newton’s Method Performance.
- In this experiment, we check how much the initial ciphertext generated according to the order of the Chebyshev polynomial affects Newton’s method.

### Experiment 3: Performance Evaluation of CryptoSqrt
- In this experiment, we evaluated the performance of two square root approximation methods, CryptoSqrt and a variation of CryptoInvSqrt, over the domain [0.001, 1000].

### Experiment 4: Evaluation in Various Domains of CryptoInvSqrt
- To demonstrate that CryptoRoot guarantees consistent performance across various domains, we conducted experiments in three extreme domains: $$[10^{-4}, 10^{4}], [10^{−6}, 10^{2}]$$, and $$ [10^{-2}, 10^{6}].$$

### Experiment 5: Key and Ciphertext Storage Performance
- In this experiment, we measure the size of the set of keys and ciphertext with $$N=2^{17}$$, and $$∆ = 40$$.


## Usage of Test
To run the n-th test, run the following command:
```
 go test experiment{n}_test.go -v
```
The -v option shows the detailed result of the experiment.

## 