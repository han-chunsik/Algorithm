n, m = input().split()
print(max(map(lambda x: x[::-1], [n,m])))