b = int(input())
n = int(input())
b_a = [int(c) for c in str(b)]

def is_adjacent(a, b):
    return a == b - 1 or a == b + 1

def is_second(a, b, a_a, b_a):
    return a_a[-1] == b_a[-1]  and a_a[-2] == b_a[-2] and a_a[-3] == b_a[-3] and a_a[-4] == b_a[-4]

def is_third(a, b, a_a, b_a):
    return a_a[-1] == b_a[-1]  and a_a[-2] == b_a[-2] and a_a[-3] == b_a[-3]

for i in range(n):
    a = int(input())
    a_a = [int(c) for c in str(a)]

    if a == b:
        print("first")
    elif is_adjacent(a, b):
        print("adjacent")
    elif is_second(a, b, a_a, b_a):
        print("second")
    elif is_third(a, b, a_a, b_a):
        print("third")
    else:
        print("blank")

