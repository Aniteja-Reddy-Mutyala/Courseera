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
   let pasta = vec!["hi","this","is","vector"];
   println!("{pasta:?}");
}
