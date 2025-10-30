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
   let cake=bake_cake();
   println!(" I now have a cake named {cake}");
}

fn bake_cake()->String{
   let cake:String=String::from("Moose");
   return cake;
}
