fn main() {
   let result=even_or_odd(10);
   println!("It is {result} that the number is even");
}
fn even_or_odd(number:i32)->bool{
    if number % 2 == 0{
       return true;
    }
    else{
        return false;
    }
}
