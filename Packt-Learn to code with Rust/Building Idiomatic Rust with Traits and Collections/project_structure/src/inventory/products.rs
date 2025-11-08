use fake::Dummy;
/// category of products
#[derive(Debug, Dummy)]
pub enum ProductCategory {
    Ladder,
    Hammer,
}
/// a concrete item in stock.
#[derive(Debug, Dummy)]
pub struct Item {
    pub name: String,
    pub category: ProductCategory,
    pub quantity: u32,
}
impl Item {
    /// create a new item
    pub fn new(name: String, category: ProductCategory, quantity: u32) -> Self {
        super::talk_to_manager();
        Self {
            name,
            category,
            quantity,
        }
    }
}
