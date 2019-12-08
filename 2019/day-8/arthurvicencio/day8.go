package main

import (
   "fmt"
   "strconv"
)

var input string = `122222222022222222222222222202222222200122212222222222222222222222122222220220222022222202222222222022222120122022222022220222222221222202002202222122122222222122222222221222222202222222201222222222222222222222222222222222222222222012222202222222222122222022122022222022221222222221222212002202222022222222222122222222220222222212222222220122212222222222222222222202022222222222222102222222222222222122222020222022222122220222222222222212102202222222222222222022222222220222222212222222222122202222222222222222222212022222220220222112222222222222222022222120222222222022220222222221222222012212222122122222222222222222221222222212222222221222222222222222222222222202222222221221222012222202222222222222222021022022222022220222222221222212022212222022222222222022222222222222222212222222202122222222222222222222222222122222221221222102222222222222222222222221222222222222222222222220222202222212222022122222222122222222222222222212222222202222202220222220222222222212112222222221222212222222222222222122222222022022222222221222222221222212012202222022122222222022222222221222222222222222200022222220222221222222222202022222220220222222202222222222222222222122022022222222222222222220222212122222022122022222222122222222221222222212222222221222212220222222222222222212201222221220222122212222222222222122222021122222222222220222222222222212112202022022122222222122222222221222222212222222201022212222222222222222222202111222222222222202212212220222222122222020122222222122221222222222222212202212122122222222222222222222221222222212122222221122202222222222222222222222020222222222222102222202220222222222222122222022222022220222222221222212102212122122022222222222222222221222222202222222202122202200222220222222222221222222220222222102202222222222222022222122222122222022221222222220222202102222022022222222222022222222220222222212222222212222202211222220222222222201211222220220221122222212220222222020222020122222222022222222222220222212012222222222222222222222222222221222222202122222210022222221222222222222222221000222222221220222222212221222222221222122222022222022222222222220222202012212122222222222222122222222222222222212122222200222222211222220222222222210101222222222221202222212222222222022222120022122222022221222222221222222202222122122122222222122222122220222222202222222201022202220222221222222222212201222221221222122222202222222222222222221222022222112222222222221222202202212122122122222222022222222220222222202022222221022212220222221222222222221020222220220220002002212220222222122222221122022222012221202222221222212012212222122222222222122222022220222222202122222222222222210222221222222222211201222221221220112222202221222222222222020022012222222221222222222222222212222222022222222222122222222202222222202222222211222222200222220222222222201000222220221222002212222222222222221222220222022222112222202222222222222112202022022122222222022222020202222222222122222210122222200222220222222222221221222220222222212212212221222222021222220122002222212220202222221222222102222222222120222222122222022212222222212122222201122212210222222222222222220210222222220220202122222222222222121222221222022222212221212222220222222012222222022120222222022222020222222222202222222211122202200222222222220222202202222221221222012012212221222222222222122022212222102221202222222222222212212222222120222222022222220221222222222022222202222212222222221222220222222222222221222221122022222221122222020222020222212222002221222222220222212102222122122020202222122222020211122222222222222200122202210222221222221222200020222220220222102222202221022222022222122222122222222221222222222222212212202022022120222222222222221111222222212222222202022222202222222222221222212001222220220221022022222221022222122222020022002222222222202222221222202102222222222022202222122222022000222222222122222222222212211222221222222022211020222220220220212022202222122222022222020020122222112220222222221222212222202022222220202222022222122112122222212122222222222222212222220222221022201000222220220220122112222222122222220222122222102222102221222222222222212022222222022120212222122222222000022222212022222220222222222222221222220022222222222221222221012122222221122222022222022121102222202222202222222222222212202222122221202222222222102101222222202022222222222202220222221222221122212122222222222220212020212220022222021222120022002222022221212222221222222212212122222121222222222222221221122222202022222220222202212222221222221222202001222201220220102201222222222222220222221221012222102220202222222222222202222022022022222222022222100112222222212120222221122222220222222222222222200111222221222222002110202222222222121222120122212222202222212122221222222122202222222220212222022222022222122222202020222221022222212222220222220222202112222212220220202002202220022222120222021120012222102220202122221222202112212222222121202222122222210100122222212220122201122222200222221222220120221100222212220222212202202220022222221222220022102222122222202222222222212202002022122022222222122222112010122222212222221200222212220222220222220022211200222200222220022112212222222222221222020022012222122222212122220222222112002222122221202222122222100200022222212122122201222222200222220222222122211022222222220222212000202220022222122222221222212222202220212122222222222122102122222022222222122222202002022222222222222221122202210222220222222121211202222222220220122200222221122222221222022220222222112220202222222222002012012122122121212222222222002111020222212220022202122212210222221222220222201020222222222220212111222222222222222222022221212222202220202222220222122112102022222020202222122221122101220222212120120222222202202222222222222121201101222201222221122200222221122222122222121122022222112221102122220222112222202022222021202222222220222112120222202120021201022212211222220222222120201020222221220221002022222222122222020222221220012222112221202022222222002002202222022121222222222220100122022222222222222221222222221222220222222022212101222202220222212210212221022222022222220120002222112222212222221222102102022022222121222222012222102001121222212121020201222212220222222222221022211110222201220221102020202222222222221222021120202222002220122122221222022202212222022020202222202220112111022222202222121210022212222222221222221021220121222210221222002201202222222222120222122121002222112222212122222222222212122122022220222222012222112210020222202120220220022202211202221222222022211100222220220222212001202222222222222222021221112222002220111022220222002102002122222122222222202221002022221222202020222200122222212212221222220122212011222222222222122121202221122222120222121020012222122222000122220222202022022022222021202222202221211121220222202222222211222212211222221222222020211120222211221222002022222221122222020222022022022222222220202022221222002222222122122220212222020221111201121222222220222202022202200202222222220121201210222202222221002122202221222222021222221221022212102220211022222222102022002022222221202222222220002020222222212120020210122202211212221222221220222202222201221220012112212221022222120222122221222212012221020022221222202002012022222021212222102222001000221222222120121201122222210202221222221120211000222211222221002001222221122222120222220221212202022221221022221222202022212222122021202222010222012102120222212021220202022222222222221212222222202021222212221222022022222220222222222222221120212202222221012222221222202102212212222221222222212220120200122222222221222210222212211222221202221120220111222200222221022012212221122222020222102021222202022222111022220222202222002102022021212222102221210120120220202220022222122202221212220212222221220101222212222201022011222221222222120222001020112222122222212022222222122212002202022220212222112221000022121221222221120220222212210222222202221022221010222210220210022021222220222222120222010122002222112220212222222222222102222022022122202222000221212020220222222121021200222202201202222202221221222101222210222220222101212222222222222222022121222222012220012222221222202012202122122022202222020222011102220220212222222212022222221212222002222011222020222212220221002120202222222222022222022220022212012222001022222222202112102122022220222222000222010122021222212221221220022212220202222020222211210021222202222202102121222220022222122222110222122222022220202022221222212022112122122020202222122220102010022022202222221202222222201102220121221210211001222210222210112221222221122222122222011121212222202201201022222222112012112112222021212222120222022020122221212020222222222222210102020202221221222210222210221212102221202221022202221222012120002202112220102022221222212212102022022020212222202220121102222120202221022201122222221122020001221221202201222201220202112000212222122212221222002121022212102211101022222222112012222012122121202222012220011011220022202121222220022202201112220221221010201020222211220220212020222220122202222222102022012202222212212222222222112012022212222021202222210221100212020022222220120212022212202022120001222201200000222211220210112002222220122202122222211121222202112202002122220222222112022222202222222222001222212022022022212120021221122202210102222222220001222111222211222202222102202222222222020222111121212202212211112122221222112012102112022021212222022221001202021222212122220201122202211122222200222211202202222211221211112000212221022222222202110220222222222221201022220222002212212202012021222222202221110220022220202021020221022222212112021201220000200010222222222221222111212222022202020202112222222222012212122222222222012002002002102220212222021221120222021020222120121221022202220022222011221201220100222221220212222212222220222212022222210122112222102210200222220222112002012222012221202222201222112222120021222022021222222202201112022012222010210002222220221220022021202221022212221212220121012202202201221022222222012212202012022122212222022222022120220122212122021222122212211112220100222212210121222200222221122001212220122222112222202220212212012210022122220222112102012022102222222222220221111022022122212222221211022212220002021012220120221100222211222212122221212220022202220222212220202202212212020222220222022002122022102122212222020202220122020221222122121221222212201212020012222110211122222201221221022220222200122212101202102022212202112212202122222222022222022012102222212222021220000102121121202120221200122222201201120200220001212122222211220221122010212200022222002222002120122212022212121222222222202222002222012120212222022212111102122021222020120220022202222011220000220001221212222210220201002111222221122212101222010220212212122220112022221222202022112002112120222222211212011000122222222120021211022212200000222202222002220201222200222202122212212222122212000212000120122222102200212122220222112222212202122122212222201221201222021021212121122220022202222120120220220100220021222220222202222021112201222212022212020122212222022222102222222222122222102022202220212222222202222222221022212201221200222202210222122020220211220220222202222212212222122220222212012202000121112220212202112222222222212202022122212221202222220201200122020220202202222221222202221022120121222201212110222200222211012200222210222202211212200122222200102200200022222222112112212012022221222222000201011012122120222210022222022212221011122000221222202201222201221202122100212200222222212202111022202200002210010222220222202222222202202120202222120222010012222121212111221221112212211210220002220111212210222201220210012000022202022202200202102022212000122220101022222222112222002122002220222222100212020201120020222020221220222212211101221222221021211110222200221202112122202220022222102202211022222210112202102122222222221122122212222122202222221221211121022020202211122211002212222100021012220212220222222200221220212020110220022222212202200121212200112202021022220222210122212122112122002222002221002201121021202220220200002202210100222112221201212200222221222201002201221220022202202202222221022212202211112022221222122212112222222020012222221201222011120020202201021201002212220021020200222022211101222201221222112011011200022202111202120020222110012211002222221222212222202222102120212222211200002211120120212022020220222202212122122121222120202221222212220211202122101212122212200202001222122010202210012222220222000112212112122122212222012220000002220221202011120201102212200210120111220222222012222210221211222211210210122212200202220121212100102212122022220222012012012012212122112222022221221020020021202101220212122202200211021102221200111002222202222212212111002202222222101222121020012121202210210222220222220022112222222120212222020222210100221220222101022202112222211021221010222201211011222112220211022211221211222202110222110121022010122121201122222222201212012222012122202222121200200112020222212222022202202202220022021110220001002001022221222200122121011222222202121212220120022001202220222122221222121202012112002221202222000221120012222220202102020201122222212221022111222001102211222200221221122000000211122222100202222021112120002122102222220222212222202212222221122222022220210211021121202110221210000012201202020111221110210121122011222200012212100220222222121210110221202010012022000122221222220222212012002022002222011200021211121021212020122220021002222101020202220201100112022221121221102000100201122222022221220121122111122111000222220212012002012222112021022222202211112212122222202022020202000112220120120101221022212022122110021221102000121220222202001210011022102001202202111122220222002002022202112022022122211011111012221222222012021222200212201021022110220211012102222011120210222012100202022222012210000220212121002020121222222202010122122222112022102022022002022201222121212200022201200102200000220201220022100022122212222222122022020210122202202212122121022010202011000222222222120202212002222021122222102220212002020021202102020201120022212011121101220111220201222210121202212001121102122222010221100120122201222220210122222222120022122102122121222122021222021102022122222202121200022102202022220022222101200020221122220212202100102120222222220201101022012001102120222022220202021111212021212221212022220021001202020122202122222220012202210212020000221022212100222222020211112210200020222202021220122222102210022122020222222222200221012210222120200022010000200010021022202122021220110002220002221002222121112002022212222220202121212100222212100221120120002012112212020022220222001220022222122221202222210121202002120221202011121222100012222201220100220020202211020120022211122221202221222222020200121020012111212020111212221202202201112102222221001022110220001210221022202022021200210212201021122110220001201012121121120211212002202112122212002200111020212112122121212202020212022222122110002021010022202021100211120120222012122212110222220200222212220020011222120121222220212220012012122222210202112222122020022021222102022202121102002202110121210222212111221222122020222020022212001122200221222110220220012201022120122221011000101210222202022202021222002211202221022022122202220011011022222001200110102112020010210010020222101012110100111220001211000101000001212001000012012022010000101001121001020201100121010102111211001110011121101110211`
var width int = 25
var height int = 6

func main() {
   var tl int = len(input) / (width * height)
   var layers [][][]int = make([][][]int, tl)
   // var layerLen int = width * height
   
   // minZero := 0
   // sLayer := -1

   inputP := 0

   for layer := 0; layer < tl; layer++ {
      // zeroCount := 0

      for h := 0; h < height; h++ {
         for w := 0; w < width; w++ {
            d, _ := strconv.Atoi(string(input[inputP]))
            // if d == 0 {
               // zeroCount++
            // }
            if layers[layer] == nil {
               layers[layer] = make([][]int, height)   
            }
            if layers[layer][h] == nil {
               layers[layer][h] = make([]int, width)   
            } 
            layers[layer][h][w] = d   
            inputP++
         }
      }

      // for i := 0; i < layerLen; i++ {
      //    d, _ := strconv.Atoi(string(input[i+(layer * layerLen)]))
      //    if d == 0 {
      //       zeroCount++
      //    }
      //    layers[layer] = append(layers[layer], d)
      // }
      // if zeroCount < minZero || minZero < 1 {
         // sLayer = layer
         // minZero = zeroCount
      // }
   }

   image := make([][]string, height)
   for l := len(layers) - 1; l >= 0; l-- {
      for h := 0; h < height; h++ {
         if image[h] == nil {
            image[h] = make([]string, width)   
         }
         for w := 0; w < width; w++ {
            val := ""
            if layers[l][h][w] == 1 {
               val = "1"
            }
            if layers[l][h][w] == 0 {
               val = "0"
            }
            if val == "" {
               continue
            }
            image[h][w] = val
         }
      }
   }

   for h := 0; h < height; h++ {
      for w := 0; w < width; w++ {
         fmt.Print(image[h][w])
      }
      fmt.Println()
   }
   // one := 0
   // two := 0
   // for _, v := range layers[sLayer] {
      // if v == 1 {
      //    one++
      // }
      // if v == 2 {
      //    two++
      // }
      // for h := 0; h < height; h++ {
      //    for w := 0; w < width; w++ {
      //       if layers[sLayer][h][w] == 1 {
      //          one++
      //       }
      //       if layers[sLayer][h][w] == 2 {
      //          two++
      //       }
      //    }
      // }
   // }
   
   // fmt.Println(one * two)
}
