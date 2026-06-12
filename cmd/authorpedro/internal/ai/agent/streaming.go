package agent

import (
	"context"

	"github.com/soypete/pedro-agentware/go/llm"
)

type StreamingBackend struct {
	llm.Backend
	OnChunk func(string)
}

func (s *StreamingBackend) StreamComplete(ctx context.Context, req *llm.Request, onChunk func(llm.Response)) (*llm.Response, error) {
	originalOnChunk := onChunk
	wrappedOnChunk := func(chunk llm.Response) {
		if s.OnChunk != nil && len(chunk.Content) > 0 {
			s.OnChunk(chunk.Content)
		}
		if originalOnChunk != nil {
			originalOnChunk(chunk)
		}
	}
	return s.Backend.StreamComplete(ctx, req, wrappedOnChunk)
}

func WrapBackendForStreaming(backend llm.Backend, onChunk func(string)) llm.Backend {
	return &StreamingBackend{
		Backend: backend,
		OnChunk: onChunk,
	}
}
