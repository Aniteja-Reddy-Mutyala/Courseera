fn main() {
   let mut pizza_diameters:Vec<i32>=vec![8,10,12,14];
   println!("{pizza_diameters:?}");
   pizza_diameters.push(16);
   pizza_diameters.push(18);
   println!("{pizza_diameters:?}");
   pizza_diameters.insert(0, 4);
   println!("{pizza_diameters:?}");
   pizza_diameters.pop();
   println!("{pizza_diameters:?}");
   pizza_diameters.remove(3);
     println!("{pizza_diameters:?}");
   let pasta = vec!["hi".to_string(),"this".to_string(),"is".to_string(),"vector".to_string()];
   let _hi=&pasta[0];
   let is=pasta.get(4);
   println!("{pasta:?}");
//    println!("{is:?}");
   match is{
    Option::Some(value)=>{println!("the value is {value}");}
    Option::None=>{println!("Error");}
   }
}
