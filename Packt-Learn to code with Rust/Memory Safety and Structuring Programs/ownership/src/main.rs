fn main() {
   let my_value:i32=10;
   let reference:&i32=&my_value;
   println!("{}",*reference);
    let my_string_value=String::from("Welcome");
    let borrowed_value:&String=&my_string_value;
   
    println!("{borrowed_value:p}");
   

}
