package goutil

import (
	"crypto/md5"
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"net"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// The same as php ip2long
func IP2long(ipstr string) uint32 {
	ip := net.ParseIP(ipstr)
	if ip == nil {
		return 0
	}
	ip = ip.To4()
	return binary.BigEndian.Uint32(ip)
}

// The same as php long2ip
func Long2IP(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)
	return ip.String()
}

// Generate Md5 of the given string
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// The same microtime as php
func Microtime() float64 {
	return float64(time.Now().UnixNano()) / 1000000000.0
}

// Delete the given key in slice
func DeleteSliceItem(key int, items interface{}) {
	refValue := reflect.ValueOf(items)
	refType := reflect.TypeOf(items)

	if refValue.Kind() == reflect.Ptr {
		refValue = refValue.Elem()
		refType = refType.Elem()
	}

	if refType.Kind() != reflect.Slice {
		panic("the items must be array or slice")
	}

	length := refValue.Len()
	if 0 > key || key > length {
		return
	}

	refValue.Set(reflect.AppendSlice(
		refValue.Slice(0, key),
		refValue.Slice(key+1, length)),
	)
}

// Return part of a string
//
// Examples:
// string   start length return
// "abcdef" 0     1      a
// "abcdef" 1     2      bc
// "abcdef" -2    1      e
// "abcdef" -2    0      ef
// "abcdef" 1     -2     bcd
// "abcdef" -3    -2     d
// "abcdef" -2    -4     ""
// "abcdef" -20   -4     ab
// "abcdef" -20   -10    ""
func SubStr(str string, start, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl + start
		if start < 0 {
			start = 0
		}
	}

	if 0 == length {
		end = rl
	} else if 0 > length {
		end = rl + length
		if end < 0 {
			end = 0
		}
	} else {
		end = start + length
		if end > rl {
			end = rl
		}
	}

	if start > end {
		return ""
	}

	return string(rs[start:end])
}

// Rounds a float
func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

// Remove the html tags
func StripHtmlTags(src string) string {
	//all to lower case
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//remove STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//remove SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//remove all html
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "")

	//remove all duplicate line break
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	return strings.TrimSpace(src)
}

// Convert any type of int or float or string to int64
// Use case, like timestamp
func ToInt64(val interface{}) (int64, error) {
	switch v := val.(type) {
	case int:
		return int64(v), nil
	case int16:
		return int64(v), nil
	case int32:
		return int64(v), nil
	case int64:
		return v, nil
	case float32:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case string:
		i, err := strconv.Atoi(v)
		return int64(i), err
	}

	return 0, errors.New("cannot convert to int64")
}
