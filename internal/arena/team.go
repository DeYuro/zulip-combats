package arena

import "github.com/deyuro/zulip-combats/internal/figther"

type Team struct {
	Members []figther.Fighter[uint]
}
