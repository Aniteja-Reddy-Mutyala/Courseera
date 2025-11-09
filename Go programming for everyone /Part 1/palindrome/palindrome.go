package main

import ("fmt"
		"strings"
)

func check_palindrome(isPalindrome *string) bool{
	start:=0
	end:=len(*isPalindrome)-1
	for start < end{
     if (*isPalindrome)[start]!=(*isPalindrome)[end]{
        return false
	 }else{
		start +=1
		end -=1
	 }
	}
	return true
    
}
func main(){
	var isPalindrome string
	fmt.Println("Enter your word");
	fmt.Scanf("%s\n",&isPalindrome);
	isPalindromeTest:=strings.ToLower(isPalindrome);
	
	fmt.Println(check_palindrome(&isPalindromeTest));
}
