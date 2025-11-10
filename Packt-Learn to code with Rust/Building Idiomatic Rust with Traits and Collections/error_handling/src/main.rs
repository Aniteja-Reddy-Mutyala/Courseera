use std::fs::File;

fn main() {
    let array=[1,2];
    println!("Hello, world!");
       //exit(1);
   // println!("{:?}",array[2]);
    //panic!("Something went wrong");
    let file=File::open("story.txt");
    match file{
        Result::Ok(value)=>{println!("{value:?}");}
        Result::Err(_)=>{eprintln!("Error in opening file");}
    }

    eprintln!("Something went wrong!!!")
}
