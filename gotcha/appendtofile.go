package gotcha

//	This code appends a line `of text to the file text.log. It creates the file if it doesn’t already exist.
//
//	f, err := os.OpenFile("text.log",
//	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
//	if err != nil {
//	log.Println(err)
//	}
//	defer f.Close()
//	if _, err := f.WriteString("text to append\n"); err != nil {
//	log.Println(err)
//	}
//
//	If you’re appending text to` a file for logging purposes, see Write log to file.
//
//https://yourbasic.org/golang/append-to-file/
