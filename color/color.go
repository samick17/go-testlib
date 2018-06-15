package color

import (
    "fmt"
    "regexp"
)

const (
    Red = "\x1b[0;31m"
    Green = "\x1b[0;32m"
    Yellow = "\x1b[0;33m"
    Blue = "\x1b[0;34m"
    Purple = "\x1b[0;35m"
    Cyan = "\x1b[0;36m"
    White = "\x1b[0;37m"
    BrightBlack = "\x1b[0;90m"
    BrightRed = "\x1b[0;91m"
    BrightGreen = "\x1b[0;92m"
    BrightYellow = "\x1b[0;93m"
    BrightBlue = "\x1b[0;94m"
    BrightMagenta = "\x1b[0;95m"
    BrightCyan = "\x1b[0;96m"
    BrightWhite = "\x1b[0;97m"
    EndOfColor = "\x1b[0m"
)

var Colors = map[string]string {
    "Red": Red,
    "Green": Green,
    "Yellow": Yellow,
    "Blue": Blue,
    "Purple": Purple,
    "Cyan": Cyan,
    "White": White,
    "BrightBlack": BrightBlack,
    "BrightRed": BrightRed,
    "BrightGreen": BrightGreen,
    "BrightYellow": BrightYellow,
    "BrightBlue": BrightBlue,
    "BrightMagenta": BrightMagenta,
    "BrightCyan": BrightCyan,
    "BrightWhite": BrightWhite,
    "EndOfColor": EndOfColor,
}

func ToColor(color string, text string) string {
    if val, ok := Colors[color]; ok {
        return val + text + EndOfColor
    } else {
        return text
    }
}

func CreateText(message string, args ... interface{}) *Text {
    text := &Text{value: message}
    var msg string
    msg = message
    for i, arg := range args {
        valOfIndex := fmt.Sprint(i)
        var re = regexp.MustCompile(`\{`+valOfIndex+`\}`)
        valOfArg := fmt.Sprint(arg)
        msg = re.ReplaceAllLiteralString(msg, valOfArg)
    }
    text.value = msg
    return text
}

type Text struct {
    value string
    color string
}

func (t *Text) String() string {
    return ToColor(t.color, t.value)
}

func (t *Text) Color(color string) *Text {
    if _, ok := Colors[color]; ok {
        t.color = color
    }
    return t
}