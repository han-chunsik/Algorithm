h, m = map(int, input().split())
t = int(input())

if 60 <= t:
    h+=(t // 60)
    m+=(t % 60)
    if 60 <= m:
        h+=1
        m-=60
    if 24 <= h:
        h-=24
else:
    m+=t
    if 60 <= m:
        h +=1
        m-=60
    if 24 <= h:
        h-=24

print(h, m)   
