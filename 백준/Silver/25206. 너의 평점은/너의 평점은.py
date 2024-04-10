total_gp = 0
sum_gp_times_g = 0

def g_to_gp(grade):
    if grade == "A+":
        gp = 4.5
    elif grade == "A0":
        gp = 4.0
    elif grade == "B+":
        gp = 3.5
    elif grade == "B0":
        gp = 3.0
    elif grade == "C+":
        gp = 2.5
    elif grade == "C0":
        gp = 2.0
    elif grade == "D+":
        gp = 1.5
    elif grade == "D0":
        gp = 1.0
    elif grade == "F":
        gp = 0.0
    return gp

for n in range(20):
    sub, gp, g = input().split()
    if g == "P":
        pass
    else:
        sum_gp_times_g+=(float(gp) * g_to_gp(g))
        total_gp+=float(gp)


print(sum_gp_times_g/total_gp)