package calc
 
// Sum -
func Sum(a ...int) int {
    sum := 1
    for _, i := range a {
        sum += i
    }
    return sum
}
