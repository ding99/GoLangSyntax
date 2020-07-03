package main

import . "fmt"

type Phone interface{ call() }

type NokiaPhone struct {}
func (nokiaPhone NokiaPhone) call(){ Println("I am Nokia, I can call you!") }

type IPhone struct {}
func (iPhone IPhone) call() { Println("I am iPhone, I can call you!") }

func calling(){
  Println("-- Phones")
  var phone Phone
  
  phone = new(NokiaPhone)
  phone.call()
  
  phone = new(IPhone)
  phone.call()
}


type Person interface{ name() string; age() int }

type Woman struct{}
func (woman Woman) name() string{ return "Julia" }
func (woman Woman) age() int { return 23 }

type Man struct {}
func (man Man) name() string { return "Jason" }
func (man Man) age() int { return 27 }

func persons(){
  Println("-- Person")
  
  var person Person;
  
  person = new(Woman);
  Println(person.name())
  Println(person.age())
  
  person = new(Man)
  Println(person.name())
  Println(person.age())
}

type IFPhone interface{ call01(); call02() }

type Phone1 struct { id int; name string; category_id int; category_name string }
func (test Phone1) call01(){ Println("First class, First method ", Phone1{id:1,name:"laugh"}) }
func (test Phone1) call02(){ Println("First class, Second method ", Phone1{1,"laugh",4,"names" }) }

type Phone2 struct { member_id int; member_balance float32; member_sex bool; member_nick string }
func (test Phone2) call01(){ Println("Second class, First method ", Phone2 {22,15.23,false,"laugh 18" }) }
func (test Phone2) call02(){ Println("Second class, Second method ", Phone2{44,100,true,"steven" }) }

func calls(){
  Println("-- phones")
  var phone IFPhone
  phone = new(Phone1); phone.call01(); phone.call02()
  phone = new(Phone2); phone.call01(); phone.call02()
}

type PhoneCell interface { call() string }

type AndroidCell struct { brand string }
type IPhoneCell struct { version string }

func (android AndroidCell) call() string { return "I am Android " + android.brand }
func (iphone IPhoneCell) call() string { return "I am IPhone " + iphone.version }
func printCall(p PhoneCell) { Println(p.call() + ", I can call you!") }

func paras(){
  Println("-- parameters")
  var vivo = AndroidCell{ brand:"Vivo" }
  var hw = AndroidCell{ "HuaWei" }
  i7 := IPhoneCell{ "7 Plus" }
  ix := IPhoneCell{ "X" }
  
  printCall(vivo); printCall(hw); printCall(i7); printCall(ix)
}

type fruit interface { getName() string; setName(name string) }
type apple struct { name string }
func (a *apple) getName() string { return a.name }
func (a *apple) setName(name string){ a.name = name }

func pointers(){
  Println("-- pointers")
  
  a := apple{"Fuji"}; Println(a.getName())
  a.setName("Gala"); Println(a.getName())
}

func main(){
  calling()
  persons()
  calls()
  paras()
  pointers()
}
