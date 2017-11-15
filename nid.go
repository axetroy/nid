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

/**
Generate a new random number id
 */
func New(length int) (nid string) {
  nid = ""
  rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

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
