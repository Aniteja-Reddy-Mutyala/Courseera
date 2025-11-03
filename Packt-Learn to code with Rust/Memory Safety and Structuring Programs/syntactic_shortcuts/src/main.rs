fn main() {
   let action_hero:String=String::from("batman superman");
   let first_hero:&str=&action_hero[..6];
   println!("First hero is {first_hero}");
   let second_hero:&str=&action_hero[7..];
   println!("Last hero is {second_hero}");
   let full_name:&str=&action_hero[..];
   println!("Full names are {full_name}");
   let action_hero_ref=&action_hero;
   do_hero_stuff(action_hero_ref);
   println!("{action_hero_ref}");
   let values = [2,3,7,4,5];
   let my_slice=&values[2..4];
   println!("My slice is {my_slice:?}");
   print_length(my_slice);
   let mut my_array=[10,20,30,40,50];
   let my_array_slice=& mut my_array[2..4];
   my_array_slice[1]=25;
   println!("My array slice after modifiying {my_array_slice:?}");

    


}
fn do_hero_stuff(hero_name:&str){
    println!("{hero_name} saves the day");
}
fn print_length(slice:&[i32]){
    println!("length is {}",slice.len());
}
