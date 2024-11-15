package pachages

import f "fmt" 
func SayHello(){
	f.Println("hello brp ")
}


func hello(){
	f.Println("hello brp ")
}

func init(){
	hello()
}

func Sum(a int ,b int ) int{

	return a+b
}
func Total( x ... int )int{
	var t int = 0
	for i := 0; i < len((x)); i++ {
		t += x[i]
	}
	return t
}

func Max( x ... int )int{
	var m int =0

	for i := 0; i < len((x)); i++ {
		m=x[i+1]
	}
	return m		
}

