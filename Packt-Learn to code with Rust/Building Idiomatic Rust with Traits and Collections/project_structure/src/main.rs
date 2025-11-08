use fake::{Fake, Faker};
use warehouse::{FLOOR_SPACE, INVENTORY_MANAGER, Item, ORDERS_MANAGER, ProductCategory};
///primary entry point into our warehouse program
fn main() {
    println!("Hi {}", INVENTORY_MANAGER);
    println!("My orders manager is {}", ORDERS_MANAGER);
    println!(
        "Our managers are {} and {}.We have room space for {} sft.",
        INVENTORY_MANAGER, ORDERS_MANAGER, FLOOR_SPACE
    );

    let my_favourite_product = ProductCategory::Hammer;
    println!("{my_favourite_product:?}");
    let check = ProductCategory::Hammer;
    println!("{check:?}");
    let tall_ladder = Item::new(String::from("Huge ladder"), my_favourite_product, 100);

    println!("{tall_ladder:#?}");
    let fake_item: Item = Faker.fake();
    println!("{fake_item:?}");
}
