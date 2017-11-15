package nid

import (
  _ "github.com/lib/pq"
  "testing"
  "time"
)

func Test_New(t *testing.T) {
  nid1 := New(8)

  time.Sleep(1)

  nid2 := New(8)

  if len(nid1) != 8 || len(nid2) != 8 {
    t.Errorf("长度不一致")
  }

  if nid1 == nid2 {
    t.Errorf("两次id一致")
  }

}
