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

## Beat 3: Wrong Tool Calls Are Missing Context

The first failure stream is too many tool calls—pulling all the data from MCPs, searching all the files and all the repos, getting the picture before understanding the question. The agent retrieves information without direction, and then incorrect information and incorrect assumptions proliferate in the session. Tool selection requires contextual understanding of what's available, what the user actually wants, and what information would be relevant versus noise. Without scoped retrieval, agents drown in data and surface the wrong conclusions.

This also happens with AI-generated skills—tools that were auto-generated without careful descriptions, boundaries, or scoped purposes. When the tool definition itself lacks proper context about what it's for and when to use it, the agent has no chance of selecting correctly.

## Beat 4: Agent Loops Are Missing Exit Criteria

Loops happen when agents aren't given correct context to achieve a goal, or when tool calls don't match what the model can actually do—calling `directory` when the model only understands `directories`, for example. This is where guardrails from the Forge repo come in. Another common failure: not limiting turns. When you need many SQL queries to get information—one to get an ID, another to get a second ID, and so on—and a column name is misspelled, the model will "learn" and self-heal. But a better workflow with proper context upfront would have accomplished it in fewer turns. The loop persists because the model doesn't know when success has been achieved or what information would break the cycle.

## Beat 5: Permission Failures Are Unconstrained Context

When agents first started automating jobs, there was a lot of user impersonation—agents acted as the user via prompt instructions, inheriting tokens and credentials. This opened the door to prompt injection: someone could convince the bot to act on behalf of a person they are not. Access and permissions should be scoped to the user that instigated the agent or session. The agent should only have access to what the user themselves has access to—we can use existing authorization structures for data access instead of relying on prompts to convey permissions.

## Beat 6: Cost Overruns Are Unmeasured Context

Larger context means larger bills—AI is billed by token. The money optimization question is shorter agent loops. The failures I mostly see are around compaction: how do we shorten context without losing relevant information? Without context budget controls, systems consume unlimited tokens, make redundant calls, or retrieve unnecessary data. The context failure here is failing to measure, limit, and optimize what gets included.