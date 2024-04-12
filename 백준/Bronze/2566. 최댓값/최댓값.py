mtx = []
for i in range(9):
    line = list(map(int, input().split()))
    mtx.append(line)

max_line = 0
max_num = 0
for i in mtx:
    if max_num < max(i): 
        max_num = max(i)
        max_line = mtx.index(i)

print(max_num)
print(max_line+1, mtx[max_line].index(max_num)+1)
