//use std::fs::File;
use std::fs;
use std::io::{Error, Read, stdin};
//use std::process;

fn main() {
    let file_result = read_file();
    match file_result {
        Result::Ok(contents) => {
            println!("{contents}");
        }
        Result::Err(error) => {
            eprintln!("There was a error.The error is {error}");
        }
    }
}
fn read_file() -> Result<String,Error> {
    println!("Please enter the name of file");
    let mut input = String::new();
    stdin().read_line(& mut input)?;
   /*  match stdin().read_line(&mut input) {
        Result::Ok(_) => {}
        Result::Err(error) => {
            return Err(error);
        }
    }*/
    fs::read_to_string(&input.trim())
    //let mut file=File::open(input.trim())?;
    /*let mut file: File = match File::open(input.trim()) {
        Result::Ok(file) => file,
        Result::Err(error) => {
            return Result::Err(error);
        }
    };*/
    //let mut file_content = String::new();
    //file.read_to_string(&mut file_content)?;
    /*match read_operation {
        Result::Ok(_) => {}
        Result::Err(error) => {
            return Result::Err(error);
        }
    }*/
    //Result::Ok(file_content)
}
