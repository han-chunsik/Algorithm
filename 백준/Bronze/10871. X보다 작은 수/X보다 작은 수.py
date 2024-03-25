a,b = map(int, input().split())
num_list = list(input().split())

result_list = []

for i in num_list:
    if int(i) < b:
        result_list.append(i)

print(" ".join(result_list))