fn main() {
    open_store("Monroe");
}
fn open_store(neighbour:&str){
    println!("Welcome to my store in {neighbour}");
    about_store();
}
fn about_store(){
    println!("This is my store");
    thank_you();
}
fn thank_you(){
    println!("Thank you for visiting my store");
}