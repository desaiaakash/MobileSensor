package loadbalancer

//import "fmt"

var slaveDns = map[int]map[string]interface{}{
0 : {"connectstring": "127.0.0.1:3501", "weight": 2},
1 : {"connectstring": "127.0.0.1:3502", "weight": 4},
}


var i int = -1 
var cw int = 0 
var gcd int = 2 

func GetDns() string {
	//fmt.Println("GetDns Function called successfully")
	for {
		i = (i + 1) % len(slaveDns)
		if i == 0 {
			cw = cw - gcd
			//fmt.Println(gcd,cw,i)
			if cw <= 0 {
				//	fmt.Println(gcd,cw,i)
				cw = getMaxWeight()
				if cw == 0 {
					return ""
				}
			}
		}

		if weight, _ := slaveDns[i]["weight"].(int); weight >= cw {
			return slaveDns[i]["connectstring"].(string)
		}
	}
}

func getMaxWeight() int {
	max := 0
	for _, v := range slaveDns {
		if weight, _ := v["weight"].(int); weight >= max {
			max = weight
		}
	}

	return max
}

