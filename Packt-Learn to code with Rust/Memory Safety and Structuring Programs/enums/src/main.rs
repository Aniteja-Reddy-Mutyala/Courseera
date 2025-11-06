#[derive(Debug)]
enum CardSuit{
    Hearts,
    Diamonds,
    Spades,
    Clubs
}
#[derive(Debug)]
enum PaymentMethodType{
    CreditCard(String),
    DebitCard(String),
    PayPal(String, String)
}
struct Card{
    rank:String,
    suit:CardSuit
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
    let mut my_payment_method= PaymentMethodType::CreditCard(String::from("123"));
    my_payment_method=PaymentMethodType::PayPal(String::from("ani@123.com"),String::from("12345"));
     println!("{my_payment_method:?}");
}
