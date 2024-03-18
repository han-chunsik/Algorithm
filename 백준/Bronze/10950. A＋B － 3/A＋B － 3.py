range_num = int(input())
plus_list = []

for i in range(0,range_num):
    a,b = map(int, input().split())
    plus_list.append(a+b)

for j in plus_list:
    print(j)
