# Evaluating AI Systems

## Beat 1: The Reliability Question

How do we prove improvement?

This is the question every AI systems team must answer. In traditional software, it's straightforward: you write a test, the test passes, you ship. The test checks deterministic behavior. Given the same input, you get the same output. You can measure coverage, assert on return values, and know exactly what broke when something fails.

AI doesn't work this way. Given the same input, you might get different outputs. The same prompt that produces a correct answer today might produce a slightly different answer tomorrow—not because something broke, but because the model is sampling from a probability distribution. Your "test" that passed last week might fail this week not because code changed, but because the model's behavior drifted.

This is the fundamental challenge of AI evaluation: you're evaluating probabilistic behavior, not deterministic outputs. And that changes everything about how you build test suites, measure reliability, and prove that your system is improving.

## Beat 2: Evals as Automated Judgment

An eval is an automated judgment about AI output quality. Unlike a unit test that asserts on exact values, an eval judges whether the output meets some standard—correctness, coherence, usefulness, safety.

The most common approach is LLM-as-judge: you ask another AI model to evaluate the output. You provide a prompt that describes the criteria, show the input and output, and get back a score or judgment. This works because evaluating text is something language models do well, and it scales in a way that human review cannot.

But LLM-as-judge has limits. Judges have their own biases—they might prefer verbose outputs, or be too lenient, or fail to catch subtle errors. The judge model itself can drift. You're evaluating probabilistic behavior with another probabilistic system. This is why golden datasets matter: a set of inputs with known-correct outputs that you can use to measure whether your eval is actually working. If your judge says a correct answer is wrong, your golden dataset catches that.

Golden datasets also solve the regression problem. You build a collection of representative inputs—the edge cases, the common paths, the failure modes you've seen in production—and you run your system against them over time. When outputs change, you can see whether they're better or worse. The golden dataset is your ground truth, the fixed point against which you measure change.

Effective evals share common characteristics: they're automated, reproducible, and correlated with real-world performance. An eval that doesn't predict production behavior is just busywork. The best evals are the ones that catch real bugs—the ones that would have reached users if you hadn't run the eval first.

## Beat 3: Benchmarks—When to Use, When to Ignore

Benchmarks are standardized eval sets that let you compare performance across systems. MMLU measures general knowledge. HumanEval measures code generation. SWE-Bench measures software engineering capability. The promise is simple: run the benchmark, get a score, know where you stand.

This is useful when you're choosing between models. If you need code generation capability, you compare HumanEval scores across available models and pick the best. Benchmarks give you a common yardstick for model selection.

But benchmarks have a crucial limitation: they measure capability, not reliability in your specific system. A model that scores well on HumanEval might fail consistently on the particular code patterns your application uses. The benchmark measures what the model can do in abstract; your eval measures what your system actually does.

Ignore benchmarks when they're driving design decisions they shouldn't influence. If you're optimizing for benchmark scores instead of user experience, you've lost the plot. The benchmark is a signal, not a target. Your system serves users, not test sets.

Use benchmarks for model selection and for tracking capability trends over time. If a new model version improves your primary benchmark, that's a signal to test it in your system. But always validate with your own evals on your own data.

## Beat 4: Regression Testing—Comparing Outputs Across Versions

Regression testing for AI systems means comparing outputs across versions. You take the same input, run it through the current system and the previous version, and measure differences. This is how you answer the reliability question: are we improving or degrading?

The challenge is that some difference is expected. AI outputs vary. A change in whitespace, a reworded explanation, a different but equally correct answer—these aren't regressions. Your regression test needs to distinguish meaningful changes from acceptable variance.

One approach is output comparison with allowed variation. You define what counts as equivalent: same structured data, same key facts, same tool calls. Then you assert on those invariants and allow flexibility on everything else. If your system returns a list of search results, you assert that the relevant results are present, not that they're in a specific order.

Another approach is behavioral regression: you don't compare outputs directly, you compare outcomes. Does the user get a correct answer? Does the tool selection succeed? Does the workflow complete? These are the things that matter for reliability. A different path to the same correct outcome isn't a regression.

Build regression suites from production failures. Every bug that reaches users should produce a test case that goes into your regression suite. Over time, your regression suite becomes a living record of failure modes—the things your system has already broken and been fixed. Running this suite before every deployment is how you prove that you're not reintroducing known problems.

## Beat 5: Reliability Metrics

Beyond evals and regression tests, you need metrics that measure system-level reliability. These are the numbers that answer: is the system working?

Latency matters because AI is slow. Every token generated takes time, and longer outputs mean more latency. But latency isn't just about speed—it's about reliability under load. If latency spikes during peak traffic, that's a reliability failure. Track p50, p95, and p99 latency. Watch for degradation across model versions.

Error rate is the percentage of requests that produce an error—timeouts, API failures, invalid outputs. This is the most basic reliability metric. If your error rate is climbing, something is breaking. Track it by model version, by endpoint, by time of day. A sudden spike in error rate means something is wrong, even if you don't yet know what.

Retrieval accuracy measures whether your context system is giving the model what it needs. If your retrieval system returns irrelevant documents, the model can't produce good answers. Measure recall: of the documents that would make the answer correct, how many did retrieval return? Measure precision: of the documents retrieved, how many are actually relevant? Both matter. Poor recall means missing information. Poor precision means overwhelming the model with noise.

Context quality is harder to measure but critical. Is your prompt structure effective? Are you including the right instructions? Is your examples set representative? You can measure this indirectly through downstream metrics—if improving the prompt improves output quality, the prompt was lacking. But context quality is something you often need to evaluate manually, especially when you're iterating on prompt design.

## Beat 6: The Eval Pipeline

Think of your eval system as a pipeline with three stages, each testing at a different granularity.

Unit tests for tools test individual components in isolation. Does your retrieval function return relevant documents? Does your prompt template render correctly? Does your tool schema validate? These are deterministic checks on discrete functions. They catch the bugs that are easy to catch.

Integration tests for workflows test the composition of components. Does the full pipeline—retrieve, prompt, generate, parse—produce correct outputs? These tests exercise the system end-to-end but with fixed inputs and known expected outputs. They catch integration bugs, the failures that happen when components are wired together wrong.

Evals for outputs test the AI behavior itself. Given a prompt and context, is the output correct, useful, and safe? This is where LLM-as-judge lives, where golden datasets matter, where you're judging probabilistic behavior. Evals run on every deployment and catch the failures that unit and integration tests miss—the ones that come from the model itself, not from your code.

The pipeline runs automatically. Every code change triggers unit tests. Every model change triggers integration tests. Every deployment triggers evals. Results flow into dashboards that show reliability over time. When evals regress, you don't ship. When latency spikes, you investigate. When error rates climb, you alert.

This is how you prove improvement: you measure, you compare, you track over time. The numbers tell the story. You ship when the numbers are good, and you don't ship when they're not.

## Beat 7: Evaluation Is Continuous

The hard truth about AI evaluation is that there's no finish line. Models update. User behavior changes. Failure modes evolve. Your eval suite is never complete—it grows with your system, incorporating every bug you've seen, every edge case you've encountered, every regression you've suffered.

Build the culture around this. Treat evals as first-class artifacts, as important as the code itself. When something breaks in production, write a test for it. When a new failure mode appears, add it to your golden dataset. The eval suite is your system's memory of everything that's gone wrong—and your best defense against it happening again.

Reliable AI systems aren't built by hoping for the best. They're built by measuring everything, comparing rigorously, and proving through data that they're improving. The reliability question has an answer: you prove it with evals.