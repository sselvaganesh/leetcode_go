package easy

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
            if (i+1) < len(s) {
                    if nxtVal=getIntVal(rune(s[i+1])); curVal >= nxtVal {
                      result = result+curVal
                    } else {
                        result = result-curVal+nxtVal
                        i=i+1
                    }
            } else {
                result = result + curVal
            }
    }

    return result
}


