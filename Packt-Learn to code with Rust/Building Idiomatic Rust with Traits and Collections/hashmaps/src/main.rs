
use std::collections::HashMap;
fn main() {
    let mut menu:HashMap<String, f64>=HashMap::new();
    menu.insert("Steak".to_string(), 29.99);
    menu.insert("Tuna".to_string(), 29.99);
    menu.insert("Chicken".to_string(),25.99);
    println!("{menu:?}");
    let mut country_capitals:HashMap<&str,&str>=HashMap::new();
    country_capitals.insert("France","Paris");
    country_capitals.insert("Germany","Berlin");
    println!("{country_capitals:?}");

}
