mod inventory;
mod orders {
    pub const MANAGER: &str = "Aniteja orders";
}
fn main() {
    println!("Hi {}", inventory::MANAGER);
    println!("My orders manager is {}", orders::MANAGER);
}
