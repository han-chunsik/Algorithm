result_list = []
for _ in range(10):
    a = int(input())
    result_list.append(a%42)

print(len(set(result_list)))