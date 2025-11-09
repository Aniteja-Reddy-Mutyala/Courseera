fn main() {
  let  first_name:String=String::from("Aniteja Reddy");
  let  last_name= String::from("Mutyala");
  let full_name = first_name + &last_name;
  //full_name.push_str(last_name);
  println!("{full_name}");
  println!("{last_name}");
}
