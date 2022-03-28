func lengthOfLongestSubstring(s string) int {
    count:=0
    database :=make(map[byte]int)
    tmp_index := 0
    var i int
    for i=0;i<len(s);i++{
        if v,ok := database[s[i]];ok{
            if v < tmp_index{
                 database[s[i]] = i
                 continue
            }
            if i-tmp_index > count{
                count = i- tmp_index
            }
            tmp_index = v+1
            database[s[i]] = i
        }else{
            database[s[i]] = i
        }
    }

    if i-tmp_index > count{
        count = i- tmp_index
    }
    return count
}
