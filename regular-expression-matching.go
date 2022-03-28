func isMatch(s string, p string) bool {
    p=preprocess(p)
    return match(s,p)
}

func match(s string,p string) bool{
    j:=0  //j is the index of p

    for i:=0;i<len(s);i++{
        v:= s[i]
        if j == len(p) {
            return false
        }
        switch byte(p[j]){
            case byte('.'):
                if j< len(p)-1&& p[j+1] == byte('*'){
                    i--
                }
                j++
                continue
            case byte(v):
                if j< len(p)-1&& p[j+1] == byte('*'){
                    i--
                }
                j++
                continue
            case byte('*'):
                // if j == len(p) - 1{
                //     continue
                // }

            if p[j-1] != byte('.') && p[j-1] != byte(v){
                i--
                j++
                continue
            }
            if i<len(s)  && j<len(p) -1 && (v == p[j+1] || p[j+1] == byte('.')){

                    if match(s[i:len(s)],p[j+1:len(p)]){
                        return true
                    }

            }
            
            if i<len(s) && j<len(p) -2 && p[j+2]== byte('*'){
                    if match(s[i:len(s)],p[j+1:len(p)]){
                        return true
                    }
            }

                
            continue
            default:
                if j <len(p)-1 && p[j+1] ==byte('*'){
                    j=j+2
                    i--
                    continue
                } 
                //fmt.Printf("%s",p[0:j])
                return false
        }
    }


    if equal_to_zero(p[j:len(p)]){
        return true
    }

    if j == len(p) -1 && p[j] == byte('*'){
        return true
    }

    if j==len(p)  {
        return true
    }else{
        return false
    }
}

func equal_to_zero(p string)bool{
    if len(p) == 0{
        return true
    }
    if p[0] == byte('*'){
        return equal_to_zero(p[1:len(p)])
    }
    for len(p) >= 2 && p[1] == byte('*'){
        if len(p) == 2{
            return true
        }
        p = p[2:len(p)]
    }
    return false
}



func preprocess(p string)string{
    tmp:=""
    for i:=0;i<len(p);i++{
        if i<len(p)-2{
            if p[i+1] == byte('*') && p[i] == p[i+2]{
                if i<len(p)-3{
                    if p[i+3] == byte('*'){
                        tmp = tmp + string(p[i]) + string(p[i+1])
                        i = i+3
                        continue
                    }
                }else{
                    tmp = tmp + string(p[i])+string(p[i]) + string(p[i+1])
                    i = i+2
                    continue
                }
            }
        }
        tmp = tmp + string(p[i])
    }
    return tmp

}
