import string
s = input()
a_list = list(string.ascii_lowercase)

result_list = []

for a in a_list:
    if a in s:
        for i, c in enumerate(s):
            if a == c:
                result_list.append(i)
                break
    else:
        result_list.append(-1)
        
for k in range(len(result_list)):
    print(result_list[k], end=" ")