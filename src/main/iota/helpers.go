package iota

import (
  "math/rand"
  "time"
  "log"

  "github.com/iotaledger/iota.go/mam/v1"
  "github.com/iotaledger/iota.go/consts"
  "github.com/iotaledger/iota.go/api"
)

const endpoint = "https://nodes.devnet.iota.org"

// difficulty of the proof of work required to attach a transaction on the tangle
const mwm = 9

// how many milestones back to start the random walk from
const depth = 3

func GenerateSeed() string {
  var seed string
  alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ9"
  r := rand.New(rand.NewSource(time.Now().Unix()))

  for i := 0; i < 81; i++ {
    seed += string(alphabet[r.Intn(len(alphabet))])
  }
  return seed
}

func GetTransmitter(t *mam.Transmitter, api *api.API, cm mam.ChannelMode) *mam.Transmitter {
  switch {
    case t != nil:
      return t
    default:
      transmitter := mam.NewTransmitter(api, GenerateSeed(), uint64(mwm), consts.SecurityLevelMedium)
      if err := transmitter.SetMode(cm, sideKey.Get()); err != nil {
        log.Fatal(err)
      }
      return transmitter
  }
}
