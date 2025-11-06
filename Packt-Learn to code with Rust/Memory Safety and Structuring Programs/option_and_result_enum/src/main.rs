fn main() {
  
   let musical_instruments=[
    String::from("Guitar"),
    String::from("Drums"),
    String::from("bass"),

   ];
   let bass=musical_instruments.get(3);
   println!("{bass:?}");
   match bass {
      
      Option::Some(instrument)=>{println!("The instrument is {instrument}");}
      Option::None=>{println!("Index out of bounds");}
   }
   //let valid_instrument=bass.unwrap();
   //println!("{valid_instrument}");
   let invalid_instrument=musical_instruments.get(3);
   //invalid_instrument.unwrap();
   //invalid_instrument.expect("Index out of bounds");
   
}
