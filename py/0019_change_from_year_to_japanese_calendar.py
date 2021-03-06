import datetime

input_line = input()
input_line_a = input_line.split(" ")
y = int(input_line_a[0])
m = int(input_line_a[1])
d = int(input_line_a[2])
date = datetime.date(y, m, d)

# ・明治は1912年7月29日まで
# ・大正は1912年7月30日から1926年12月24日まで
# ・昭和は1926年12月25日から1989年1月7日まで
# ・平成は1989年1月8日から2019年4月30日まで
# ・令和は2019年5月1日から
if date <= datetime.date(1912, 7, 29):
    print(f"明治年{m}月{d}日")
elif datetime.date(1912, 7, 30) <= date <= datetime.date(1926, 12, 24):
    print(f"大正年{m}月{d}日")
elif datetime.date(1926, 12, 25) <= date <= datetime.date(1989, 1, 7):
    print(f"昭和年{m}月{d}日")
elif datetime.date(1989, 1, 8) <= date <= datetime.date(2019, 4, 30):
    print(f"平成年{m}月{d}日")
else:
    print(f"令和年{m}月{d}日")

