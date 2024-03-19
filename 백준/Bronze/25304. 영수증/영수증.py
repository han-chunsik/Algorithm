total = int(input())
kind = int(input())

buy_list = []

for i in range(0,kind):
    a,b = map(int,input().split())
    buy_list.append(a*b)

if total == sum(buy_list):
    print("Yes")
else:
    print("No")