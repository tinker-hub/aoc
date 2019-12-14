package main

import (
   "fmt"
   "strings"
   "strconv"
)

/*
[63] = 68926676
[63] = 1
(53)

*/

var input string = `3,8,1005,8,324,1106,0,11,0,0,0,104,1,104,0,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1001,8,0,29,1,1107,14,10,1006,0,63,1006,0,71,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,1,10,4,10,1002,8,1,61,1,103,18,10,1006,0,14,1,105,7,10,3,8,1002,8,-1,10,101,1,10,10,4,10,1008,8,1,10,4,10,101,0,8,94,1006,0,37,1006,0,55,2,1101,15,10,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,0,10,4,10,101,0,8,126,2,1006,12,10,3,8,102,-1,8,10,101,1,10,10,4,10,1008,8,1,10,4,10,1001,8,0,152,3,8,102,-1,8,10,1001,10,1,10,4,10,108,0,8,10,4,10,101,0,8,173,1006,0,51,1006,0,26,3,8,102,-1,8,10,101,1,10,10,4,10,1008,8,0,10,4,10,1001,8,0,202,2,8,18,10,1,103,19,10,1,1102,1,10,1006,0,85,3,8,102,-1,8,10,1001,10,1,10,4,10,108,0,8,10,4,10,1001,8,0,238,2,1002,8,10,1006,0,41,3,8,102,-1,8,10,1001,10,1,10,4,10,108,0,8,10,4,10,101,0,8,267,2,1108,17,10,2,105,11,10,1006,0,59,1006,0,90,3,8,1002,8,-1,10,1001,10,1,10,4,10,1008,8,1,10,4,10,1001,8,0,304,101,1,9,9,1007,9,993,10,1005,10,15,99,109,646,104,0,104,1,21102,936735777688,1,1,21101,341,0,0,1105,1,445,21101,0,937264173716,1,21101,352,0,0,1106,0,445,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,3,10,104,0,104,1,3,10,104,0,104,0,3,10,104,0,104,1,21101,3245513819,0,1,21102,1,399,0,1105,1,445,21102,1,29086470235,1,21102,410,1,0,1105,1,445,3,10,104,0,104,0,3,10,104,0,104,0,21101,825544712960,0,1,21102,1,433,0,1106,0,445,21102,825460826472,1,1,21101,0,444,0,1106,0,445,99,109,2,22102,1,-1,1,21101,0,40,2,21101,0,476,3,21102,466,1,0,1105,1,509,109,-2,2105,1,0,0,1,0,0,1,109,2,3,10,204,-1,1001,471,472,487,4,0,1001,471,1,471,108,4,471,10,1006,10,503,1101,0,0,471,109,-2,2106,0,0,0,109,4,2101,0,-1,508,1207,-3,0,10,1006,10,526,21101,0,0,-3,21202,-3,1,1,21201,-2,0,2,21101,0,1,3,21101,0,545,0,1105,1,550,109,-4,2105,1,0,109,5,1207,-3,1,10,1006,10,573,2207,-4,-2,10,1006,10,573,21202,-4,1,-4,1106,0,641,21202,-4,1,1,21201,-3,-1,2,21202,-2,2,3,21101,0,592,0,1105,1,550,22101,0,1,-4,21101,1,0,-1,2207,-4,-2,10,1006,10,611,21102,1,0,-1,22202,-2,-1,-2,2107,0,-3,10,1006,10,633,22101,0,-1,1,21102,633,1,0,105,1,508,21202,-2,-1,-2,22201,-4,-2,-4,109,-5,2105,1,0`
// var input = `204,-7,99`
// var input = `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`
// var input = `109,19,204,-34,109,-34,104,1980,99`
// var input = `109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`
// var input = `1102,34915192,34915192,7,4,7,99,0`
// var puts int64 = 2   

var rb int64

func main() {
   // fmt.Println(34463338+34463338)
   var ip []string = strings.Split(input, ",")
   var o map[int64]int64 = make(map[int64]int64)
   for i, v := range ip {
     // var n int64
     n, _ := strconv.ParseInt(v, 10, 64)
     o[int64(i)] = n
     // o = append(o, n )
   }
   // o[43] = 90
   /*
      0 => up
      1 => right
      2 => down
      3 => left
   */
   cColor := ""
   x := 0
   y := 0

   minX := 0
   maxX := 0
   minY := 0
   maxY := 0

   coord := make(map[string]string)

   coord["0,0"] = "#"

   cDir := 0

   var inp int64 = 0
   input := 1

   var coms []int64

   for ;  ; {

      xx := strconv.Itoa(x)
      yy := strconv.Itoa(y)
      key := xx + "," + yy

      coms, inp = get(o, int64(input), inp)

      if o[inp] == 99 {
         break
      }

      if coms[0] == 1 {
         cColor = "#"
         input = 1
      } else {
         cColor = "."
         input = 0
      }

      if o[inp] == 99 {
         break
      }
      coord[key] = cColor

      if coms[1] == 1 {
         cDir++
      } else {
         cDir--
      }
      if cDir > 3 {
         cDir = 0
      }
      if cDir < 0 {
         cDir = 3
      }
      if cDir == 0 {
         y--
      }
      if cDir == 1 {
         x++
      }
      if cDir == 2 {
         y++
      }
      if cDir == 3 {
         x--
      }
      if x > maxX {
         maxX = x
      }
      if x < minX {
         minX = x
      }
      if y > maxY {
         maxY = y
      }
      if y < minY {
         minY = y
      }
      // in := strconv.Itoa(o[i])
   }
   for my := minY; my <= maxY; my++ {
      for mx := minX; mx <= maxX; mx++ {
         xx := strconv.Itoa(mx)
         yy := strconv.Itoa(my)
         key := xx + "," + yy
         // fmt.Print( yy + "," + xx, " ")
         c, exists := coord[key]
         if exists {
            fmt.Print(c)
         } else {
            fmt.Print(" ")
         }
      }
      fmt.Println()
   }
   // fmt.Println(minX,maxX,minY,maxY)
   fmt.Println(len(coord))
}

func get(o map[int64]int64, puts int64, start int64) ([]int64, int64) {

   // var rb int64 = 0
   r := make([]int64, 0)

   i := start
   for ; ; i++ {
      // rb = 0
      in := strconv.FormatInt(o[i], 10)
      l := len(in)

      opCode := -1
      mode1 := 0
      mode2 := 0
      mode3 := 0

      if l < 5 {
         in = strings.Repeat("0", 5 - l) + in
      }

      opCode, _ = strconv.Atoi(in[3:])
      mode1, _ = strconv.Atoi(string(in[2]))
      mode2, _ = strconv.Atoi(string(in[1]))
      mode3, _ = strconv.Atoi(string(in[0]))

      if opCode == 99 {
         break
      } else if opCode == 1 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         } 
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         if mode2 == 0 {
            p2 = o[p2]
         } 
         if mode2 == 2 {
            p2 = o[p2 + rb]
         }

         if mode3 == 2 {
            p3 = p3 + rb
         }

         o[p3] = p1 + p2

         i+=3
      } else if opCode == 2 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         }
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }
         if mode2 == 2 {
            p2 = o[p2 + rb]
         }

         if mode3 == 2 {
            p3 = p3 + rb
         }

         o[p3] = p1 * p2
         i+=3
      } else if opCode == 3 {
         p1 := o[i + 1]

         if mode1 == 2 {
            p1 = p1 + rb
         }

         if puts == -1 {
            return r, i
         }
         o[p1] = puts

         puts = -1

         i+=1
      } else if opCode == 4 {
         p1 := o[i + 1]

         if mode1 == 0 {
            p1 = o[p1]
         } 
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         r = append(r, p1)
         i++

         if len(r) > 1 {
            i++
            return r, i
         }
      } else if opCode == 5 {
         p1 := o[i + 1]
         p2 := o[i + 2]

         if mode1 == 0 {
            p1 = o[p1]
         }
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }
         if mode2 == 2 {
            p2 = o[p2 + rb]
         }

         if p1 != 0 {
            i = p2
            i--
         } else {
            i+=2
         }
      } else if opCode == 6 {
         p1 := o[i + 1]
         p2 := o[i + 2]

         if mode1 == 0 {
            p1 = o[p1]
         }
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }
         if mode2 == 2 {
            p2 = o[p2 + rb]
         }

         if p1 == 0 {
            i = p2
            i--
         } else {
            i+=2
         }
      } else if opCode == 7 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         }
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }
         if mode2 == 2 {
            p2 = o[p2 + rb]
         }

         if mode3 == 2 {
            p3 = p3 + rb
         }

         if p1 < p2 {
            o[p3] = 1
         } else {
            o[p3] = 0
         }
         i+=3
      } else if opCode == 8 {
         p1 := o[i + 1]
         p2 := o[i + 2]
         p3 := o[i + 3]

         if mode1 == 0 {
            p1 = o[p1]
         }
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         if mode2 == 0 {
            p2 = o[p2]
         }
         if mode2 == 2 {
            p2 = o[p2 + rb]
         }

         if mode3 == 2 {
            p3 = p3 + rb
         }
         // fmt.Println(p3)
         if p1 == p2 {
            o[p3] = 1
         } else {
            o[p3] = 0
         }
         i+=3
      } else if opCode == 9 {

         p1 := o[i + 1]

         if mode1 == 0 {
            p1 = o[p1]
         }
         if mode1 == 2 {
            p1 = o[p1 + rb]
         }

         rb += p1
         i+=1
      } else {
         break
      }

   }

   return r, i
}


