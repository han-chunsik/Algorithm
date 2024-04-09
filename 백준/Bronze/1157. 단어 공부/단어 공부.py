s = input().upper()
s_list = list(set(s))

cnt = []
for i in s_list:
    chr_count = s.count
    cnt.append(chr_count(i))

if cnt.count(max(cnt)) > 1:
    print("?")
else:
    print(s_list[cnt.index(max(cnt))])