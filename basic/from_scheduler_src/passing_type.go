package main

type TypeBridgePassing struct {

}




type FuncType func(ip string, meta interface{}) TypeBridgePassing


func CalculateFunc(ip string, meta interface{}) TypeBridgePassing {
	return TypeBridgePassing{
	}
}




func MainCallFunc(mapFunction FuncType){
}

func main() {
	MainCallFunc(CalculateFunc)
}