use warehouse::{FLOOR_SPACE, INVENTORY_MANAGER, ORDERS_MANAGER};
/// get a summary of our current managers
fn main() {
    println!(
        "Our managers are {} and {}.we have {} sft floor space",
        INVENTORY_MANAGER, ORDERS_MANAGER, FLOOR_SPACE
    );
}
