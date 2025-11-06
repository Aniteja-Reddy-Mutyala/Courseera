#[derive(Debug)]
struct TreasureChest<T> {
    captain: String,
    treasure: T,
}
impl TreasureChest<String>{
    fn clean_treasure(&mut self){
       self.treasure= self.treasure.trim().to_string();
    }
}
impl TreasureChest<[&str;3]>{
   fn amount_of_treasure(&mut self)->usize{
    self.treasure.len()
   }
}
impl <T> TreasureChest<T>{
    fn capital_captain(& mut self){
        self.captain=self.captain.to_uppercase();
    }
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
    let mut gold_chest=TreasureChest{
        captain:String::from("Goldberg"),
        treasure:"Gold"
    };
    gold_chest.capital_captain();
    let mut silver_chest=TreasureChest{
        captain:String::from("Silverberg"),
        treasure:String::from("   Silver    ")
    };
    silver_chest.clean_treasure();
    silver_chest.capital_captain();
    let mut  special_chest=TreasureChest{
        captain:String::from("special"),
        treasure:["Gold","Silver","Platinum"]
    };
    special_chest.capital_captain();
    let length=special_chest.amount_of_treasure();
    println!("{special_chest:?}");
    println!("{gold_chest:?}");
    println!("{silver_chest:?}");
    println!("{length}");


}
