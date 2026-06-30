# Every Failure Is a Context Failure

## Beat 1: The Demo-to-Production Gap

This is normal in data work. The problem is data is unpredictable. And when you add user interaction, you add more unpredictability.

For example: if you ask someone what a strawberry tastes like, you'll get several different answers. The same with AI. Ask it to write code and solve a problem—there are always many ways to solve a problem. Data is probabilistic. It's full of statistical distributions. This is why we deal with AI differently.

In demos, we want to limit outcomes to make it look good, so we use prompts and theoretical constraints. We're targeting only a small sample set. The problem is in the real world, we can and should put up guardrails for when we run into outliers.

## Beat 2: Hallucinations Are Missing Context

A hallucination is like a bug in software. What is it? When the model does something the user doesn't expect—especially as we use these models for automation.

Smaller "dumber" models are more likely to regurgitate erroneous information when it comes to spitting out facts. But when it comes to predicting outcomes based on provided information or tool calls, they do really well with good prompting.

The key insight: if you provide all three—semantics, pragmatics, and data—the model will act with minimum probability of error. And that's getting better. But if you do not provide the correct amount of all three, it will mess up almost assuredly.

This is the first pillar of context engineering: the model needs the right information, expressed the right way, in the right amount. Fail on any dimension, and hallucinations follow.