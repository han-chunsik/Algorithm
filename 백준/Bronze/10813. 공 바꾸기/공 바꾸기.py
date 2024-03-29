n, m = map(int,input().split())

buckets = list(range(1,n+1))

for z in range(m):
    i,j = map(int, input().split())
    buckets[i-1], buckets[j-1] = buckets[j-1], buckets[i-1] 

for k in range(n):
    print(buckets[k], end=" ")