a = int(input())

result_list = []

for i in range(a):
    p = ""
    r, s = input().split()
    for j in range(len(s)):
        p = p + (s[j]*int(r))
    result_list.append(p)

for result in result_list:
    print(result)