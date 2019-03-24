package foobargo

import ("testing")


func TestEcho(t *testing.T){

	got := Echo("Hello")
	want := "HELLO"

	if got != want {
		t.Errorf("Unexpected result. Expected [%s] but got [%s]", want, got)
	}

}
