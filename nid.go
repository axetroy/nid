/**
Author: Axetroy
Github: https://github.com/axetroy
 */

package nid

import (
  "math/rand"
  "strconv"
  "time"
)

var seed = time.Now().UnixNano()

var rnd = rand.New(rand.NewSource(seed))

/**
Generate a new random number id
 */
func New(length int) (nid string) {
  nid = ""

  for i := 0; i < length; i++ {
    var (
      r1 int
    )

    r1 = rnd.Intn(9)

    if i == 0 {
      for r1 == 0 {
        r1 = rnd.Intn(9)
      }
    }

    nid += strconv.Itoa(r1)
  }
  return
}

func SetSeed(s int64) {
  seed = s
  rnd = rand.New(rand.NewSource(seed))
}
