## Engine

Engine is the main module of CryptoRoot.
Engine consists of four sub-modules:

### crypto_inv_sqrt
The module provides the HE-based inverse square root operation using our method: CryptoInvSqrt.

- CryptoInvSqrt 
  - Input
    - eval *ckks.Evaluator: Evaluator of CKKS scheme
    - params ckks.Parameters: Parameters for CKKS operations.
    - ct *rlwe.Ciphertext: Input ciphertext.
    - d int: Degree (Log) of Chebyshev polynomial.
    - A float64: Start value of the domain.
    - B float64: End value of domain.
    - types int: It is not used in this version.
  - Output
    - ct *rlwe.Ciphertext: Inverse Square Root of Input ciphertext.

- GetChebyshevPoly
  - Input
    - K float64
    - degree int
    - f64 func(x float64) (y float64)
  - Output
    - poly bignum.Polynomial

### crypt_sqrt
The module provides the HE-based square root operation using our method: CryptoSqrt.

- CryptoSqrt
  - Input
    - eval *ckks.Evaluator: Evaluator of CKKS scheme
    - params ckks.Parameters: Parameters for CKKS operations.
    - ct *rlwe.Ciphertext: Input ciphertext.
    - d int: Degree (Log) of Chebyshev polynomial.
    - A float64: Start value of the domain.
    - B float64: End value of domain.
  - Output
    - ct *rlwe.Ciphertext: Square Root of Input ciphertext.

### he_module
The module provides the basic parameters and modules for HE.  

- GetParam
  - Input
    - LogN int: Log of Ring degree.
    - LEVEL int: Level of ciphertext.
    - SCALE int: Scale factor of ciphertext. Default is 40. 

  - Output
    - ct ckks.Parameters: Parameters for CKKS operations.

### newton_iteration
The module provides the Newton-Rapshon's iteration method (Newton's method). 

- NewtonInvSqrt
 - Input
  - x0: Initial number
  - d: Iteration number
  - y0: Initial point. If y is nil then y is set to 1.0
- Output
  - y: Result of approximation

- HENewtonInvSqrt
  - Input
    - eval *ckks.Evaluator: Evaluator of CKKS scheme
    - params ckks.Parameters: Parameters for CKKS operations.
    - x0 *rlwe.Ciphertext: Input ciphertext.
    - d int: Degree (Log) of Chebyshev polynomial.
    - y *rlwe.Ciphertext: Initial ciphertext (value) of Newton's method.
  - Output
    - ct *rlwe.Ciphertext: Result of Newton's method.

### pivot_tangent
The implementation of the architecture (Pivot-Tangent) in panda' et al [1].

- TwoLineApprox
  - Input
    - eval *ckks.Evaluator: Evaluator of CKKS scheme
    - params ckks.Parameters: Parameters for CKKS operations.
    - ecd *ckks.Encoder: Encoder of CKKS scheme.
    - enc *rlwe.Encryptor: Encryptor of CKKS scheme.
    - x0 *rlwe.Ciphertext: Input ciphertext.
    - d int: Degree (Log) of Chebyshev polynomial.
    - A float64: Start value of the domain.
    - B float64: End value of domain.
  - Output 
    - ct *rlwe.Ciphertext: Output of the initial value using the Pivot-Tangent method.

[1] Panda, Samanvaya. "Polynomial approximation of inverse sqrt function for fhe." International Symposium on Cyber Security, Cryptology, and Machine Learning. Cham: Springer International Publishing, 2022.