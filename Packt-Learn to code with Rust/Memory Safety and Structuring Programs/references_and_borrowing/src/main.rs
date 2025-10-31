fn main() {
    let mut current_meal:String=String::new();
    add_flour(& mut current_meal);
    show_my_meal(&current_meal);
}
fn add_flour( meal:& mut String){
    meal.push_str("Add flour");
    
}
fn show_my_meal( meal:&String){
print!("Meal steps:{meal}");
}
