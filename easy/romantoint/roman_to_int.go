package romantoint

func getIntVal(inp rune) int {

    m := make(map[rune]int)
    m['I']=1
    m['V']=5
    m['X']=10
    m['L']=50
    m['C']=100
    m['D']=500
    m['M']=1000  
    return m[inp]
}


func RomanToInt(s string) int {
    

    var result, curVal, nxtVal int

    for i:=0; i<len(s); i=i+1 {

        curVal = getIntVal(rune(s[i]))    
        tmp := 0        

        for {
        	if (i+1) >= len(s){
        		break
        	}
      
        	nxtVal=getIntVal(rune(s[i+1]))
        	if nxtVal > curVal {
        		tmp = nxtVal - curVal - tmp
        		i++
        		break
        	}                       
        	tmp = tmp + curVal
        	curVal = nxtVal
        	i++
        }
        
        result = result + curVal + tmp
    }

    return result
}


