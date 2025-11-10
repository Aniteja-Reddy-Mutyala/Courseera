use std::fs::File;
use std::io::stdin;
use std::process;

fn main() {
    let array=[1,2];
    println!("Please enter the name of file");
    let mut input =String::new();
       //exit(1);
   // println!("{:?}",array[2]);
    //panic!("Something went wrong");
   
   match stdin().read_line(&mut input){
    Result::Ok(_)=>{}
    Result::Err(_)=>{eprintln!("Something went wrong");process::exit(1);}
   }
    let file=File::open(input.trim());
    match file{
        Result::Ok(value)=>{println!("{value:?}");}
        Result::Err(_)=>{eprintln!("Error in opening file");}
    }

   // eprintln!("Something went wrong!!!")
}
