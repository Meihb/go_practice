package main
import "fmt"

func main(){
	arr := []int{1,2,3,4,5,6,7}
	Loop:
		for k,v := range arr{
			fmt.Printf("k is %d,v is %d\r\n",k,v)
			// if v<4 {
			// 	break;
			// }

			switch{
			case v<4:
				fmt.Println("in switch")
				if v==3 {
					break
				} 
			}
			if k+v ==9 {
				fmt.Println("should get out the loop")
				break Loop
			}
		}
}