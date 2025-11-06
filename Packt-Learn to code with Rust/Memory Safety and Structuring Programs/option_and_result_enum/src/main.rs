fn main() {
   let musical_instruments=[
    String::from("Guitar"),
    String::from("Drums"),
    String::from("bass"),

   ];
   let bass=musical_instruments.get(2);
   println!("{bass:?}");
   let valid_instrument=bass.unwrap();
   println!("{valid_instrument}");
   let invalid_instrument=musical_instruments.get(3);
   //invalid_instrument.unwrap();
   invalid_instrument.expect("Index out of bounds");
   
}
