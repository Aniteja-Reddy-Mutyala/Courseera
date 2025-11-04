#[derive(Debug)]
struct Coffee {
    price: f64,
    name: String,
    is_hot: bool,
}
#[derive(Debug)]
struct TaylorSwiftSong {
    title:String,
    release_year:u32,
    duration_secs:u32
}

impl TaylorSwiftSong{
    fn new(title:String,release_year:u32,duration_secs:u32)-> Self{
        TaylorSwiftSong { title, release_year, duration_secs }
    }
    fn display_song_info(&self){
           println!("The name of song is {}",self.title);
           println!("Years since release {}",self.years_since_release());
           println!("The duration is {} seconds",self.duration_secs);
    }
    fn double_length(&mut self){
          self.duration_secs=2 *self.duration_secs;
          
    }
    fn is_longer_than(&self,other:&Self) -> bool{
             self.duration_secs > other.duration_secs
    }
    fn years_since_release(&self) ->u32{
        2025-self.release_year
    }
}
fn main() {
    let mut mocha: Coffee = Coffee {
        price: 4.99,
        name: String::from("Mocha"),
        is_hot: false,
    };
    let song=TaylorSwiftSong{
        title:String::from("Blank space"),
        release_year:2014,
        duration_secs:231
    };
    song.display_song_info();
    let  mut song1 = TaylorSwiftSong{
        title:String::from("Blank space"),
        release_year:2014,
        duration_secs:120
    };
    let one :bool=song.is_longer_than(&song1);
    song1.double_length();
    let two: bool=song.is_longer_than(&song1);
    println!("{one},{two}");

    let song2 =TaylorSwiftSong::new(String::from("Shake it off ") ,2015,210);
    song2.display_song_info();
    let latte = make_coffee(String::from("Latte"), 3.99, true);
    println!(
        "My morning {} costs {} and it is {} that it was hot.",
        mocha.name, mocha.price, mocha.is_hot
    );
    mocha.price = 6.00;
    println!("The latest price of {} is {} ", mocha.name, mocha.price);
    println!(
        "The latest price of {} is {} and it is {} that it is hot.",
        latte.name, latte.price, latte.is_hot
    );
    let caramel_macchiato: Coffee = Coffee {
        name: String::from("Caramel Macchiato"),
        ..mocha
    };
    println!(
        "The lastest price of {} is {} and it is {} that it is hot",
        caramel_macchiato.name, caramel_macchiato.price, caramel_macchiato.is_hot
    );
    drink_coffee(& mut mocha);
    println!("{:?}",mocha);
}
fn make_coffee(name: String, price: f64, is_hot: bool) -> Coffee {
    Coffee {
        price: price,
        name: name,
        is_hot: is_hot,
    }
}
/*fn drink_coffee(coffee: Coffee) {
    println!("Drinking my delicious {}", coffee.name);
}*/
fn drink_coffee(coffee:& mut Coffee){
    println!("Drinking my delicious {}", (*coffee).name);

}
