package main

import ("fmt")


func addition(n1 int,n2 int)int{
  sum:=n1+n2
  return sum
 }
 
 
func main(){
 var n1,n2 int
 fmt.Printf("ENter a digit")
 fmt.Scanf("%d",&n1)
 fmt.Printf("ENter another digit")
 fmt.Scanf("%d",&n2)
 sum1:=addition(n1,n2)
 fmt.Printf("The sum is %d",sum1)
 }
  
