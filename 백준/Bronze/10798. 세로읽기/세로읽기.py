ws = []
for i in range(5):
    ws.append(input())

result = ""
for i in range(15):
    for j in range(5):
        if len(ws[j]) > i:
            result += ws[j][i]

print(result)