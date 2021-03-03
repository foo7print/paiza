n = int(input())
d = {
    "ball": 0,
    "strike": 0
}
for str in range(n):
    s = input()
    if s == "ball":
        d["ball"] += 1
        if d["ball"] == 4:
            print("fourball!")
        else:
            print("ball!")
    elif s == "strike":
        d["strike"] += 1
        if d["strike"] == 3:
            print("out!")
        else:
            print("strike!")

