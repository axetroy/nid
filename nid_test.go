package nid

import (
  _ "github.com/lib/pq"
  "testing"
  "strings"
)

func Test_New(t *testing.T) {
  length := 12
  nid1 := New(length)

  nid2 := New(length)

  if len(nid1) != length || len(nid2) != length {
    t.Errorf("长度不一致")
  }

  if nid1 == nid2 {
    t.Errorf("两次id一致")
  }

  // 生成10000次，不重样
  // 不以0开头
  var (
    nidList [100000]string
  )
  for i := 0; i < 100000; i++ {
    nidList[i] = New(length)
  }

  dic := make(map[string]bool)

  for _, nid := range nidList {

    ss := strings.Split(nid, "")

    if ss[0] == "0" {
      t.Errorf("有一个ID以0开头： " + nid)
      break
    }

    if !dic[nid] {
      dic[nid] = true
    } else {
      t.Errorf("生产10000次，有重复的nid, id： " + nid)
      break
    }
  }

}
