#[derive(Debug)]
enum CardSuit{
    Hearts,
    Diamonds,
    Spades,
    Clubs
}
#[derive(Debug)]
struct Credentials{
    username:String,
    password:String
}
#[derive(Debug)]
enum PaymentMethodType{
    CreditCard(String),
    DebitCard(String),
    PayPal{username:String,password:String}
}
struct Card{
    rank:String,
    suit:CardSuit
}
#[derive(Debug)]
enum Meat{
    Chicken,
    Steak
}
#[derive(Debug)]
enum RestaurantItem{
   Burrito(Meat),
   Bowl(Meat),
   VeganPlate
}
fn main() {
    let first_card:CardSuit=CardSuit::Hearts;
    let mut second_card:CardSuit=CardSuit::Spades;
    second_card=CardSuit::Diamonds;
    println!("{second_card:?}");
    let card_suits=[CardSuit::Clubs,CardSuit::Spades];
    let visa =PaymentMethodType::CreditCard(String::from("1234-5678"));
    let master_card=PaymentMethodType::DebitCard(String::from("5678-9012"));
    println!("{visa:?}");
    println!("{master_card:?}");
     let paypal_credentials=PaymentMethodType::PayPal {
        username:String::from("abc@123.com"),
        password:String::from("12345")
     };
    let mut my_payment_method= PaymentMethodType::CreditCard(String::from("123"));
    my_payment_method=paypal_credentials;
     println!("{my_payment_method:?}");
     let lunch =RestaurantItem::Burrito(Meat::Steak);
     let dinner =RestaurantItem::Bowl(Meat::Chicken);
     let abandoned_meal= RestaurantItem::VeganPlate;
     println!("Lunch was {lunch:?}");
     println!("Dinner was {dinner:?}");
     println!("Abandaoned meal was {abandoned_meal:?}");
     
    
}
