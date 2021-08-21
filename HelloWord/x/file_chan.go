package main

//var numChan =make(chan int)
//
//func readFile(infile string){
//	fin,err :=os.Open(infile)
//	if err !=nil{
//		fmt.Println(err)
//		return
//	}
//	defer fin.Close()
//	reader :=bufio.NewReader(fin)
//	for{
//		line,_,err :=reader.ReadLine()
//		if err !=nil{
//			if err == io.EOF{
//				break
//			}else{
//				fmt.Println(err)
//			}
//		}
//	}
//
//}

func calTotle(line string) int{
	sum :=0
	for _,c:=range line{
		sum +=int(c)
	}
	return sum
}

//func writeFile(outfile string){
//	fout,err:=os.OpenFile(outfile,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,)
//	if err !=nil{
//		fmt.Println(err)
//		return
//	}
//	defer fout.Close()
//	writer :=bufio.NewWriter(fout)
//	for{
//		n,ok :=<-numChan
//		if !ok {
//			break
//		}else{
//			writer.WriteString(strconv.Itoa(n))
//			writer.WriteString("\n")
//		}
//	}
//}
