#[derive(Debug)]
enum CardSuit{
    Hearts,
    Diamonds,
    Spades,
    Clubs
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
}
