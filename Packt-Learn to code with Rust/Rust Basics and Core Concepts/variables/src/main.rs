#![allow(unused_variables)]
const YEAR:i32=1996;
type Number=i32;
fn main() {
    let  apples :Number= 50;
    #[allow(unused_variables)]
    let oranges = "20";
    let oranges:Number = 20;
    let fruits = apples + oranges;
    
    println!(
        "My garden has {0} apples and {1} oranges in the year {YEAR}.I did i got my desired target of {0} apples",
        apples, oranges
    );
    {    
        let  fruit_price=2;
        println!("Inside inner scope.No of fruits are {}.Total price is {}.",fruits,fruits*fruit_price);
    }
}
