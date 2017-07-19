package hello
import "testing"

func TestHello(t *testing.T) {
  name := "hoge"
  actual := Hello(name)
  expected := "Hello World! hoge"
  if actual != expected {
    t.Errorf("got %v\nwant %v", actual, expected)
  }
}
