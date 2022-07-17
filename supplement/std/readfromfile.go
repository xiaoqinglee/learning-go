package std

//	Read a file (stdin) line by line
//
//
//	Read from file
//	Read from stdin
//	Read from any stream
//
//	Read from file
//
//	Use a bufio.Scanner to read a file line by line.
//
//	file, err := os.Open("file.txt")
//	if err != nil {
//	log.Fatal(err)
//	}
//	defer file.Close()
//
//	scanner := bufio.NewScanner(file)
//	for scanner.Scan() {
//	fmt.Println(scanner.Text())
//	}
//
//	if err := scanner.Err(); err != nil {
//	log.Fatal(err)
//	}
//
//	Read from stdin
//
//	Use os.Stdin to read from the standard input stream.
//
//	scanner := bufio.NewScanner(os.Stdin)
//	for scanner.Scan() {
//	fmt.Println(scanner.Text())
//	}
//
//	if err := scanner.Err(); err != nil {
//	log.Println(err)
//	}
//
//	Read from any stream
//
//	A bufio.Scanner can read from any stream of bytes, as long as it implements the io.Reader interface. See How to use the io.Reader interface.
//
//https://yourbasic.org/golang/read-file-line-by-line/

//	input := bufio.NewScanner(os.Stdin)
//	//默认的分割函数是bufio.ScanLines 分割行
//	input.Split(bufio.ScanWords) //分割单词
