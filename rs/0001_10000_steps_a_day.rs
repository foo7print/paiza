use std::io::Read;

fn main(){
  let mut buf = String::new();

  // 標準入力をStringで読み込み。IOは失敗するのでunwrap()
  std::io::stdin().read_to_string(&mut buf).unwrap();

  // 空白区切りのiteratorに分割
  let mut iter = buf.split_whitespace();
  
  // iteratorからnext()で値を取り出し、parseでi32に変換。next(), parse()ともに失敗するのでunwrap()
  let a: i32 = iter.next().unwrap().parse().unwrap();
  let b: i32 = iter.next().unwrap().parse().unwrap();
  
  let steps_count = (a * 1000 * 100) / b;
  let answer = if steps_count >= 10000 { "yes" } else { "no" };
  println!("{}", answer);
}

