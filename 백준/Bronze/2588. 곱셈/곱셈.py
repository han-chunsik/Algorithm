a = int(input())
b = int(input())

num_list_b = list(map(int, str(b)))[::-1]

for i in num_list_b:
    print(a*i)
print(a*b)