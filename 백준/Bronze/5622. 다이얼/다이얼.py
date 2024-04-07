s = input()
t = 0

for j in s:
    if 'A' <= j <= 'C':
        n = 2
    elif 'D' <= j <= 'F':
        n = 3
    elif 'G' <= j <= 'I':
        n = 4
    elif 'J' <= j <= 'L':
        n = 5
    elif 'M' <= j <= 'O':
        n = 6
    elif 'P' <= j <= 'S':
        n = 7
    elif 'T' <= j <= 'V':
        n = 8
    elif 'W' <= j <= 'Z':
        n = 9
    t += n+1
print(t)