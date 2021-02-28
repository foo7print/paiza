input_line = input()
a = input_line.split(" ")
output_a = []
for str in a:
    if str not in output_a:
        print(f'{str} {a.count(str)}')
    output_a.append(str)
