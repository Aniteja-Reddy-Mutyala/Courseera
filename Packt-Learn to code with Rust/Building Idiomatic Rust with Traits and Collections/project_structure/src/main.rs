mod inventory;
mod orders;
use inventory::products::{self,ProductCategory};
use inventory::{FLOOR_SPACE,MANAGER,talk_to_manager};

fn main() {
    println!("Hi {}", inventory::MANAGER);
    println!("My orders manager is {}", orders::MANAGER);
    println!(
        "Our managers are {} and {}.We have room space for {} sft.",
        MANAGER,
        orders::MANAGER,
        FLOOR_SPACE
    );
    talk_to_manager();
    let my_favourite_product = ProductCategory::Hammer;
    println!("{my_favourite_product:?}");
    let check = ProductCategory::Hammer;
    println!("{check:?}");
    let tall_ladder = products::Item {
        name: String::from("Huge ladder"),
        category: my_favourite_product,
        quantity: 100,
    };
    println!("{tall_ladder:#?}");
}
