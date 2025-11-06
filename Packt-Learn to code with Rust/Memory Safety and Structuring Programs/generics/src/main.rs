fn identity<T>(value:T)->T{
   value
}
fn make_tuple<T,U>(first:T,second:U)->(T,U){
    (first,second)
}
fn main() {
    println!("Hello, world!");
    let i32_value= identity::<u32>(5);
    println!("identity is {i32_value}");
    let bool_value =identity(true);
    println!("Identity is {bool_value}");
    let f64_value= identity::<f32>(9.056);
    println!("Identity is {f64_value}");
    let new_tuple=make_tuple(String::from("hello"),10);
    println!("{new_tuple:?}");

}