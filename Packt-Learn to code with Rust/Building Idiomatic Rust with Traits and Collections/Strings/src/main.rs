

fn main() {
  //let  first_name:String=String::from("Aniteja Reddy");
  //let  last_name= String::from("Mutyala");
  //let full_name = first_name + &last_name;
  //full_name.push_str(last_name);
  //let icon= format!("{first_name} {last_name}");
  //println!("{icon}");
  //println!("{first_name}");
  //println!("{last_name}");
  let music_genre= "   rock ,metal,rap,pop";
  println!("{}",music_genre.trim());
  println!("{}",music_genre.trim().to_uppercase());
  println!("{}",music_genre.replace("a", "@"));
  let genres:Vec<&str>=music_genre.split(',').collect();
  println!("{:?}",genres);
}
