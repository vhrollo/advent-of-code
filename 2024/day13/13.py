import numpy as np

# this was too good to not use lin alg
M = np.fromregex('input.txt', r'\d+', [('', int)]
    *6).view(int).reshape(-1,3,2).swapaxes(1,2)

S = np.kron(np.eye(2), M[..., :2])
P = np.concatenate([M[..., 2:], M[..., 2:]+1e13], 1)
R = np.linalg.solve(S, P).round().astype(np.int64)
print(*np.sum((R.reshape(-1,2,2) @ [3,1]) * (S @ R == P).reshape(-1,2,2).all(2), 0))