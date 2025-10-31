#![allow(unused_variables)]
const YEAR:i32=1996;
const TOUCHDOWN_POINTS:i32=6;
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
    let season= "Fall";
    let mut points_scored = 28;
    points_scored = 35;
    let event_time= "06:00";
    let event_time= 6 ;
    println!("My favourite season is {season}.The team scored {points_scored}.The time is {event_time}.A touchdown is worth {TOUCHDOWN_POINTS}.");

    let favourite_beverage="Apple juice" ;
     
}
