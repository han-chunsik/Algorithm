n = int(input())

s_list = []

for i in range(n):
    s = input()
    s_list.append(f"{s[0]}{s[-1]}")

for j in s_list:
    print(j)