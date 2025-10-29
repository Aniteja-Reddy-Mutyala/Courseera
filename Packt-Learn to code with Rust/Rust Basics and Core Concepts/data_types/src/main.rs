fn main() {
    println!("Hi Everyone.\n.I am in new line");
    println!("I said \"Hi everyone\"");
    let path=r"HI.\n";
    println!("{path}");
    let value:i32=-20;
    println!("{}",value.abs());
    let empty_space:&str="       messsage  ";
    println!("{}",empty_space.trim());
    println!("{}",value.pow(3));
    let pi:f64=3.141598908765544321213445677888;
    println!("The value of pi is : {pi}");
    println!("Floor of the pi is {}",pi.floor());
    println!("Ceil of pi is {}",pi.ceil());
    println!("Pi is rounded to {}",pi.round());
}