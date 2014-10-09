package main

import "bufio"
import "fmt"
import "math"
import "os"
import "sort"
import "strings"
import "strconv"

const (
    RadToDeg  = 180 / math.Pi
    DegToRad  = math.Pi / 180
    RadToGrad = 200 / math.Pi
    GradToDeg = math.Pi / 200
)

type defibrillator struct {
    name      string
    longitude float64
    latitude  float64
}

type defibrillators []*defibrillator

type byDistance struct {
    ds defibrillators
    longitude float64
    latitude  float64
}

func (d byDistance) Len() int {
    return len(d.ds)
}
func (d byDistance) Swap(i, j int) {
    d.ds[i], d.ds[j] = d.ds[j], d.ds[i]
}
func (d byDistance) Less(i, j int) bool {
    var d0, d1 = d.ds[i], d.ds[j]
    return distance(d.longitude, d.latitude, d0.longitude, d0.latitude) <
    distance(d.longitude, d.latitude, d1.longitude, d1.latitude)
}

func distance(lon1, lat1, lon2, lat2 float64) float64 {
    var x = (lon2 - lon1) * math.Cos((lat1+lat2)/2)
    var y = lat2 - lat1
    return math.Sqrt(math.Pow(x, 2)+math.Pow(y, 2)) * 6371
}

func main() {
    var in = bufio.NewReader(os.Stdin)

    var line = readline(in)
    line = strings.Replace(line, ",", ".", -1)
    var longitude, _ = strconv.ParseFloat(line, 64)
    longitude *= DegToRad

    line = readline(in)
    line = strings.Replace(line, ",", ".", -1)
    var latitude, _ = strconv.ParseFloat(line, 64)
    latitude *= DegToRad

    line = readline(in)
    var n, _ = strconv.Atoi(line)

    var ds = make(defibrillators, n)

    for i := 0; i < n; i++ {
        line = readline(in)
        var split = strings.Split(line, ";")

        ds[i] = new(defibrillator)
        ds[i].name = split[1]

        split[4] = strings.Replace(split[4], ",", ".", -1)
        var longitude, _ = strconv.ParseFloat(split[4], 64)
        ds[i].longitude = longitude * DegToRad

        split[5] = strings.Replace(split[5], ",", ".", -1)
        var latitude, _ = strconv.ParseFloat(split[5], 64)
        ds[i].latitude = latitude * DegToRad
    }

    sort.Sort(byDistance{ds, longitude, latitude})
    fmt.Println(ds[0].name)
}

func readline(in *bufio.Reader) string {
    var line, _ = in.ReadString('\n')
    line = strings.TrimSpace(line)
    return line
}
