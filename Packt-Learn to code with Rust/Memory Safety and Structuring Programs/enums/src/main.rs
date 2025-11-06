#[derive(Debug)]
enum OperatingSystem{
    Windows,
    MacOS,
    Linux
}
fn main(){
    let my_os=OperatingSystem::MacOS;
    let age_os= years_since_release(my_os);
    println!("Age is {age_os}");
    let dad_os=OperatingSystem::Windows;
    let age_dad_os=years_since_release(dad_os);
    println!("My dad's computer is {age_dad_os} years.");
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