n, m = map(int,input().split())

procession = []
sum_procession = []
for i in range(n*2):
    line = list(map(int, input().split()))
    procession.append(line)

for j in range(n):
    tmp = []
    for k in range(m):
        a = procession[j][k] + procession[j+n][k]
        tmp.append(a)
    sum_procession.append(tmp)

for p in sum_procession:
    print(' '.join(map(str, p))) 