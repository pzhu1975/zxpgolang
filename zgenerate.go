// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package parser implements a parser for Go source files. Input may be
// provided in a variety of forms (see the various Parse* functions); the
// output is an abstract syntax tree (AST) representing the Go source. The
// parser is invoked through one of the Parse* functions.
//
// The parser accepts a larger language than is syntactically permitted by
// the Go spec, for simplicity, and for improved robustness in the presence
// of syntax errors. For instance, in method declarations, the receiver is
// treated like an ordinary parameter list and thus may contain multiple
// entries where the spec permits exactly one. Consequently, the corresponding
// field in the AST (ast.FuncDecl.Recv) field is not restricted to one entry.
//
package zgenerate

import (
	"fmt"
	"bytes"
	"io"
	"os"
	"io/ioutil"
	"errors"
	"github.com/wordanalysis/zscanner"
	"github.com/wordanalysis/ztoken"
)

// The parser structure holds the parser's internal state.
type Zgerneratego struct {
	inputfilename    string  //file operate parameter
	inputdata		[]byte
	filestatus		int
	linedata		[]byte
	outputfilename	string  //file comment
	filset  *ztoken.FileSet
	file    *ztoken.File
	zscanner zscanner.Scanner

	comments    	string
	// Next token
	pos ztoken.Pos   // token position
	tok ztoken.Token // one token look-ahead
	lit string      // token literal	
	// Tracing/debugging
	mode   			uint // parsing mode
	trace  			bool // == (mode & Trace != 0)	
	indent			int  //scanner	
}

const  defaultoutfilename string= "xptest.go"
const (
	INITAL int = iota
	OPENED
	CLOSED
	FINISHED
)

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func Readtext(filename string, src interface{}) ([]byte, error) {
	if src != nil {
		switch s := src.(type) {
		case string:
			return []byte(s), nil
		case []byte:
			return s, nil
		case *bytes.Buffer:
			// is io.Reader, but src is already available in []byte form
			if s != nil {
				return s.Bytes(), nil
			}
		case io.Reader:
			return ioutil.ReadAll(s)
		}
		return nil, errors.New("invalid source")
	}
	return ioutil.ReadFile(filename)
}

func FilefromRead(filename string) ([]byte,error) {
    fileObj,err := os.Open(filename)  
    if err != nil{  
        panic(err)  
    }  
	buftemp := make([]byte,512)
    buf := make([]byte,1024)  
    for{  
        n,err := fileObj.Read(buf)  
        if err != nil && err != io.EOF{panic(err)}  
        if 0 ==n {break}  
        buf=append(buf,buftemp...)
        // fmt.Println(string(buf[:n]))  
	}  
	defer fileObj.Close() 	
    return buf,err  	
    // if fileObj,err := os.Open(filename);err == nil {
    //     defer fileObj.Close()
    //     //在定义空的byte列表时尽量大一些，否则这种方式读取内容可能造成文件读取不完整
    //     // if _,err = fileObj.Read(databuf);err == nil {
    //     //        fmt.Println("Use os.Open and File's Read method to read a file:",string(databuf))
	// 	//     }
	// 	return ioutil.ReadAll(fileObj) 
	// }
	// return databuf,err
}

//使用io.WriteString()函数进行数据的写入
func Writetofile(filename,srccontent string) (n int,err error) {
    fileObj,err := os.OpenFile(filename,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
    if err != nil {
        fmt.Println("Failed to open the file",err.Error())
        os.Exit(2)
	}
    if  n,err = io.WriteString(fileObj,srccontent);err == nil {
        fmt.Println("Successful os.OpenFile and io.WriteString.",srccontent)
	}
	defer fileObj.Close()	
	return n,err
}

func DelFile(filename string) error{
	return os.Remove(filename)
}

func AppendToFile(fileName string, content string) error {
	// 以只写的模式，打开文件
	f, err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	if err != nil {
	   fmt.Println("cacheFileList.yml file create failed. err: " + err.Error())
	} else {
	   // 查找文件末尾的偏移量
	   n, _ := f.Seek(0, os.SEEK_END)
	   // 从末尾的偏移量开始写入内容
	   _, err = f.WriteAt([]byte(content), n)
	}
defer f.Close()   
return err
}

func (p *Zgerneratego) init(fromname string,toname string, src []byte) {
	p.inputfilename=fromname
	p.inputdata=src
	p.filset = ztoken.NewFileSet()
	if(toname==nil)
		p.outputfilename=string(defaultoutfilename)
	else
		p.outputfilename=toname
	p.filestatus=INITAL
	file := fset.AddFile("", fset.Base(), len(text)) // register input "file"
}

func (p *Zgerneratego) trace() {
	p.indent++  // "(" counter
	return p
}

// Usage pattern: defer un(trace(p, "..."))
func (p *Zgerneratego) un() {
	p.indent--		// ")" counter
}

func (p *Zgerneratego) Wrtietofile(databuf string) (error){
	var err error
	var fexist bool
	if(databuf==nil||databuf==""){
		return errors.New("data is null")
		}
	if(p.outputfilename==nil||p.outputfilename==""){
		return errors.New("file is null")
		}
	if(p.filestatus==INITAL){
		fexist,err=FileExists(p.outputfilename)
		if(fexist){
			err=DelFile(p.outputfilename)
		}
		_,err=Writetofile(p.outputfilename,databuf)
		p.filestatus=CLOSED
		return err
		} else {
		err = AppendToFile(p.outputfilename,databuf)
		}
	return err
}

// Advance to the next token.
func (p *Zgerneratego) next0() {
	// Because of one-token look-ahead, print the previous token
	// when tracing as it provides a more readable output. The
	// very first token (!p.pos.IsValid()) is not initialized
	// (it is token.ILLEGAL), so don't print it .
	if p.trace && p.pos.IsValid() {
		s := p.tok.String()
		switch {
		case p.tok.IsLiteral():
			p.printTrace(s, p.lit)
		case p.tok.IsOperator(), p.tok.IsKeyword():
			p.printTrace("\"" + s + "\"")
		default:
			p.printTrace(s)
		}
	}

	p.pos, p.tok, p.lit = p.scanner.Scan()
}

func (p *Zgerneratego) Parse() ([]byte,error){

}
