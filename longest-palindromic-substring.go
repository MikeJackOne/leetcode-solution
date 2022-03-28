func longestPalindrome(s string) string {
    var maxstring string
    var length int
    if len(s) <=0{
        return ""
    }
    maxstring = string(s[0])
    length = 1
    for i,v := range s{
        if len(s) - i < length/2{
            return maxstring
        }
        if i+1 < len(s) && v == rune(s[i+1]){
            tmp_result := compare(s[0:i+1],s[i+1:len(s)])
            //fmt.Printf("%s,%s\n",s[0:i+1],s[i+1:len(s)])
            if len(tmp_result) * 2 > length{
                length = len(tmp_result) * 2
                tmp_len := len(tmp_result)
                maxstring = s[i- tmp_len + 1:i+1] + tmp_result 
            }
        } 
        if i+1 < len(s) && i-1 >=0 && s[i+1] == s[i-1]{
            tmp_result := compare(s[0:i],s[i+1:len(s)])
            if len(tmp_result) * 2 + 1> length{
                length = len(tmp_result) * 2 + 1
                tmp_len := len(tmp_result)
                maxstring = s[i- tmp_len:i+1] + tmp_result 
            }
        }
    }
    return maxstring
}

func compare(s1 string,s2 string)string{
    result := ""
    max := min(len(s1),len(s2))
    s1_length := len(s1) - 1
    for i:=0 ;i>= -1 *(max-1);i--{ 
        if s1[s1_length + i] != s2[-1*i]{
            return result
        }else{
            result += string(rune(s2[-1*i]))
        }
    }
    return result
}

func min(a int,b int)int{
    if a > b{
        return b
    }
    return a
}
