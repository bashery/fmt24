// formating any PM / AM time format  to 24 format
// this package not very usefull,
// You can use this Package under any you like

package hours24


import (
    "fmt"
    "strconv"
    "strings"
)


// @Format acceptin time as string and return it as string
func Format(s string) string {

    add := 0

    if string(s[8]) == "P" {
        add = 12
    }

    hm := []string{string(s[0]), string(s[1])}

    hor1 := string(s[0])
    hor2 := string(s[1])
    hm = []string{hor1, hor2}
    strhor := strings.Join(hm, "")
    if strhor == "12" {
        strhor = "00"
    }
    numhor, _ := strconv.Atoi(strhor)

    fmt.Printf("%v, %v\n", strhor, numhor)

    time := []string{}


    numhor += add
    strhor = strconv.Itoa(numhor)


    if len(strhor) < 2 {
        strhor = "0"+strhor
    }
    time = append(time, strhor)
    for _,v := range(s[2:8]) {
        time = append(time, string(v))
        fmt.Println(time)
    }
    resault := strings.Join(time, "")
    return resault
}








