fn main() {
    struct Coffee {
        price: f64,
        name: String,
        is_hot: bool,
    }
    let mut mocha: Coffee = Coffee {
        price: 4.99,
        name: String::from("Mocha"),
        is_hot: false,
    };

    println!("My morning {} costs {} and it is {} that it was hot.",mocha.name,mocha.price,mocha.is_hot);
    mocha.price=6.00;
    println!("The latest price of {} is {} ",mocha.name,mocha.price);
}
