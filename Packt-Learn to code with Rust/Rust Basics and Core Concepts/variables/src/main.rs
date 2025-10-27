fn main() {
    let  apples = 50;
    let oranges = "20";
    let oranges=20;
    let fruits = apples + oranges;
    
    println!(
        "My garden has {0} apples and {1} oranges.I did i got my desired target of {0} apples",
        apples, oranges
    );
    {    
        let  fruit_price=2;
        println!("Inside inner scope.No of fruits are {}.Total price is {}.",fruits,fruits*fruit_price);
    }
}
