array = [[0] * 100 for _ in range(100)]

for _ in range(int(input())):
    y, x = map(int, input().split())

    for x1 in range(x, x+10):
        for y1 in range(y, y+10):
            array[x1][y1] = 1

r = 0
for i in range(100):
    r +=array[i].count(1)

print(r)
