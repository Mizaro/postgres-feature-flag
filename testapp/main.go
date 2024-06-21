package main

import (
	"context"
	"fmt"
	"github.com/Mizaro/postgrest-feature-flag/pkg/postgrestfeatureflags"
	"github.com/open-feature/go-sdk/openfeature"
)

func main() {
	// Register your feature flag postgrestfeatureflags
	_ = openfeature.SetProvider(postgrestfeatureflags.NewPostgrestProvider("http://localhost:3000/"))
	// Create a new client
	client := openfeature.NewClient("app")
	// Evaluate your feature flag
	v2Enabled, _ := client.BooleanValue(
		context.Background(), "v2_enabled", true, openfeature.EvaluationContext{},
	)
	// Use the returned flag value
	if v2Enabled {
		fmt.Println("v2 is enabled")
	} else {
		fmt.Println("v2 is disabled")
	}
}
