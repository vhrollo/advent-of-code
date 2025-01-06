from z3 import *

# Define Z3 variables
opt = Optimize()
s = BitVec('s', 64)
a, b, c = s, 0, 0

output_sequence = [2,4,1,2,7,5,1,3,4,4,5,5,0,3,3,0]

# Simulate program
for x in output_sequence:
    b = a % 8
    b = b ^ 2
    c = a / (1 << b)
    b = b ^ 3
    b = b ^ c
    opt.add((b % 8) == x)
    a = a / (1 << 3)
opt.add(a == 0)
opt.minimize(s)
assert str(opt.check()) == "sat"
print(opt.model().eval(s))