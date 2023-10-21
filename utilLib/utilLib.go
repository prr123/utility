package utilLib
//
//
// library that parses text for chars
//
// Created: 12/4/2022
// Author: prr, azul software
// Modified: 21/10/2023
// copyright (c) 2022, 2023 prr, azul software
//
// mod 21/10/2023
// added string to byte conversion with no memory allocation
//

import (
	"os"
	"fmt"
	"strings"
	"unsafe"
	"reflect"
)

func CheckFilnam(filnam, ext string)(res bool) {
// check file name extension
	res = true
    idx := strings.Index(filnam, ext)
    if idx < 0 { res=false}
    return res
}

// routine that reads cli and returns a map of key and values
func ParseFlagsStart(args []string, flags []string, flagStart int) (argmap map[string]interface{}, err error){

	argmap = make(map[string]interface{})
    numArg:= len(args)
    if numArg < 2 { return argmap, nil}
	if (len(flags) == 0 ) { return argmap, fmt.Errorf("insufficient flags!")}
	if flagStart == 0 {flagStart = 1}
	pos:= 0

	for j:=flagStart; j< numArg; j++ {
		str := args[j]
		if len(str) < 2 { return argmap, fmt.Errorf("invalid flag too short: %d %s",j, str)}
		if str[0] != '/' {return argmap, fmt.Errorf("invalid flag no slash: %d %s",j, str)}

		pos = 0
		for j:=1; j<len(str); j++ {
			if str[j] == '=' {
				pos = j
				break
			}
		}
		flagKey := ""
		flagVal := ""
		if pos==0 {
			flagKey = str[1:]
			flagVal = "none"
		} else {
			flagKey = str[1:pos]
			flagVal = str[pos+1:]
		}


		found := false

		for i:=0; i<len(flags); i++ {

			if flagKey == flags[i] {
				argmap[flagKey] = flagVal
				found = true
			}
			if found { break}
		}
		if !found { return argmap, fmt.Errorf("arg %d: %s is not a viable flag!", j, str) }
	}

	return argmap, nil
}

// routine that reads cli and returns a map of key and values
func ParseFlags(args []string, flags []string) (argmap map[string]interface{}, err error){

	argmap = make(map[string]interface{})
	numArg:= len(args)
	if numArg < 2 {
		return argmap, nil
	}

	pos:= 0

	for j:=1; j< numArg; j++ {
    	str := args[j]
		if len(str) < 2 { return argmap, fmt.Errorf("invalid flag too short: %d %s",j, str)}
		if str[0] != '/' {return argmap, fmt.Errorf("invalid flag no slash: %d %s",j, str)}

		pos = 0
		for j:=1; j<len(str); j++ {
			if str[j] == '=' {
				pos = j
				break
			}
		}
		flagKey := ""
		flagVal := ""
		if pos==0 {
			flagKey = str[1:]
			flagVal = "none"
		} else {
			flagKey = str[1:pos]
			flagVal = str[pos+1:]
		}

		found := false

		for i:=0; i<len(flags); i++ {
			if flagKey == flags[i] {
				argmap[flagKey] = flagVal
				found = true
			}
			if found { break}
		}
		if !found { return argmap, fmt.Errorf("arg %d: %s is not a viable flag!", j, str) }
	}
	return argmap, nil
}

// routine that reads cli and returns a map of key and values
func GetFlags(args []string) (argmap map[string]interface{},  err error){

	numArg:= len(args)
	if numArg < 2 {
		return argmap, fmt.Errorf("insufficient args!")
	}

	pos:= 0
	argmap = make(map[string]interface{})
	for i:=2; i< numArg; i++ {
		str := args[i]
		if str[0] != '/' {
			return argmap, fmt.Errorf("invalid flag no/: %d %s",i, str)
		}
		for j:=1; j<len(str); j++ {
			if str[j] == '=' {
				pos = j
				break
			}
		}

		if pos==0 {
			return argmap, fmt.Errorf("invalid flag no=: %d %s",i, str)
		}
		flag := str[1:pos]
		flagVal := str[pos+1:]
		argmap[flag]=flagVal
	}
	return argmap, nil
}

// function that parse a string and returns a slice of strings for each word ending with a comma
func ParseList(src string)(dest *[]string, err error) {

    var list [10]string
    count:= 0
    stPos:=0
    for i:=0; i< len(src); i++ {
        if src[i] == ',' {
            list[count] = string(src[stPos:i])
            stPos = i+1
            count++
			if count == 11 {
				return nil, fmt.Errorf("maximum number of items 10  exceeded!")
			}
        }
    }

    list[count] = string(src[stPos:])
    count++
    lp := list[:count]

    return &lp, nil
}

func PrintList (list *[]string) {

    fmt.Printf("items: %d\n", len(*list))
    for i:=0; i< len(*list); i++ {
        fmt.Printf("%d: %s\n", i+1, (*list)[i])
    }

}


// function that tests whether byte is alpha
func IsAlpha(let byte)(res bool) {
	res = false
	if (let >= 'a' && let <= 'z') || (let >= 'A' && let <= 'Z') { res = true}
	return res
}

// function that tests whether byte is not alpha
func NotAlpha(let byte)(res bool) {
	res = false
	if (let < 'a' || let > 'z') && (let < 'A' || let > 'Z') {res = true}
	return res
}

// function that tests whether byte is aphanumeric
func IsAlphaNumeric(let byte)(res bool) {
	res = false
	tbool := (let >= 'a' && let <= 'z') || (let >= 'A' && let <= 'Z')
	if tbool || (let >= '0' && let <= '9') { res = true }
    return res
}

// function that tests whether byte is numeric
func IsNumeric(let byte)(res bool) {
	res = false
	if (let >= '0') && (let <= '9') { res = true }
	return res
}

// function that tests whether byte is whitespace
func IsWsp(let byte)(res bool) {
	res = false
	if let ==' ' { res = true}
	return res
}

// https://github.com/valyala/fasthttp
// no memory allocation conversions
func Byt2Str(b []byte) string {
    return *(*string)(unsafe.Pointer(&b))
}

func Str2Byt(s string) (b []byte) {
    bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
    sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
    bh.Data = sh.Data
    bh.Cap = sh.Len
    bh.Len = sh.Len
    return b
}


// function that displays a console error message and exits program
func FatErr(fs string, msg string, err error) {
	if err != nil {
		fmt.Printf("error %s:: %s!%v\n", fs, msg, err)
	} else {
		fmt.Printf("error %s:: %s!\n", fs, msg)
	}
	os.Exit(-1)
}

// creates file in file path and returns file handle
func CreateOutFil(folderPath, filNam, filExt string) (outfil *os.File, err error) {
	var fullFilNam, filpath string

	if len(filNam) == 0 {
		return nil, fmt.Errorf("file name is empty!")
	}
	// create full file name: filnam + ext
	// check file extension
	if len(filExt) == 0 {
		ext := false
		for i:=len(filNam) -1; i>=0; i-- {
			if filNam[i] == '.' {
				ext = true
				break
			}
		}
		if !ext {return nil, fmt.Errorf("no file extension provided!")}
		fullFilNam = filNam
	} else {
		// check extension
		if filExt[0] == '.' {
			fullFilNam = filNam + filExt
		} else {
			fullFilNam = filNam + "." + filExt
		}
	}

	lenFP := len(folderPath)
	if lenFP > 0 {
		filinfo, err := os.Stat(folderPath)
		if err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf("folderPath %s is not valid!", folderPath)
			}
		}
		if !filinfo.IsDir() {
			return nil, fmt.Errorf("folderPath %s is not a folder!", folderPath)
		}
		if folderPath[lenFP-1] == '/' {
			filpath = folderPath + fullFilNam
		} else {
			filpath = folderPath + "/" + fullFilNam
		}
	} else {
		filpath = fullFilNam
	}

	// check whether file exists
	_, err = os.Stat(filpath)
	if !os.IsNotExist(err) {
		err1:= os.Remove(filpath)
		if err1 != nil {
			return nil, fmt.Errorf("os.Remove: cannot remove existing file: %s! error: %v", filpath, err1)
		}
	}

	outfil, err = os.Create(filpath)
	if err != nil {
		return nil, fmt.Errorf("os.Create: cannot create file: %s! %v", filpath, err)
	}
	return outfil, nil
}

// method that creates a file folder, returns the full path, true if folder exists
func CreateFileFolder(path, foldnam string)(fullPath string, existDir bool, err error) {

    // check if foldenam is valid -> no whitespaces
	fnamValid := true
	for i:=0; i< len(foldnam); i++ {
		if foldnam[i] == ' ' {
			fnamValid = false
			break
		}
	}

	if !fnamValid {
		return "", false, fmt.Errorf("error -- not a valid folder name %s!", foldnam)
	}

    // check whether foldnam folder exists
	fullPath =""
	switch {
	case len(path) == 0:
		fullPath = foldnam

	case path[0] == '/':
		return "", false, fmt.Errorf("error -- absolute path!")

	case path[len(path)  -1] == '/':
		fullPath = path + foldnam

	default:
		fullPath = path + "/" + foldnam
	}

	// check path with folder name
    // add trimming wsp to left
	if _, err1 := os.Stat(fullPath); !os.IsNotExist(err1) {
		return fullPath, true, nil
	}

    // path does not exist, we need to create path
	ist:=0
	for i:=0; i<len(fullPath); i++ {
		if fullPath[i] == '/' {
			parPath := string(fullPath[ist:i])
			if _, err1 := os.Stat(parPath); os.IsNotExist(err1) {
				err2 := os.Mkdir(parPath, os.ModePerm)
				if err2 != nil {
					return "", false, fmt.Errorf("os.Mkdir: lev %d %v", err2, i)
				}
			}
		}
    }
	err = os.Mkdir(fullPath, os.ModePerm)
	if err != nil {
		return "", false, fmt.Errorf("full Path os.Mkdir: %v", err)
	}

	return fullPath, false, nil
}

