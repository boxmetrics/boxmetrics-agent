package info

import (
	// "github.com/boxmetrics/boxmetrics-agent/internal/pkg/errors"
	"strconv"
	"time"
)

type info interface {
	String() string
	Text() string
	Format() string
}

type byte float64

func (b byte) String() string {
	return strconv.FormatFloat(float64(b), 'f', 2, 64)
}

func (b byte) Text() string {
	s := b.String()

	return s + "B"
}

func (b byte) Format() string {
	switch {
	case b >= pb:
		v := b / pb
		return v.String() + "PB"
	case b >= tb:
		v := b / tb
		return v.String() + "TB"
	case b >= gb:
		v := b / gb
		return v.String() + "GB"
	case b >= mb:
		v := b / mb
		return v.String() + "MB"
	case b >= kb:
		v := b / kb
		return v.String() + "KB"
	}
	return b.Text()
}

type byteSpeed float64

func (bs byteSpeed) String() string {
	return strconv.FormatFloat(float64(bs), 'f', 2, 64)
}

func (bs byteSpeed) Text() string {
	s := byte(bs).Text()

	return s + "/S"
}

func (bs byteSpeed) Format() string {
	b := byte(bs)
	switch {
	case b >= pb:
		v := b / pb
		return v.Format() + "/S"
	case b >= tb:
		v := b / tb
		return v.Format() + "/S"
	case b >= gb:
		v := b / gb
		return v.Format() + "/S"
	case b >= mb:
		v := b / mb
		return v.Format() + "/S"
	case b >= kb:
		v := b / kb
		return v.Format() + "/S"
	}
	return bs.Text()
}

type percent float64

func (p percent) String() string {
	return strconv.FormatFloat(float64(p), 'f', 2, 64)
}

func (p percent) Text() string {
	s := p.String()

	return s + "%"
}

func (p percent) Format() string {
	return p.Text()
}

type frequency float64

func (f frequency) String() string {
	return strconv.FormatFloat(float64(f), 'f', 2, 64)
}

func (f frequency) Text() string {
	s := f.String()

	return s + " MHz"
}

func (f frequency) Format() string {
	return f.Text()
}

type jiffy float64

func (j jiffy) String() string {
	return strconv.FormatFloat(float64(j), 'f', 4, 64)
}

func (j jiffy) Text() string {
	s := j.String()

	return s + " jiffies"
}

func (j jiffy) Format() string {
	s := strconv.FormatFloat(float64(j/100), 'f', 4, 64) + "s"

	d, _ := time.ParseDuration(s)

	return d.String()
}

type number int

func (n number) String() string {
	return strconv.Itoa(int(n))
}

func (n number) Text() string {
	return n.String()
}

func (n number) Format() string {
	return n.Text()
}

type systemtime int64

func (st systemtime) String() string {
	return strconv.FormatInt(int64(st), 10)
}

func (st systemtime) Text() string {
	str := st.String()

	return str + "s"
}

func (st systemtime) Format() string {
	t := time.Unix(int64(st), 0)

	return t.String()
}

type second float64

func (s second) String() string {
	return strconv.FormatFloat(float64(s), 'f', 4, 64)
}

func (s second) Text() string {
	str := s.String()

	return str + "s"
}

func (s second) Format() string {
	d, _ := time.ParseDuration(s.Text())

	return d.String()
}

type millisecond float64

func (m millisecond) String() string {
	return strconv.FormatFloat(float64(m), 'f', 4, 64)
}

func (m millisecond) Text() string {
	s := m.String()

	return s + "ms"
}

func (m millisecond) Format() string {
	d, _ := time.ParseDuration(m.Text())

	return d.String()
}

type temperature float64

func (t temperature) String() string {
	return strconv.FormatFloat(float64(t), 'f', 2, 64)
}

func (t temperature) Text() string {
	s := t.String()

	return s + "°C"
}

func (t temperature) Format() string {
	
	return t.Text()
}

const (
	b byte = 1 << (10 * iota)
	kb
	mb
	gb
	tb
	pb
)

type numberSlice []number

func (r numberSlice) Convert() []info {
	infos := make([]info, len(r))
	for i, v := range r {
		infos[i] = v
	}

	return infos
}

func newNumberSlice32(slice []int32) numberSlice {
	var ns numberSlice

	for _, val := range slice {
		ns = append(ns, number(val))
	}

	return ns
}


func convertInfoSlice(slice []info, method string) []string {
	var ret []string

	for _, inf := range slice {
		switch method {
		case "string":
			ret = append(ret, inf.String())
		case "text":
			ret = append(ret, inf.Text())
		case "format":
			ret = append(ret, inf.Format())
		default:
			ret = append(ret, "")
		}
	}

	return ret
}