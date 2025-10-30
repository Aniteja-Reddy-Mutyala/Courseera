fn main() {
   let my_value:i32=10;
   let reference:&i32=&my_value;
   println!("{}",*reference);
    let my_string_value=String::from("Welcome");
    let borrowed_value:&String=&my_string_value;
   
    println!("{borrowed_value:p}");
    let mut string="hi";
    string="hello";
   let  mut one:i32=5;
   let two:&i32=&one;
   let three:&i32=&one;
   println!("{three}");
  let apples:String=String::from("oranges");
  print_my_value(apples);
   println!("{apples}");
}

fn print_my_value(value:String){
    println!("My value is {value}");
}
