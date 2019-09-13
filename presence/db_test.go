package presence

import (
	"time"

	"github.com/BorisBorshevsky/timemock"
	"github.com/nymtech/directory-server/models"
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
)

var _ = Describe("Presence Db", func() {
	Describe("listing network topology", func() {
		Context("when no presence has been registered by any node", func() {
			It("should return an empty topology object", func() {
				db := newPresenceDb()
				assert.Len(GinkgoT(), db.List(), 0)
			})
		})
	})
	Describe("for coconodes", func() {
		var (
			presence1 models.Presence
		//presence2 models.Presence
		)
		var db *db
		BeforeEach(func() {
			db = newPresenceDb()

			// Set up fixtures
			var mix1 = models.HostInfo{
				Host:   "foo.com:8000",
				PubKey: "pubkey1",
			}
			presence1 = models.Presence{
				HostInfo: mix1,
				LastSeen: timemock.Now().Unix(),
			}

			//var mix2 = models.HostInfo{
			//	Host:   "bar.com:8000",
			//	PubKey: "pubkey2",
			//}
			//presence2 = models.Presence{
			//	HostInfo: mix2,
			//	LastSeen: timemock.Now().Unix(),
			//}
		})
		Describe("adding presence", func() {
			Context("1st presence", func() {
				It("adds properly", func() {
					db.AddCocoNode(presence1)
				})
			})
		})
	})
	Describe("for mixnodes", func() {
		var (
			presence1 models.MixNodePresence
			presence2 models.MixNodePresence
		)
		var db *db
		BeforeEach(func() {
			db = newPresenceDb()

			// Set up fixtures
			var mix1 = models.MixHostInfo{
				HostInfo: models.HostInfo{
					Host:   "foo.com:8000",
					PubKey: "pubkey1",
				},
				Layer: 1,
			}
			presence1 = models.MixNodePresence{
				MixHostInfo: mix1,
				LastSeen:    timemock.Now().Unix(),
			}

			var mix2 = models.MixHostInfo{
				HostInfo: models.HostInfo{
					Host:   "bar.com:8000",
					PubKey: "pubkey2",
				},
				Layer: 2,
			}
			presence2 = models.MixNodePresence{
				MixHostInfo: mix2,
				LastSeen:    timemock.Now().Unix(),
			}
		})
		Describe("adding mixnode presence", func() {
			Context("1st presence", func() {
				It("returns the map correctly", func() {
					db.AddMix(presence1)
					assert.Len(GinkgoT(), db.List(), 1)
				})
				It("gets the presence by its public key", func() {
					db.AddMix(presence1)
					assert.Equal(GinkgoT(), presence1, db.get(presence1.PubKey))
				})
			})
			Context("adding two mixnode presences", func() {
				It("returns the map correctly", func() {
					db.AddMix(presence1)
					db.AddMix(presence2)
					assert.Len(GinkgoT(), db.List(), 2)
				})
				It("contains the correct presences", func() {
					db.AddMix(presence1)
					db.AddMix(presence2)
					assert.Equal(GinkgoT(), presence1, db.get(presence1.PubKey))
					assert.Equal(GinkgoT(), presence2, db.get(presence2.PubKey))
				})
			})
			Describe("Presences", func() {
				Context("more than 5 seconds old", func() {
					It("are not returned in the topology", func() {
						oldtime := time.Now().Add(time.Duration(-5 * time.Second)).Unix()
						presence1.LastSeen = oldtime
						db.AddMix(presence1)
						db.AddMix(presence2)
						assert.Len(GinkgoT(), db.List(), 1)
						assert.Equal(GinkgoT(), presence2, db.get(presence2.PubKey))
					})
				})
			})
		})
	})
})
