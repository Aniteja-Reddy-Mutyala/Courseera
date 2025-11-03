fn main() {
    struct Coffee {
        price: f64,
        name: String,
        is_hot: bool,
    }
    let mocha: Coffee = Coffee {
        price: 4.99,
        name: String::from("Mocha"),
        is_hot: false,
    };
}
