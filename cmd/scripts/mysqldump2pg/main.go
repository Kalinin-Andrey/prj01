package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

const (
	workDir	= "./deployment/data"
	fileIn	= "./db_source/auto_rus.sql"
	fileOut	= "./db/catalog_auto.sql"
	emptyLine = ""
)

type fileWalkFunc func(string) error
type strConvFunc func(string) (string, error)

var commentRegexp	= regexp.MustCompile("COMMENT '.+?'")
var intRegexp		= regexp.MustCompile(`int\(.+?\)`)
var tinyintRegexp		= regexp.MustCompile(`tinyint\(.+?\)`)

func main() {
	err := os.Chdir(workDir)
	if err != nil {
		log.Fatal(err)
	}

	in, err := os.OpenFile(fileIn, os.O_RDONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer in.Close()

	out, err := os.OpenFile(fileOut, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	isLineWithComma		:= false
	isPrevLineWithComma	:= false
	//_, err = fmt.Fprintln(w, "SET PGCLIENTENCODING=UTF8")

	err = fileWalkByLine(r, func(line string) error {
		resLine, err := strConv(line)
		if err != nil {
			return err
		}

		if resLine == "" {
			return nil
		}

		isLineWithComma = rmComma(resLine, &resLine)
		if isPrevLineWithComma && resLine != ");" {
			resLine = ", " + resLine
		}

		_, err = fmt.Fprintln(w, resLine)
		if err != nil {
			return err
		}
		isPrevLineWithComma = isLineWithComma
		return nil
	})


}

func fileWalkByLine(r *bufio.Reader, f fileWalkFunc) error {
	for {
		line, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				return err
			}
		}
		err = f(line)
	}
	return nil
}

func strConv(in string) (string, error) {
	in = strings.TrimSpace(in)
	in = rmbacktick(in)
	in = rmComment(in)

	switch {
	case rmSetSQLMode(in, &in):
	case rmSetTimeZone(in, &in):
	case rmTableSettings(in, &in):
	case rmAlter(in, &in):
	case rmAdd(in, &in):
	case rmModify(in, &in):
	case rmSQLComment(in, &in):
	case rmKeys(in, &in):
	case rmReplace(in, &in):
	case rmCharacterSet(in, &in):
	case rmTinyint(in, &in), rmInt(in, &in), hasInt(in):
		rmIntUnsigned(in, &in)
		rmIntAutoIncriment(in, &in)
	}
	return in, nil
}

func rmSetSQLMode(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, "SET SQL_MODE")
	if res {
		*out = emptyLine
	}

	return res
}

func rmSetTimeZone(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, "SET time_zone")
	if res {
		*out = emptyLine
	}

	return res
}

func rmAlter(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, "ALTER")
	if res {
		*out = emptyLine
	}

	return res
}

func rmAdd(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, "ADD")
	if res {
		*out = emptyLine
	}

	return res
}

func rmModify(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, "MODIFY")
	if res {
		*out = emptyLine
	}

	return res
}

func rmSQLComment(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, "/*!")
	if res {
		*out = emptyLine
	}

	return res
}

func rmbacktick(l string) string {
	l = strings.ReplaceAll(l, `\'`, `\"`)
	return strings.ReplaceAll(l, "`", "")
}

func rmTableSettings(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, ") ENGINE")
	if res {
		*out = ");"
	}

	return res
}

func rmComment(l string) string {
	return commentRegexp.ReplaceAllString(l, "")
}

func rmTinyint(l string, out *string) bool {
	*out = tinyintRegexp.ReplaceAllString(l, "smallint")
	return l != *out
}

func rmInt(l string, out *string) bool {
	*out = intRegexp.ReplaceAllString(l, "int")
	return l != *out
}

func rmIntUnsigned(l string, out *string) bool {
	*out = strings.ReplaceAll(l, "int UNSIGNED", "int")
	*out = strings.ReplaceAll(*out, "int unsigned", "int")
	return l != *out
}

func rmIntAutoIncriment(l string, out *string) bool {
	*out = strings.ReplaceAll(l, "NOT NULL AUTO_INCREMENT", "PRIMARY KEY")
	return l != *out
}

func rmReplace(l string, out *string) bool {
	*out = strings.ReplaceAll(l, "REPLACE", "INSERT")
	return l != *out
}

func rmCharacterSet(l string, out *string) bool {
	*out = strings.ReplaceAll(l, "CHARACTER SET utf8", "")
	*out = strings.ReplaceAll(*out, "COLLATE utf8_general_ci", "")
	return l != *out
}

func hasInt(s string) bool {
	return strings.Index(s, "int") > -1
}

func rmKeys(l string, out *string) (res bool) {

	res = strings.HasPrefix(l, "UNIQUE KEY") || strings.HasPrefix(l, "KEY") || strings.HasPrefix(l, "PRIMARY KEY") || strings.HasPrefix(l, "CONSTRAINT") || strings.HasPrefix(l, "LOCK TABLES") || strings.HasPrefix(l, "UNLOCK TABLES")
	if res {
		*out = emptyLine
	}

	return res
}

func rmComma(l string, out *string) (res bool) {

	s := strings.TrimRight(l, ",")
	res = l != s
	if res {
		*out = s
	}

	return res
}
