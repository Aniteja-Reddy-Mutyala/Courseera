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
    let season:&str="winter";
    if season == "summer"{
        println!("School's out");
    }
    else if season == "winter"{
        println!("So cold");
    }
    else {
        println!("Rain!!!");
    }
    match season {
        "summer" => println!("School's out"),
        "winter" => println!("So cold"),
         _ => println!("Lots of rain"),
        
    };
}
