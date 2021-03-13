let input_line = readLine()!
let input_line_a = input_line.split { $0 == " " }
let a = Int(input_line_a[0])!
let b = Int(input_line_a[1])!
print(a + b)

