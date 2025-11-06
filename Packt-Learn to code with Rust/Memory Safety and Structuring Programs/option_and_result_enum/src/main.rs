fn main() {
  
   /*let musical_instruments=[
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
   //let invalid_instrument=musical_instruments.get(3);
   //invalid_instrument.unwrap();
   //invalid_instrument.expect("Index out of bounds");*/
   let availability=is_item_in_stock(true,false);
   println!("{availability:?}");
   match availability{
      Some(value)=>{println!("The item is available:{value}");}
      None=>{println!("The item doesn't exist in our system");}
   }
   
}
fn is_item_in_stock(item_is_in_system:bool,item_is_in_stock:bool)->Option<bool>{
  if item_is_in_system && item_is_in_stock{
    Option::Some(true)
  }
  else if item_is_in_system{
       Option::Some(false)
  }
  else{
     Option::None
  }
}
