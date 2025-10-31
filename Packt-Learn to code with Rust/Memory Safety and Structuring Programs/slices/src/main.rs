fn main() {
   let action_hero:String=String::from("Hi Batman");
   let first_part:&str=&action_hero[0..3];
   println!("{first_part}");
   let second_part:&str=&action_hero[3..9];
   println!("{second_part}");

}
