fn main() {
   let mut food:&str="apple";
   println!("{food}");
   food="banana";
   println!("{food}");
   let text=String::new();
    let mut new_text:String=String::from("Hello");
    let len:i32=new_text.capacity().try_into().unwrap();
    println!("{len}");
    new_text.push_str("Everyone");
    println!("{new_text}");

}
