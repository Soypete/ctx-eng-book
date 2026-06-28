# Forge — Tool-Calling Reliability Framework

## Overview

A Python framework for self-hosted LLM tool-calling and multi-step agentic workflows. Provides a reliability layer that sits inside an agentic loop and makes tool calls reliable.

**Repository:** https://github.com/antoinezambelli/forge
**Stars:** 2.1k
**License:** MIT

## Key Capabilities

### Three Usage Modes

1. **Proxy Server** — Drop-in proxy between clients and local models. Supports OpenAI chat-completions and Anthropic Messages APIs. Point clients (opencode, Continue, aider, Claude Code) at the proxy and forge applies guardrails transparently.

2. **WorkflowRunner** — Define tools, pick a backend, run structured agent loops. Manages full lifecycle: system prompts, tool execution, context compaction, guardrails.

3. **Guardrails Middleware** — Composable reliability stack for use inside your own orchestration loop.

### Guardrail Features

- **Response validation** — validates tool calls against the tools array
- **Rescue parsing** — extracts structured tool calls from malformed output (Mistral `[TOOL_CALLS]`, Qwen `<tool_call>` XML, fenced JSON)
- **Retry loop with error tracking** — retries inference on validation failures with corrective nudges
- **Synthetic `respond` tool** — injects a tool for small models that can't reliably choose between text and tool calls
- **Step enforcement** — enforces required steps and prerequisites (WorkflowRunner only)
- **Context compaction** — rolling window management (WorkflowRunner only)

### Backends Supported

- Ollama
- llama-server (llama.cpp)
- Llamafile
- vLLM
- Anthropic

## Evaluation Results

- Takes an 8B local model from single digits to **84%** across 26-scenario eval suite
- Lifts Sonnet 4.6 from **85% to 98%** on same workload
- Eval suite: 26 scenarios (OG-18 baseline + 8 advanced_reasoning)

## Paper

Published in ACM/IEEE:
> Zambelli, A. *Forge: A Reliability Layer for Self-Hosted LLM Tool-Calling.* https://doi.org/10.1145/3786335.3813193

## Go Port

Ported to Go in [incode-agentware](https://github.com/incode-agentware/forge):
- Same guardrail functionality in Go
- Useful for Go-based agent frameworks

## Relevance to Context Engineering

- Demonstrates that tool-calling reliability requires **structured constraints**, not just better prompts
- Shows value of **guardrails** as a layer between model and tools
- Validates the principle that context engineering includes **validation and error recovery**
- Useful for examples in the book on building reliable tool-using agents