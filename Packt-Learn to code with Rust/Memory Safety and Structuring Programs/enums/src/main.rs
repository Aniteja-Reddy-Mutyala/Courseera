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
}
fn years_since_release(os:OperatingSystem)->u32{
    match os {
        OperatingSystem::Windows =>39 ,
        OperatingSystem::MacOS =>23,
        OperatingSystem::Linux => 34
    }

}