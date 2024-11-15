// iota is a const number  have 0 as a default value and he increments  by +1

package main

import (
	// "bufio"
"fmt"
	// "log"
	"os"
	// "time"
)

//"reflect"
//"strings"

//	"workplace/pkj"  ==> get a file packages
//"strconv" =>strconv.parseFloat  === to change string to float



func main(){

///////////////////////////////////////////////////////////
//var x = map[string]string{"name":"reda","age":"5"}  
 
/////////////////////////////////////////////////////////////
//fmt.Println(pachages.Total(5,5,5,44,2,4))
//fmt.Println(1%2)
//fmt.Printf("%d",Total,"%t","totaaaaal ajmiii ")


////////////////////////////////////////////////////////////////////
//for key,value:=range x{fmt.Printf( "%s\n%s\n",key,value)}
/////////////////////////////////////////////////////////////////////

/*
var array =[...]string{"a", "b", "c"}
for i := 0; i < len(array); i++ {
	if (i == 2) {
		array[i] ="redaaaaaaaaaaaaaaa"
	}
	fmt.Println(array[i])
	
}
*/
///////////////////////////////////////////////////////////////////////////
/*
var array =[3][3]int {{1,2,3},{4,5,6},{7,8,9}}
var WHITEONE[]int
var NEGGERSONE[]int

for i:=0;i<len(array);i++{


for j:=0;j<len(array[i]);j++{
	if(array[i][j]%2 == 0 ){



		WHITEONE = append(WHITEONE,  array[i][j])
	}else{
		NEGGERSONE = append(NEGGERSONE,  array[i][j])
		}
}
}
fmt.Println(pos,neg)
*/
///////////////////////////////////////////////////////////////////////
/*
var name = " redaa"
var array = [...]string{"redaa", "klhkk"}


fmt.Println(reflect.TypeOf(name))
fmt.Println(len(name))
fmt.Println(strings.ToUpper(name))
fmt.Println(strings.ToLower(name))
fmt.Println(strings.HasPrefix(name,"re"))
fmt.Println(strings.HasSuffix(name,"aa"))
fmt.Println(strings.Join(array[:]," °"))
fmt.Println(strings.Repeat(name,5))
fmt.Println(strings.Contains(name,"ed"))
fmt.Println(strings.Index(name,"ed"))
fmt.Println(strings.Count(name,"a"))
fmt.Println(strings.Replace(name,"a","g",1))
fmt.Println(strings.TrimSpace(name))
fmt.Println(strings.ContainsAny(name,"e"))



*/
///////////////////////////////////////////////////////////////////////

/*
type persone struct{
	name string
	age int 
	department string
}

var ahmed = persone{name:"reda",age:45,department: "ggo "}
fmt.Println(ahmed.age)
*/


///////////////////////////////////////////////////////////////////////
/*
file,err :=os.Open("file.txt")

if (err !=nil){
	log.Fatal(err)


}
defer file.Close()
scanner :=bufio.NewScanner(file)
fmt.Println(scanner)

for scanner.Scan() {

	line :=scanner.Text()
fmt.Println(line)
time.Sleep(1*time.Second)
}
*/
/*
file:="file.txt"
		if (checkFile(file)==true) {
			fmt.Println(ReadFile(file))
		}else{
			fmt.Println("this is a wrong file ")
		}
}

func checkFile(file string) bool {
	_,err:=os.Stat(file)
	if os.IsNotExist(err){
		return false
	}
	return err==nil


}
func ReadFile(file string)  string {


	content,err :=os.ReadFile(file)
	if (err != nil){
		fmt.Println(err)
	}
	contentt :=fmt.Sprintf("%s",content)
	return contentt
	*/
///////////////////////////////////////////////////////////////////////////////////
/*
filee := "test.txt"
content :="hello my name is reda and dddddddddddddddi have 20 years old "


file ,err:= os.Create(filee)
if (err !=nil) {
panic(err)
}
defer file.Close()

_,err = file.WriteString(content)
if (err !=nil) {
	panic(err)
	}
err =file.Sync()


*/
/////////////////////////////////////////////////////////////////////
str1:= os.Args[1]
str2:= os.Args[2]

var t []int

res := ""

for i := 0; i < len(str1); i++ {
	for j := 0; j < len(str2); j++ {
		
if   string(str1[i])  ==  string(str1[j]) {
	res +=string(str1[i])

}


	}
}


fmt.Println(res)



}