mod inventory;
mod orders;

fn main() {
    println!("Hi {}", inventory::MANAGER);
    println!("My orders manager is {}", orders::MANAGER);
    println!(
        "Our managers are {} and {}.We have room space for {} sft.",
        inventory::MANAGER,
        orders::MANAGER,
        inventory::FLOOR_SPACE
    );
    inventory::talk_to_manager();
    let my_favourite_product = inventory::products::ProductCategory::Hammer;
    println!("{my_favourite_product:?}");
    let check = inventory::products::ProductCategory::Hammer;
    println!("{check:?}");
    let tall_ladder = inventory::products::Item {
        name: String::from("Huge ladder"),
        category: my_favourite_product,
        quantity: 100,
    };
    println!("{tall_ladder:#?}");
}
