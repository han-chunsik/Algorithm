sum_list = []
while True:
    a,b = map(int,input().split())
    if a==0 and b==0:
        break
    else:
        sum_list.append(a+b)

for i in sum_list:
    print(i)