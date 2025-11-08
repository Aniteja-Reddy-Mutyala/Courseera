mod inventory;
mod orders;

fn main() {
    println!("Hi {}", inventory::MANAGER);
    println!("My orders manager is {}", orders::MANAGER);
}
