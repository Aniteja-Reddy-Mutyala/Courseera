package main

import "fmt"

func main(){
	var phoneNumber=[]uint64{5554123,6665534,7684444,9897654};
	var name=[]string{"irapohl","randPill","bettyDodo","danielEuclid"};
	phone :=make(map[string]uint64);
	var nameInput string;
	for i:=0; i<len(name);i++{
		phone[name[i]]=phoneNumber[i];
	}
	fmt.Println("Enter the name of the person");
	fmt.Scanf("%s\n",&nameInput);
	fmt.Println("The phone number of ",nameInput, " is ",phone[nameInput]);

}