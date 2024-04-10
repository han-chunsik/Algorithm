n = int(input())

count = 0
for _ in range(n):
    s = input()
    last_index = 0
    for c in s:
        if last_index > s.index(c):
            count-=1
            break
        else:
            last_index = s.index(c)
    count+=1
print(count)