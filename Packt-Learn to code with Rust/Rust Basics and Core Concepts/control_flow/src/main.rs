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
    let number:i32=9;
    match number {
        check_number if check_number % 2 == 0 => println!("{check_number} is an even number"),
        check_number if check_number % 2 != 0 => println!("{check_number} is an odd number"),
        _ => unreachable!(),
    };
    let mut seconds:i32=10;
    loop{
        println!("{seconds} seconds!");
        seconds -= 1;
        if seconds < 0{
            println!(" let s start");
            break;
        }
    }
}
