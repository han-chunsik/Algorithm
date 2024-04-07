s = input()
t = 0

for j in s:
    chr = ord(j)
    n = 0
    if 65 <= chr <= 67:
        n = 2
    elif 68 <= chr <= 70:
        n = 3
    elif 71 <= chr <= 73:
        n = 4
    elif 74 <= chr <= 76:
        n = 5
    elif 77 <= chr <= 79:
        n = 6
    elif 80 <= chr <= 83:
        n = 7
    elif 84 <= chr <= 86:
        n = 8
    elif 87 <= chr <= 90:
        n = 9
    t += n+1
print(t)