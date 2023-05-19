package snowflake

import "github.com/bwmarrin/snowflake"

type Generator interface {
	Generate() snowflake.ID
}

type Provider interface {
	Generator() Generator
}
