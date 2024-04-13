mtx = []
max_len = 0
for i in range(5):
    line = list(input().replace(""," ").split())
    if max_len < len(line):
        max_len = len(line)
    mtx.append(line)

for i in mtx:
    while len(i) < max_len:
        i.extend(["@"])

result = ""
for i in range(max_len):
    for j in range(5):
        result += mtx[j][i]

print(result.replace("@",""))