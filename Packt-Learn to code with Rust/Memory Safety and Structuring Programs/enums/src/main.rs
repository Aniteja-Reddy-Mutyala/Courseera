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
fn wash_laundry(cycle:LaundryCycle){
  match cycle{
    LaundryCycle::Cold => {println!("Running the laundry in cold mode")},
    LaundryCycle::Hot { temperature }=>{ println!(" Runnig the laundry with temperature of {temperature}");},
    LaundryCycle::Delicate(fabric_type)=>{println!("Running the laundry with delicate cycle for {fabric_type}");}
  }
}
fn main(){
    let my_os=OperatingSystem::MacOS;
    let age_os= years_since_release(my_os);
    println!("Age is {age_os}");
    let dad_os=OperatingSystem::Windows;
    let age_dad_os=years_since_release(dad_os);
    println!("My dad's computer is {age_dad_os} years.");
    wash_laundry(LaundryCycle::Cold);
    wash_laundry(LaundryCycle::Hot { temperature:55 });
    wash_laundry(LaundryCycle::Delicate(String::from("Silk")));

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