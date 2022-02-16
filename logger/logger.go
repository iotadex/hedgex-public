package logger

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// LOGLEVEL log level
type LOGLEVEL int

const (
	// INFO information log
	INFO LOGLEVEL = iota
	// WARN warnning log
	WARN
	// ERROR error log
	ERROR
	// DEBUG debug log
	DEBUG
)

var leveName = [4]string{
	INFO:  "[INFO]",
	WARN:  "[WARNING]",
	ERROR: "[ERROR]",
	DEBUG: "[DEBUG]",
}

// Logger information
type Logger struct {
	filename  string      // logger file's name : test_2019-06-26_22-53.log. it's filename is 'test.log'
	file      *os.File    // logger's output file
	rotate    int         // 1 for date rotate; 2 fro size rotate; 0 for not rotate
	rDateTime time.Time   // current day
	rHour     int         // current hour
	rMinute   int         // current minute
	maxSize   int64       // file's max size
	nRotate   int         // file's count
	cRotate   int         // current rotate file
	fnRotate  []string    // rotate file name
	out       io.Writer   // destination for output
	bufChan   chan []byte // log data buf chan
	bPrefix   bool        // auto add time prefix to log
	closed    bool        // close the out or not
}

// GetPathFileName return the filename's fullpath,filename and the suffix
func GetPathFileName(fn string) (string, string, string) {
	var path, file, suffix string
	if len(fn) > 0 {
		indexFile := strings.LastIndex(fn, "/")
		indexLastDot := strings.LastIndex(fn, ".")

		if indexLastDot > 0 && indexFile < indexLastDot {
			file = fn[indexFile+1 : indexLastDot]
			if indexLastDot < (len(fn) - 1) {
				suffix = fn[indexLastDot:]
			}
		} else if indexLastDot == -1 {
			file = fn[indexFile+1:]
		}
		if len(suffix) == 0 {
			suffix = ".log"
		}
		path = fn[0:(indexFile + 1)]
		var dir string
		if (len(path) > 0 && path[0] != '/') || (len(path) == 0) {
			dir, _ = filepath.Abs(filepath.Dir(os.Args[0]))
			dir += "/"
		}
		path = dir + path
	}
	return path, file, suffix
}

// New filename is the full path filename like logs/test.log
// when rotate is 1 the rDay,rHour,rMinute are used;
// and when rotate is 2 the maxSize of bytes,nRotate are used
func New(filename string, rotate int, v1 int64, v2 int) (*Logger, error) {
	l := new(Logger)

	l.filename = filename
	p, f, s := GetPathFileName(filename)

	dealFileName := ""
	l.rotate = rotate
	if rotate == 1 {
		//change the file on the time of rHour and rMinute
		if v1 > 23 || v1 < 0 || v2 < 0 || v2 > 59 {
			return nil, errors.New("Time format is error for filename")
		}
		l.rHour = int(v1)
		l.rMinute = v2
		l.rDateTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), l.rHour, l.rMinute, 0, 0, time.Local)
		if l.rDateTime.After(time.Now()) {
			l.rDateTime = l.rDateTime.AddDate(0, 0, -1)
		}
		ts := time.Now().Format("_2006-01-02_15-04-05")
		dealFileName = p + f + ts + s
	} else if rotate == 2 {
		//change the file when the log file's size bigger than maxSize. total files count is nRotate
		if v1 < 0 || v2 < 1 {
			return nil, errors.New("Size format is error for filename")
		}
		l.maxSize = v1
		l.nRotate = v2
		l.cRotate = 0
		l.fnRotate = make([]string, l.nRotate, l.nRotate)
		for i := 0; i < l.nRotate; i++ {
			l.fnRotate[i] = p + f + strconv.Itoa(i) + s
		}
		dealFileName = l.fnRotate[0]
	}

	logFile, err := os.OpenFile(dealFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return nil, errors.New("Open the log file error. " + err.Error())
	} else {
		l.out = logFile
		l.file = logFile
	}

	l.closed = false
	l.bufChan = make(chan []byte, 10240)

	l.bPrefix = true

	//start logger to write
	go l.run()

	return l, nil
}

//SetPrefix set the bPrefix
func (l *Logger) SetPrefix(b bool) {
	l.bPrefix = b
}

func (l *Logger) Write(p []byte) (int, error) {
	var buf bytes.Buffer
	n, err := buf.Write(p)
	l.bufChan <- buf.Bytes()
	return n, err
}

// Output output the s to the Logger's pointer l
func (l *Logger) Output(level LOGLEVEL, s string) {
	var buf bytes.Buffer
	if l.bPrefix {
		buf.WriteString(time.Now().Format("[2006-01-02 15:04:05]"))
		buf.WriteString(leveName[level])
		if level > ERROR {
			funcName, file, line, ok := runtime.Caller(2)
			if ok {
				buf.WriteString(fmt.Sprintf("[%s:%d:%s]", file, line, runtime.FuncForPC(funcName).Name()))
			} else {
				fmt.Println("runtime.Caller error.")
			}
		}
		buf.WriteString(" ")
	}
	buf.WriteString(s)
	buf.WriteByte('\n')
	l.bufChan <- buf.Bytes()
}

// Info output information log
func (l *Logger) Info(format string, v ...interface{}) {
	l.Output(INFO, fmt.Sprintf(format, v...))
}

// Warn output warnning log
func (l *Logger) Warn(format string, v ...interface{}) {
	l.Output(WARN, fmt.Sprintf(format, v...))
}

// Error output Error log
func (l *Logger) Error(format string, v ...interface{}) {
	l.Output(ERROR, fmt.Sprintf(format, v...))
}

// Debug output debug log
func (l *Logger) Debug(format string, v ...interface{}) {
	l.Output(DEBUG, fmt.Sprintf(format, v...))
}

// Close close the log chan and the file's io stream
func (l *Logger) Close() {
	close(l.bufChan)
	l.file.Close()
}

func (l *Logger) run() {
	for {
		buf, ok := <-l.bufChan
		if !ok {
			break
		}
		l.doRotate()
		_, err := l.out.Write(buf)
		if err != nil {
			fmt.Println("Write log error.", err.Error())
		}
	}
}

func (l *Logger) doRotate() {
	if l.rotate == 1 {
		if time.Now().AddDate(0, 0, -1).After(l.rDateTime) {
			l.rDateTime = time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), l.rHour, l.rMinute, 0, 0, time.Local)
			if l.rDateTime.After(time.Now()) {
				l.rDateTime = l.rDateTime.AddDate(0, 0, -1)
			}
			p, f, s := GetPathFileName(l.filename)
			ts := time.Now().Format("_2006-01-02_15-04-05")
			logFile, err := os.OpenFile(p+f+ts+s, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				fmt.Println(err.Error())
				l.out = os.Stdout
			} else {
				l.out = logFile
				l.file.Close()
				l.file = logFile
			}
		}
	} else if l.rotate == 2 {
		info, err := l.file.Stat()
		if err != nil {
			fmt.Println("Rotate log error. get file info error. ", err.Error())
		} else {
			if info.Size() > l.maxSize {
				// get next filename
				l.cRotate++
				l.cRotate %= l.nRotate
				filename := l.fnRotate[l.cRotate]
				if l.nRotate > 1 { // if nRotate is 1, the log file will write out for ever
					os.Remove(filename)
					logFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
					if err != nil {
						fmt.Println(err.Error())
						l.out = os.Stdout
					} else {
						l.out = logFile
						l.file.Close()
						l.file = logFile
					}
				}
			}
		}
	}
}
