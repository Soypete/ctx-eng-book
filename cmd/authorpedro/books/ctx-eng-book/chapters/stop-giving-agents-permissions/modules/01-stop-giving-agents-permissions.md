# Stop Giving Agents Permissions

## Beat 1: The Permission Problem

Most AI systems give their agents too much access. A customer support agent can read all customer data. A code review agent can access every repository. A research agent can query the entire knowledge base. The assumption is that agents need broad access to be useful—that restricting them will limit their capabilities.

This assumption is wrong. Broad permissions create unreliable systems. An agent with access to everything will retrieve everything, contaminating its context with irrelevant, outdated, or sensitive data. It will make decisions based on information it shouldn't have. It will leak data it shouldn't see. The reliability problem is, at its root, a permission problem.

This chapter argues for a simple principle: agents should have exactly the permissions they need for the current task, no more. This is least privilege applied to AI systems. It requires rethinking authorization from the ground up—not as a gate that agents pass through, but as a capability that is granted, scoped, and revoked per operation.

**Building on Chapter 10:** Chapter 10 showed how different stores (wikis, databases, knowledge graphs) enforce permissions at the data layer. This chapter shows how agents obtain scoped capabilities that satisfy those store-level policies.

## Beat 2: Why Broad Permissions Fail

Consider a research agent that assists with technical due diligence. It needs access to certain repositories, documentation, and architecture diagrams. But the system gives it read access to the entire codebase. What happens?

First, context contamination. The agent retrieves thousands of irrelevant files. The retrieval system ranks them, but the agent still processes noise. Its answers become less precise, its reasoning harder to follow. The signal-to-noise ratio in its context determines its reliability—and broad permissions guarantee noise.

Second, dangerous retrieval. The agent with access to everything might retrieve sensitive data accidentally. Credentials in environment files. PII in logs. Internal communications in wikis. The model doesn't know what's sensitive. It retrieves what matches the query, regardless of consequences. You asked for "authentication code"—here's the file with the database password.

Third, cascade failures. An agent with broad permissions is a single point of compromise. If the agent is exploited—through prompt injection, retrieval manipulation, or context poisoning—it exposes everything. There's no blast radius limit. There's no defense in depth.

The fix is not to trust the agent more carefully. It's to give it less to work with.

## Beat 3: Least Privilege for Agents

Least privilege is a well-established security principle: every component should have only the permissions necessary to perform its function. In AI systems, this means the agent's permission scope should match its current task, not its potential scope.

Implementing least privilege for agents requires two changes to how we design systems:

First, task-scoped capabilities. Instead of giving an agent persistent access to a data store, grant it a capability for a specific operation. The capability specifies what data can be accessed, for how long, and with what constraints. After the operation completes, the capability expires.

Second, capability-based authorization. Rather than checking whether the agent has a role that permits access, verify that it presents a valid capability that authorizes the specific operation. The capability is the permission—unforgeable, scoped, and auditable.

This shifts the authorization model from "is the agent allowed?" to "does the agent have a capability for this?" The difference is subtle but critical. Role-based systems are binary—allowed or not. Capability-based systems are granular—exactly this data, exactly this operation, exactly this duration.

For context engineering, least privilege means the retrieval pipeline must enforce capability checks before returning context. The agent requests data. The system verifies the capability. The retrieval returns only permitted data. No prompt required, no model behavior assumed, no ambiguity.

## Beat 4: RBAC and Its Limits

Role-based access control is the dominant authorization model in enterprise systems. Users are assigned roles—engineer, manager, admin. Roles have permissions. The system checks: does the user's role have permission for this operation?

This model maps poorly to agents. An agent acting on behalf of a user inherits the user's role, but the user's role reflects their overall job function, not the specific task they're asking the agent to perform. An engineer might need access to production logs for debugging, but their engineer role shouldn't grant permanent production access for every agent operation.

The practical problem: RBAC is too coarse for agent authorization. A role grants permissions broadly and persistently. An agent needs permissions narrowly and temporarily.

The adaptation: map roles to task-specific capabilities. The context engineering system translates the user's role into a set of available capabilities for the requested task. The agent receives only the capabilities relevant to that task, not the full role permissions.

For example, a user with role=engineer requests agent assistance with production debugging. The system maps role=engineer + task=debugging to capability=read_logs:service=payment,timeframe=1h. The agent can read payment service logs from the last hour. Nothing else. The capability is scoped to the task.

This approach preserves RBAC's simplicity while adding the granularity agents need. Roles define what capabilities are available. The task determines which capabilities are granted. The capability determines what data is accessible.

## Beat 5: OAuth for Delegated Access

OAuth is the standard protocol for delegated authorization. A user authorizes an application to access their data without sharing credentials. The application receives an access token with scopes that define what it can do.

Agents are natural OAuth clients. When an agent acts on behalf of a user, it should obtain an OAuth token that represents the user's authorization for that specific operation. The token's scopes define what the agent can access.

The OAuth flow for agents works like this:

1. The user initiates a task through the agent
2. The system requests authorization with specific scopes—read:user_data,write:tickets
3. The user approves the scopes
4. The system issues an access token with those scopes
5. The agent uses the token to access resources
6. The token expires after the task completes

This is more secure than giving the agent persistent access. The token is scoped to the task, time-limited, and revocable. If the token is compromised, the blast radius is limited to its scope and lifetime.

For context engineering, OAuth tokens flow through the retrieval pipeline. The agent presents the token. The data store verifies the token's scopes. The retrieval returns only data permitted by those scopes. The token is the capability.

## Beat 6: OpenID Connect for Agent Identity

OAuth tells you what the agent can do. OpenID Connect tells you who the agent is acting as. It's an identity layer on top of OAuth—providing authentication in addition to authorization.

Agents need identity for two reasons. First, audit: you need to know which user an agent is acting on behalf of, for compliance and investigation. Second, attribute-based access: some authorization decisions depend on the user's attributes, not just the agent's scopes.

OpenID Connect provides an ID token alongside the access token. The ID token contains claims about the user: their ID, name, email, department, roles. The context engineering system uses these claims to make authorization decisions.

For example, a user with department=finance requests a report. The agent's OAuth token has scope=read:reports. The ID token claims include department=finance. The authorization layer checks: does scope=read:reports combined with department=finance permit access to this report? It might—for finance department reports, but not for engineering reports.

The practical pattern: the agent receives both an access token (what it can do) and an ID token (who it's acting as). The context engineering system uses both for authorization—the access token for scope checking, the ID token for attribute-based decisions. Provenance records both: this data was retrieved because token had scope X and user had attribute Y.

## Beat 7: Capability-Based Access

OAuth scopes are a form of capability, but they're limited. A scope like read:user_data is broad—it grants access to all user data. Capability-based access goes further, defining capabilities that are specific to the operation.

A capability is an unforgeable token that specifies:

- The resource: which specific data, not just which type
- The action: read, write, compute, or combinations
- The constraints: time window, row limits, field restrictions
- The grantor: who issued this capability and under what authority

Think of a capability as a signed URL with expiration. It's specific, scoped, and verifiable. The data store can verify the capability cryptographically and enforce its constraints without consulting an external authority.

For agents, capability-based access means the retrieval request includes a capability that precisely defines what's permitted. Not "read user data"—but "read user data for user_ids=[123, 456] from table=orders where created_at > 2024-01-01, for the next 30 minutes."

This granularity is essential for reliability. The agent cannot retrieve data outside its capability. The data store enforces the capability. There's no ambiguity, no reliance on prompt instructions, no way for the agent to exceed its authorization.

The implementation pattern: the context engineering system generates a capability for each retrieval operation, based on the user's authorization and the task requirements. The retrieval query includes the capability. The data store validates it and returns only permitted data.

## Beat 8: The Reliability Argument

Why does any of this matter for reliability?

Because dangerous retrieval is the primary failure mode for AI systems in production. The agent retrieves the wrong data—too much, too little, or the wrong kind—and produces unreliable output. The cause is almost always permission-related: the agent had access to data it shouldn't have, or the retrieval system couldn't filter effectively.

Capability-based access prevents dangerous retrieval by making it impossible. The agent can't retrieve data outside its capability because the data store enforces the capability. There's no prompt instruction that can be ignored. There's no retrieval logic that can be bypassed. The permission is structural, not advisory.

This is the reliability question answered at the authorization layer: how do we prevent dangerous retrieval? We stop giving agents broad permissions. We grant capabilities instead. We enforce those capabilities at the data store. We make dangerous retrieval structurally impossible.

The practical implementation: every retrieval operation carries a capability. The capability is generated from the user's authorization and the task's requirements. The data store validates the capability and returns only permitted data. Provenance records which capability authorized which retrieval.

This is not a prompt engineering problem. It's an authorization engineering problem. And it's the foundation of reliable AI systems.

(End of file)