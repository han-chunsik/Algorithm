ca = ["c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="]
s = input()
for i in ca:
    s = s.replace(i, "@")
print(len(s))