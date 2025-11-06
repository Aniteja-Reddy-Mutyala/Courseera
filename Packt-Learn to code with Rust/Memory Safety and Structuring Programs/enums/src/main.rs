#[derive(Debug)]
enum OperatingSystem{
    Windows,
    MacOS,
    Linux
}
#[derive(Debug)]
enum LaundryCycle{
    Cold,
    Hot{
        temperature:u32

    },
    Delicate(String)
}
impl LaundryCycle{
fn wash_laundry(self:&Self){
  match self{
    LaundryCycle::Cold => {println!("Running the laundry in cold mode")},
    LaundryCycle::Hot { temperature }=>{ println!(" Runnig the laundry with temperature of {temperature}");},
    LaundryCycle::Delicate(fabric_type)=>{println!("Running the laundry with delicate cycle for {fabric_type}");}
  }
}
}
#[derive(Debug)]
enum OrderOnlineStatus{
    Ordered,
    Packed,
    Shipped,
    Delivered
}
impl OrderOnlineStatus{
   fn check(self:&Self){
    match self{
       /* OrderOnlineStatus::Ordered | OrderOnlineStatus::Packed =>{
            println!("Your item is being prepped for shipment");
        }*/
        OrderOnlineStatus::Delivered=>{println!("Your item has been delivered");}
        order_status=>{println!("Your item is {order_status:?}");}
    }
   }
}
enum Milk{
    LowFat(i32),
    WholeMilk
}
impl Milk{
    fn drink(self){
        match self{
            Milk::LowFat(2) =>{println!("Delicious!!!2% milk is excellent");},
            Milk::WholeMilk =>{println!("Whole milk");},
            Milk::LowFat(percent)=>{println!("You got lowfat {percent}% milk");}
        }
    }
}
fn main(){
    let my_os=OperatingSystem::MacOS;
    let age_os= years_since_release(my_os);
    println!("Age is {age_os}");
    let dad_os=OperatingSystem::Windows;
    let age_dad_os=years_since_release(dad_os);
    println!("My dad's computer is {age_dad_os} years.");
    let wash_cold_clothes=LaundryCycle::Cold;
    wash_cold_clothes.wash_laundry();
    let wash_hot_mode=LaundryCycle::Hot { temperature:55 };
    wash_hot_mode.wash_laundry();
    let wash_delicate=LaundryCycle::Delicate(String::from("Silk"));
    wash_delicate.wash_laundry();
    // wash_laundry(LaundryCycle::Cold);
    //wash_laundry(LaundryCycle::Hot { temperature:55 });
    //wash_laundry(LaundryCycle::Delicate(String::from("Silk")));
    let item=OrderOnlineStatus::Ordered;
    item.check();
    let shipped_item=OrderOnlineStatus::Shipped;
    shipped_item.check();
    let delivered_item=OrderOnlineStatus::Delivered;
    delivered_item.check();
    let whole_milk=Milk::WholeMilk;
    whole_milk.drink();
    let two_percent=Milk::LowFat(2);
    two_percent.drink();
    let low_fat=Milk::LowFat(10);
    low_fat.drink();

}
fn years_since_release(os:OperatingSystem)->u32{
    match os {
        OperatingSystem::Windows =>{
           println!("Very old system");
           39
        } ,
        OperatingSystem::MacOS =>23,
        OperatingSystem::Linux => 34
    }
   
}