package str

import (
	"testing"
	"github.com/franela/goblin"
)

func TestString(t *testing.T) {
	g := goblin.Goblin(t)
	g.Describe("left justify", func() {
		s10 := ""
		for i:=0; i < 10; i++{
			s10 += "--"
		}
		g.It("Repeat", func() {
			g.Assert(Repeat("--",10)).Equal(s10)
			g.Assert(len(Repeat(" ",10))).Equal(10)
		})
		g.It("Ljust", func() {
			g.Assert(Ljust("dev", 10, " ")).Equal("dev       ")
			g.Assert(Ljust("dev", 3, " ")).Equal("dev")
		})

		g.It("Rjust", func() {
			g.Assert(Rjust("dev", 10, " ")).Equal("       dev")
			g.Assert(Rjust("dev", 3, " ")).Equal("dev")
		})


		g.It("Center", func() {
			g.Assert(Center("dev", 10, "-")).Equal("---dev----")
			g.Assert(Center("dev", 7, " ")).Equal("  dev  ")
			g.Assert(Center("development", 7, " ")).Equal("development")
		})

		g.It("Capitalize", func() {
			g.Assert(Capitalize("development")).Equal("Development")
		})
	})

}
