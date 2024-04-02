s = int(input())
score_list = list(map(int,input().split()))

m = max(score_list)

score_list = list(map(lambda x: x / m * 100, score_list))
print(sum(score_list)/s)