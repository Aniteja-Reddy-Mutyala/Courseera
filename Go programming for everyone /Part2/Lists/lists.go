package main
import "fmt"

type ListElement struct{
	data int
	next *ListElement
}
func createListElement(d int ,ptr *ListElement) ListElement{
	var element ListElement
	element.data=d;
	element.next=ptr;
	return element;
}
func (h *ListElement) PrintList(){
	for h!= nil{
		fmt.Println(h.data,"->");
	     h=h.next;
	}
	if h== nil{
		fmt.Println("#####");
		return;
	}
}
func FillList(dataSilce []int,h **ListElement){
	curEl:=ListElement{dataSilce[0],nil};
	fmt.Println("curEl:" ,curEl);
	*h=&curEl;
	ptrCur:=&curEl;
	for i:=1;i<len(dataSilce);i++{
		nextElem:=ListElement{dataSilce[i],nil};
		fmt.Println("Next",nextElem);
        ptrCur.next=&nextElem;
		fmt.Println("curE1",curEl);
		ptrCur=&nextElem; 
	}


}
func main(){
     var head *ListElement;
	 var e,f ListElement;
	 xD:=[]int {4,8,16,32,64,128}
	 fmt.Println("Create a list from :",xD);
	 head=&e;
	 e.data=5;
	 e.next=&f;
	 f.data=6;
	 fmt.Println("Element e is :",e,"head is :" ,head);
	 fmt.Println("Element f is : ",e.next,"head is :",head);
	 f.next=new(ListElement);
	 f.next.data=24;
	 fmt.Println("element f.next is :",f.next,"head is :",head);
	 head.PrintList();
	 FillList(xD,&head);
	 fmt.Println("Head points to : ",head);
	 head.PrintList();

}
