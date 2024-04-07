s = input()
a_list = ['ABC','DEF','GHI','JKL','MNO', 'PQRS','TUV','WXYZ']
m = 0

for i in range(len(s)):
    for j in range(8):
        if s[i] in a_list[j]:
            m+=j+3
print(m)