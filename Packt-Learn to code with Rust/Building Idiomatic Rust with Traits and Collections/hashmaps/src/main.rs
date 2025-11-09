
use std::collections::HashMap;
fn main() {
    let mut menu:HashMap<String, f64>=HashMap::new();
    menu.insert("Steak".to_string(), 29.99);
    menu.insert("Tuna".to_string(), 29.99);
    menu.insert("Chicken".to_string(),25.99);
    println!("{menu:?}");
    let capitals=[("France","Paris"),("Gernamy","Berlin")];
    let mut country_capitals:HashMap<&str,&str>=HashMap::from(capitals);
   // country_capitals.insert("France","Paris");
   // country_capitals.insert("Germany","Berlin");
   // println!("{country_capitals:?}");
    match country_capitals.remove(&"Belgium"){
        Option::Some(_)=>{println!("{country_capitals:?}");}
        Option::None=>{println!("The country doesn't exist in the list");}
    }
    let france_capital= country_capitals.get(&"France");
    println!("{}",france_capital.unwrap_or(&"Country doesn't exist in list"));
    //println!("{country_capitals:?}");

}
