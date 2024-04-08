chess = [1,1,2,2,2,8]

nums = input().split()

result = []

for i in range(6):
    if chess[i] < int(nums[i]):
        a = abs(chess[i] - int(nums[i])) * -1
    else:
        a = abs(chess[i] - int(nums[i]))
    result.append(a)

for j in result:
    print(j,end=" ")