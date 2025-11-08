mod inventory;
mod orders;
use inventory::{Item, ProductCategory};
use inventory::{FLOOR_SPACE,MANAGER as  INVENTORY_MANAGER};
use orders::MANAGER as ORDERS_MANAGER;

fn main() {
    println!("Hi {}", inventory::MANAGER);
    println!("My orders manager is {}", orders::MANAGER);
    println!(
        "Our managers are {} and {}.We have room space for {} sft.",
        INVENTORY_MANAGER,
        ORDERS_MANAGER,
        FLOOR_SPACE
    );

    let my_favourite_product = ProductCategory::Hammer;
    println!("{my_favourite_product:?}");
    let check = ProductCategory::Hammer;
    println!("{check:?}");
    let tall_ladder = Item::new(String::from("Huge ladder"), my_favourite_product, 100);

    println!("{tall_ladder:#?}");
}
