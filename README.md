## Results 2025-01-26

```golang
const NB_REQUEST = 50_000
const SEMAPHORE_SIZE = 800
```

```
V1 results:  
2903 requests done in 44.77s - 64.84req/s  
2857 errs (98.415%)  
└─┬─┬ 1037 timeouts (36.297%)  
│ └─ 1033 timeouts on host lookup (36.297%)  
├── 364 no such host (12.741%)  
├── 0 bad certificate (0.000%)  
├── 255 network is unreachable (8.925%)  
└── 1201 others (42.037%)  

V2 results:  
2637 requests done in 30.00s - 87.90req/s  
263 errs (9.973%)  
 └─┬─┬ 155 timeouts (58.935%)  
   │ └─ 29 timeouts on host lookup (58.935%)  
   ├── 21 no such host (7.985%)  
   ├── 8 bad certificate (3.042%)  
   ├── 28 network is unreachable (10.646%)  
   └── 51 others (19.392%)  

V3A results:  
15082 requests done in 30.00s - 502.74req/s  
5076 errs (33.656%)  
 └─┬─┬ 2 timeouts (0.039%)  
   │ └─ 2 timeouts on host lookup (0.039%)  
   ├── 240 no such host (4.728%)  
   ├── 0 bad certificate (0.000%)  
   ├── 2 network is unreachable (0.039%)  
   └── 4832 others (95.193%)  

2541 requests done in 30.00s - 84.70req/s  
198 errs (7.792%)  
 └─┬─┬ 70 timeouts (35.354%)  
   │ └─ 70 timeouts on host lookup (35.354%)  
   ├── 18 no such host (9.091%)  
   ├── 13 bad certificate (6.566%)  
   ├── 41 network is unreachable (20.707%)  
   └── 56 others (28.283%)  

V3C results:  
4179 requests done in 30.00s - 139.30req/s  
171 errs (4.092%)  
 └─┬─┬ 14 timeouts (8.187%)  
   │ └─ 14 timeouts on host lookup (8.187%)  
   ├── 41 no such host (23.977%)  
   ├── 18 bad certificate (10.526%)  
   ├── 12 network is unreachable (7.018%)  
   └── 86 others (50.292%)  
```
