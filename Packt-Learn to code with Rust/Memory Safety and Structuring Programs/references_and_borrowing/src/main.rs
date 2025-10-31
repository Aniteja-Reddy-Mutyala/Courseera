fn main() {
    //let mut current_meal:String=String::new();
    //add_flour(& mut current_meal);
    //show_my_meal(&current_meal);
    //let mut car:String=String::from("Red");
    //let car1:& mut String=& mut car;
    //let car2:&String=&car;
    //let car3:&String=&car;
    //println!("{car2},{}",&car);
    // mutable and immutable reference to same address can't coexist at same time,i.e last used of mutable reference must be before the first use of immutable
       let registrations=(true,false,true);
       let first=registrations.0;
       println!("{first},{registrations:?}");
       let languages=[String::from("Rust"),String::from("Go")];
       let go=&languages[1];
       println!("{}",go);
}
/*fn add_flour( meal:& mut String){
    meal.push_str("Add flour");
}    

fn show_my_meal( meal:&String){
print!("Meal steps:{meal}");

}*/
