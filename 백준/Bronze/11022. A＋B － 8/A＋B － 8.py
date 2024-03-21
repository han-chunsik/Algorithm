a = int(input())

return_list = []

for i in range(0,a):
    a,b = map(int,input().split())
    return_list.append(f"Case #{i+1}: {a} + {b} = {a+b}")

for j in return_list:
    print(j)