mod inventory;
mod orders;
use inventory::products::{Item, ProductCategory};
use inventory::{FLOOR_SPACE, MANAGER};

fn main() {
    println!("Hi {}", inventory::MANAGER);
    println!("My orders manager is {}", orders::MANAGER);
    println!(
        "Our managers are {} and {}.We have room space for {} sft.",
        MANAGER,
        orders::MANAGER,
        FLOOR_SPACE
    );

    let my_favourite_product = ProductCategory::Hammer;
    println!("{my_favourite_product:?}");
    let check = ProductCategory::Hammer;
    println!("{check:?}");
    let tall_ladder = Item::new(String::from("Huge ladder"), my_favourite_product, 100);

    println!("{tall_ladder:#?}");
}
