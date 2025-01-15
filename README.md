# CryptoRoot
## CryptoRoot: Efficient and Accurate Homomorphic Evaluation of Inverse Square Root and Square Root

Official repository for **CryptoRoot: Efficient and Accurate Homomorphic Evaluation of Inverse Square Root and Square Root**
by 
Hyunmin Choi<sup>1,2**</sup>.

<sup>1</sup> NAVER Cloud, South Korea
<sup>2</sup> Sungkyunkwan University, South Korea

** Corresponding author

! The paper is currently under review and will be officially published once the notification is released.

## Summary
$$CryptoRoot$$ is an efficient and accurate homomorphic evaluation architecture for Inverse square roots (InvSqrt) operation and square roots (Sqrt) operation. $$CryptoRoot$$ employs two strategies to efficiently and accurately perform HE-based InvSqrt and Sqrt operations: (1) $$CryptoInvSqrt$$: A method that determines a suitable initial value for Newton's method to accurately compute HE-based InvSqrt while reducing the $$depth$$. (2) $$CryptoSqrt$$: A technique designed to perform HE-based Sqrt operations with minimal $$depth$$ requirements. Numerical experiments demonstrate the high efficiency and accuracy of $$CryptoRoot$$. The $$CryptoInvSqrt$$ method achieves a mean absolute error of $$4.6528 × 10^{-4}$$ and a mean relative error of $$4.6512 × 10^{-4}$$ with a computation time of 18.7776 seconds, which is approximately 3.38 times faster than \pandas's runtime of 63.5327 seconds. The \sqrts computes HE-based Sqrt with an MAE of $$1.7779 × 10^{-5}$$ and an MRE of $$1.3008 × 10^{-4}$$, within 7.3074 seconds.

## Suggestion
Any suggestions or errors for better research are always welcome. Please post them on Issue, and contact us anytime at one of the following e-mails for collaboration proposals, etc.
- fieldmedalist@gmail.com
- hyunmin.choi@g.skku.edu
- hyunmin.choi@navercorp.com

## 1. Server Setting
- In this evaluation, a Macbook Pro is used.
  - OS: macOS
  - RAM: 32GB
  - Core: Apple M1 Pro
- The experiment will work well even in a Linux environment.

## 2. Installation
### Requirements
- Go: go1.23 or higher version
- Lattigo V6 library (https://github.com/tuneinsight/lattigo)

## 3. Run codes
- The five examples are introduced in the following directory.
```
./example
```

The detailed explanation and how to run the example code is in the README.md in the example directory.

## License
This is available for non-commercial purposes only. 
