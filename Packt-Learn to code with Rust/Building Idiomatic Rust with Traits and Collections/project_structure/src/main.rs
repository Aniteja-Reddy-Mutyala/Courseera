mod inventory {
    const FLOOR_SPACE:i32=10000;
     pub const MANAGER:&str="Aniteja inventory";
    #[derive(Debug)]
    enum ProductCategory{
        Ladder,
        Hammer
    }
    #[derive(Debug)]
    struct Item{
        name:String,
        category:ProductCategory,
        quantity:u32
    }
    fn talk_to_manager(){
      println!("Hi {MANAGER},how is your coffee");
    }
}
mod orders {
    pub const MANAGER :&str ="Aniteja orders";
}
fn main() {
    println!("Hi {}",inventory::MANAGER);
    println!("My orders manager is {}",orders::MANAGER);
}
