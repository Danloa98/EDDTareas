package main

import (
	"fmt"
)

type Student struct{
	id int
	name string
	carrera string
}


type Node struct{
	 next *Node
	 student Student

}

func NewNode(st Student) *Node {
	return &Node{
		next: nil, //In go, the equivalent for NULL seems to be the 'nil' expression
		student: st,
	}
}

type LinkedList struct{
	head *Node
	tail *Node



}

func (l *LinkedList) isEmpty() bool{

	if l.head==nil{
		return true
	} else{
		return false}


}

func (l *LinkedList) addElement (st Student) {
	//In Go the method of a Struct has to be written outside of it
	//but we use the part func (l *LinkedList) to state that we are creating a method fot the LinkedList struct
	//Note that we put *LinkedList as a pointer so we indicate that we are going to modify directly the original
	//struct created since we are pointing to it trough its memory address, If we would've put it without the *
	//we would be working on a 'copy' of it instead of working on the actual receiver/struct
	//temp:= l.head

	newnode:= &Node{ student: st}


	if l.isEmpty(){
		l.head= newnode
		l.tail= newnode
		//newnode.next=nil
		//fmt.Println("Student ",newnode.student.name, "Added to the list!" )

	}else {

		temp := l.tail
		temp.next = newnode
		newnode.next = nil
		l.tail = newnode
		//fmt.Println("Student ",newnode.student.name, "Added to the list!" )
	}
}

func (l *LinkedList) deletebyID(id int){
	current:= l.head//We create a temp Node which is going to start from the beginning of our list
	prev:= l.head
	if(!l.isEmpty()){

		if current.student.id== id{//if the element we are looking happens to be at the beginning
			l.deleteFirst()
		}else if l.tail.student.id==id{//if element we are looking to remove happens to be at the end
			l.deleteLast()
		}else{

			for current!=nil{
			if(current.student.id==id){
				prev.next=current.next
				current.next=nil
				return
			}
			prev=current
			current=current.next
			}

			fmt.Println("Element was not found, looks like it is not in this list of elements")



		}


	}else{
		fmt.Println("List does not have eny elements, try adding some :)")
	}




}

func (l *LinkedList) deleteFirst(){
if !l.isEmpty(){
	if(l.head==l.tail){//if there is only 1 element in the list
		l.head=nil
		l.tail=nil

	}else {
		l.head=l.head.next
	}


}else {
	fmt.Println("Empty list, try adding some elements")
}


}

func (l *LinkedList) deleteLast(){
	if !l.isEmpty(){
		if(l.head==l.tail){//if there is only 1 element in the list
			l.head=nil
			l.tail=nil

		}else {
			prev:= l.head
			current:=l.head

			for current!=nil{
				if(current==l.tail){
					l.tail=prev
					prev.next=nil
				}
				prev=current
				current=current.next

			}



		}


	}else {
		fmt.Println("Empty list, try adding some elements")
	}


}

func (l LinkedList) printElements(){//Here we did not put any pointer since the main objective is just to print the data
	//we dont want to make changes to the receiver, we just want to print out the information we have stored in it
	if(l.isEmpty()){

		fmt.Println("No more elements in the list, try adding some :)")
	}else{ 	temp := l.head
		fmt.Println("Nombre, ID, Carrera")

		for temp!=l.tail.next{//While cycle implementation on GO
			fmt.Println(temp.student.name,"  ", temp.student.id,"  ", temp.student.carrera)

			temp=temp.next
		}}


}

func main(){

		stu1:= Student{
			id: 5,
			name: "James",
			carrera: "Industrial",
		}

	stu2:= Student{
		id: 4,
		name: "Carlos",
		carrera: "Sistemas",
	}
	stu3:= Student{
		id: 3,
		name: "Lucia",
		carrera: "Electronica",
	}

	stu4:= Student{
		id: 2,
		name: "Maria",
		carrera: "Ambiental",
	}

		list:= &LinkedList{}
		list.addElement(stu1)
		list.addElement(stu2)
		list.addElement(stu3)
		list.addElement(stu4)
		list.deletebyID(stu2.id)
		list.printElements()


}
