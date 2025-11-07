/*#[derive(Debug,Copy,Clone)]
enum MyOption{
   Some(i32),
   None
}
impl MyOption{
   fn unwrap(self)->i32{
      match self{
         MyOption::Some(value)=>{value},
         MyOption::None=>{panic!("No!!!")}
      }

   }
   fn unwrap_or(self,default_value:i32)->i32{
      match self{
         MyOption::Some(value)=>{value},
         MyOption::None=>{default_value},
      }
   }
}
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
   let present_value=Option::Some(15);
   let missing_value:Option<i32>=Option::None;
   println!("The value is {}",present_value.unwrap_or(0));
   println!("The value is {}",missing_value.unwrap_or(0));
   let some_option=MyOption::Some(50);
   println!("{}",some_option.unwrap());
   let none_option:MyOption=MyOption::None;
   println!("{}",none_option.unwrap_or(0));
   
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
}*/
fn main(){
   let ok:Result<i32,&str> =Result::Ok(5);
   println!("{ok:?}");
   let disaster:Result<i32, &str>=Result::Err("Something went wrong");
   println!("{disaster:?}");

}
