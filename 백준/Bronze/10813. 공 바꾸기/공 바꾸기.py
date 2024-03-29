n, m = map(int,input().split())
b = []

for x in range(1, n+1):
    b.append(x)

for z in range(m):
    i,j = map(int, input().split())
    num = b[i-1]
    b[i-1] = b[j-1]
    b[j-1] = num

for k in range(n):
    print(b[k], end=" ")