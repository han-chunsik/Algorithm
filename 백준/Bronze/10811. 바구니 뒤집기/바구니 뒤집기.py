n, m = map(int, input().split())
buckets = list(range(1,n+1))

for i in range(0,m):
    j,k = map(int, input().split())
    j-=1
    buckets[j:k] = buckets[j:k][::-1]
for k in range(n):
    print(buckets[k], end=" ")