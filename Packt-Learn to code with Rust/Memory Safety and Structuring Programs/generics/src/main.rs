#[derive(Debug)]
struct TreasureChest<T> {
    captain: String,
    treasure: T,
}
fn identity<T>(value: T) -> T {
    value
}
fn make_tuple<T, U>(first: T, second: U) -> (T, U) {
    (first, second)
}
fn main() {
    println!("Hello, world!");
    let i32_value = identity::<u32>(5);
    println!("identity is {i32_value}");
    let bool_value = identity(true);
    println!("Identity is {bool_value}");
    let f64_value = identity::<f32>(9.056);
    println!("Identity is {f64_value}");
    let new_tuple = make_tuple(String::from("hello"), 10);
    println!("{new_tuple:?}");
    let gold_chest=TreasureChest{
        captain:String::from("Goldberg"),
        treasure:"Gold"
    };
    let silver_chest=TreasureChest{
        captain:String::from("Silverberg"),
        treasure:String::from("Silver")
    };
    let special_chest=TreasureChest{
        captain:String::from("special"),
        treasure:["Gold","Silver","Platinum"]
    };
    println!("{special_chest:?}");
    println!("{gold_chest:?}");
    println!("{silver_chest:?}");


}
