fn main() {
    let evaluation: bool = true;
    match evaluation {
        true => {
            println!("The evaluation is {evaluation}");
        }
        false => {
            println!("The evaluation is {evaluation}");
        }
    }
    let value: i32 = match evaluation {
        true => 20,
        false => 40,
    };
    println!("The value is {value}");
}
