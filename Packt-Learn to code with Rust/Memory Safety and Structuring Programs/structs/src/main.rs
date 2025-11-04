struct Coffee {
    price: f64,
    name: String,
    is_hot: bool,
}
fn main() {
    let mut mocha: Coffee = Coffee {
        price: 4.99,
        name: String::from("Mocha"),
        is_hot: false,
    };
    let latte = make_coffee(String::from("Latte"), 3.99, true);
    println!(
        "My morning {} costs {} and it is {} that it was hot.",
        mocha.name, mocha.price, mocha.is_hot
    );
    mocha.price = 6.00;
    println!("The latest price of {} is {} ", mocha.name, mocha.price);
    println!(
        "The latest price of {} is {} and it is {} that it is hot.",
        latte.name, latte.price, latte.is_hot
    );
    let caramel_macchiato: Coffee = Coffee {
        name: String::from("Caramel Macchiato"),
        ..mocha
    };
    println!(
        "The lastest price of {} is {} and it is {} that it is hot",
        caramel_macchiato.name, caramel_macchiato.price, caramel_macchiato.is_hot
    );
    drink_coffee(& mut mocha);
}
fn make_coffee(name: String, price: f64, is_hot: bool) -> Coffee {
    Coffee {
        price: price,
        name: name,
        is_hot: is_hot,
    }
}
/*fn drink_coffee(coffee: Coffee) {
    println!("Drinking my delicious {}", coffee.name);
}*/
fn drink_coffee(coffee:& mut Coffee){
    println!("Drinking my delicious {}", (*coffee).name);
    
}
