fn main() {
   let action_hero:String=String::from("batman superman");
   let first_hero:&str=&action_hero[..6];
   println!("First hero is {first_hero}");
   let second_hero:&str=&action_hero[7..];
   println!("Last hero is {second_hero}");
   let full_name:&str=&action_hero[..];
   println!("Full names are {full_name}");
   println!("{full_name}");
}
