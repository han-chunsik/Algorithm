a = int(input())

val_dict = {}

for i in range(0, a):
    a,b = map(int, input().split())
    val_dict[i+1] = a+b

for key, value in val_dict.items():
    print(f"Case #{key}: {value}")