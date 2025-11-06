fn main() {
   let musical_instruments=[
    String::from("Guitar"),
    String::from("Drums"),
    String::from("bass"),

   ];
   let bass=musical_instruments.get(2);
   println!("{bass:?}");
   
}
