package util

func ArrayContains(array []string, term string) bool {
    for _, a := range array {
        if a == term {
            return true
        }
    }
    return false
}
