package main

import("fmt")
type Report struct{
	Content string
	Title string
}
func(report Report)GetContent(){
	fmt.Println("The content is:",report.Content,"\n")
}
func(report Report)SendEmail(email string){
	fmt.Println("Email sends to :",email)
}
func(report Report )Save(){
	fmt.Println("saved")
}
func main(){
  report1:=Report{
      Content: "hi yash",
	  Title: "Annual report",
  }
  report1.GetContent()
  report1.Save()
  report1.SendEmail("yashgoyal322023@gmail.com")
}