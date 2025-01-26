## Baseline (2025-01-26 )

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

## Varying SEMAPHORE_SIZE

```
SEMAPHORE_SIZE = 500
V2 results:
3190 requests done in 30.00s - 106.33req/s
105 errs (3.292%)
 └─┬─┬ 31 timeouts (29.524%)
   │ └─ 2 timeouts on host lookup (29.524%)
   ├── 25 no such host (23.810%)
   ├── 15 bad certificate (14.286%)
   ├── 4 network is unreachable (3.810%)
   └── 30 others (28.571%)

SEMAPHORE_SIZE = 600
V2 results:
2926 requests done in 30.00s - 97.53req/s
148 errs (5.058%)
 └─┬─┬ 77 timeouts (52.027%)
   │ └─ 3 timeouts on host lookup (52.027%)
   ├── 22 no such host (14.865%)
   ├── 11 bad certificate (7.432%)
   ├── 3 network is unreachable (2.027%)
   └── 35 others (23.649%)

SEMAPHORE_SIZE = 750
V2 results:
2774 requests done in 30.00s - 92.47req/s
264 errs (9.517%)
 └─┬─┬ 170 timeouts (64.394%)
   │ └─ 9 timeouts on host lookup (64.394%)
   ├── 22 no such host (8.333%)
   ├── 12 bad certificate (4.545%)
   ├── 6 network is unreachable (2.273%)
   └── 54 others (20.455%)

SEMAPHORE_SIZE = 1000
V2 results:
2737 requests done in 30.00s - 91.23req/s
410 errs (14.980%)
 └─┬─┬ 300 timeouts (73.171%)
   │ └─ 10 timeouts on host lookup (73.171%)
   ├── 26 no such host (6.341%)
   ├── 11 bad certificate (2.683%)
   ├── 5 network is unreachable (1.220%)
   └── 68 others (16.585%)

```
