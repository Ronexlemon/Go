package randomGreeting

import (
	"testing"
	"regexp"
)

func TestRandomGreeting(t *testing.T){
	name:="Ronex"
	want:= regexp.MustCompile(`\b`+name+`\b`)
	msg,err:= RandomGreeting("Ronex")

	if !want.MatchString(msg) || err != nil {
        t.Fatalf(`RandomGreeting("Ronex") = %q, %v, want match for %#q, nil`, msg, err, want)
    }

}

func TestEmptyRandomGreeting(t *testing.T){
	msg,err:= RandomGreeting("")
	if msg != "" || err == nil {
        t.Fatalf(`RandomGreeting("") = %q, %v, want "", error`, msg, err)
    }
}