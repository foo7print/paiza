n = int(input())
a = []
for i in range(n):
    w = input()
    if w in a:
        a.remove(w)
    a.insert(0, w)

for str in a:
    print(str)

