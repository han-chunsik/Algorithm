nums="0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
n, b = input().split()

n = n[::-1]
b = int(b)

sum = 0

for i, n_str in enumerate(n):
    nf = nums.find(n_str)
    sum += nf*b**i
print(sum)