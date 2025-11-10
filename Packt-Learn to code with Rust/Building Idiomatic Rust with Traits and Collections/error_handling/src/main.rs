use std::fs::File;
use std::io::{stdin,Read};
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
    let mut file:File=match File::open(input.trim()){
  
        Result::Ok(file)=>{file}
        Result::Err(error)=>{eprintln!("Error in opening file.The error is {error}");process::exit(1);}
    };
    let mut file_content=String::new();
    let read_operation=file.read_to_string(& mut file_content);
    match &read_operation{
        Result::Ok(_)=>{println!("{file_content}");}
        Result::Err(error)=>{eprint!("Unable to read file due to {error}");process::exit(1);}
    }

   // eprintln!("Something went wrong!!!")
}
