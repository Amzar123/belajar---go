func canConstruct(ransomNote string, magazine string) bool {
    
    for i := 0, i < len(ransomNote); i++ {
        for j:=0; j<len(magazine); j++ {
            if ransomNote[i] == magazine[j] {
                magazine = RemoveIndex(magazine, j)
                j = len(magazine)
            } else {
                return false
            }
        }
    }
    return true
}