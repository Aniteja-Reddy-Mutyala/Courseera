fn main() {
    let mut current_meal:String=String::new();
    add_flour(& mut current_meal);
    show_my_meal(&current_meal);
    let car:String=String::from("Red");
    let car1:&String=&car;
    let car2:&String=&car;
    println!("{car1},{car2},{}",&car);

}
fn add_flour( meal:& mut String){
    meal.push_str("Add flour");
    
}
fn show_my_meal( meal:&String){
print!("Meal steps:{meal}");

}
