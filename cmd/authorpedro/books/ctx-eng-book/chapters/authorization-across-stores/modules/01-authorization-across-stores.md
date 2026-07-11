# Authorization Across Data Stores

## Beat 1: The Authorization Question

Every production AI system eventually confronts a simple question: who is allowed to see what?

This question seems basic—almost trivial. Every software system that stores data must answer it. Yet AI systems consistently defer it, treating authorization as an afterthought to be solved with system prompts. The reasoning goes: "Tell the model what it can and can't access, and it will follow instructions."

This approach fails in production. Models are not access control systems. They are pattern matchers that generate likely continuations. Prompting them to enforce authorization is unreliable, inconsistent, and impossible to audit. The model might follow the instruction today and ignore it tomorrow. It might correctly restrict access in one context and leak sensitive data in another. You have no audit trail, no enforcement guarantee, and no fallback when the model gets it wrong.

Authorization is a data platform problem. It belongs at the storage layer, where it can be enforced consistently, logged reliably, and audited precisely. Context engineering must bridge the permission models of each underlying store—wikis, databases, knowledge graphs—into a coherent authorization layer that the AI system can trust.

This chapter argues that authorization across heterogeneous data stores is a core context engineering challenge. Different stores have different permission models. Context engineering must translate between them, enforce consistent policy, and provide provenance for every piece of context the model receives.

## Beat 2: Authorization Is Not a Prompt Problem

The temptation to solve authorization in the prompt is strong. It's immediate, flexible, and requires no architectural changes. A prompt like "Only access data the user has permission to see" seems to work—at least in testing.

The problems emerge in production. Models are non-deterministic in ways that affect security decisions. The same instruction produces different results depending on context, temperature, and model version. There's no enforcement mechanism when the model fails. There's no audit log of access decisions. The security boundary is imaginary—a text instruction that the model might follow or might not.

Consider a concrete failure: a customer service agent that accesses user data. The prompt instructs it to only retrieve data for the authenticated user. Under normal conditions, it works. But when the prompt is contaminated with adversarial context, or when the model hallucinates a permission it doesn't have, or when a retrieval bug returns data for the wrong user—the model doesn't know. It processes the wrong data and produces output that leaks it. The authorization failure happens at the data layer, but the prompt can't prevent it.

Now consider the same system with proper authorization at the storage layer. The database enforces row-level security. The wiki enforces page permissions. The knowledge graph enforces vertex-level access control. The model receives only the data it is permitted to access—no prompt required, no reliance on model behavior, no ambiguity.

This is the fundamental argument: authorization belongs in the data platform, not in the prompt. The prompt can express intent. The data platform must enforce it.

## Beat 3: ABAC for Agents

Attribute-Based Access Control provides the conceptual framework for agent authorization. Rather than fixed roles or static permissions, ABAC evaluates attributes of three things: the subject (the agent or user making the request), the resource (the data being accessed), and the action (read, write, delete, compute).

A subject attribute might be: role=engineer, department=infrastructure, clearance_level=3. A resource attribute might be: type=document, classification=internal, project=aurora. An action attribute might be: operation=read, requires_auth=true. The authorization decision combines these: does the subject's attributes satisfy the resource's requirements for the requested action?

This model maps directly to AI systems. The subject is the agent acting on behalf of a user. The resource is the context being retrieved—document, database row, knowledge graph vertex. The action is the retrieval itself.

In practice, ABAC for agents requires attribute assignment at request time. The context engineering system must know: who is this agent acting for, what are their attributes, what data are they requesting, and what attributes does that data have? The authorization layer evaluates these attributes and returns only permitted data.

This is more flexible than role-based access control. A role is a static assignment—engineer, manager, admin. An attribute is a property that can be combined, weighted, and evaluated dynamically. The agent requesting data about project=aurora might have attribute project=aurora, while another agent requesting the same data might have project=beta. The authorization layer evaluates these different attribute sets and returns different results.

The practical implication: your context engineering system must carry subject attributes through the entire retrieval pipeline. The agent's context includes not just what it knows, but who it's acting as and what attributes define that identity.

## Beat 4: Wiki Permissions

Wikis store unstructured text—meeting notes, design documents, RFCs, knowledge base articles. Wiki permission models vary, but the common pattern is page-level or namespace-level access control.

A namespace is a grouping mechanism. All pages under namespace=engineering might be readable by role=engineer but writable only by role=engineering-lead. A page might override namespace defaults with specific permissions—granting a specific user read access to a sensitive document.

For context engineering, the challenge is bridging wiki permissions into the retrieval pipeline. When the agent requests "relevant documents about the authentication system," the retrieval system must:

1. Identify which wiki pages match the query
2. For each page, check whether the requesting subject has read permission
3. Filter the results to only permitted pages
4. Annotate the returned context with provenance (which wiki, which page, which namespace)

This requires the context engineering system to understand wiki permission structures and enforce them during retrieval. You can't retrieve everything and filter in the prompt—the model can't evaluate permissions it doesn't see. The retrieval must be permission-aware from the start.

Most wikis expose permissions through APIs or database queries. The context engineering layer queries these, combines them with subject attributes, and constructs a permission filter for the retrieval query. The agent never sees documents it shouldn't access.

## Beat 5: Row-Level Security in Databases

Relational databases have mature row-level security (RLS). PostgreSQL, SQL Server, and other systems support security policies that filter rows based on the executing user's identity.

The mechanism is straightforward: a security policy defines a predicate that must evaluate to true for a row to be visible. The predicate can reference the current user, session attributes, or external data. When the user queries the table, the database automatically injects the predicate and returns only permitted rows.

For context engineering, RLS is powerful because it's enforced at the database layer. The context retrieval query doesn't need to manually filter—the database does it. A query like SELECT * FROM user_orders returns only rows for the current user, automatically, regardless of what the context engineering system requests.

This integrates cleanly with agent authorization. The context engineering system establishes a database connection with the agent's identity—the user they're acting on behalf of. Every query runs with that identity, automatically filtered by RLS policies. The retrieved data is inherently permitted.

The practical pattern: the agent's context includes a database connection authenticated as the user. Retrieval queries use that connection. Row-level security policies enforce access control at the database. The context engineering system doesn't need to implement authorization logic—it delegates to the database's proven security model.

This is the data platform approach in action. Authorization is not a prompt instruction. It's a database enforcement mechanism that the context engineering system leverages.

## Beat 6: Knowledge Graph Access Control

Knowledge graphs represent entities and relationships as vertices and edges. Access control must apply at both levels—controlling which vertices a subject can read and which edges they can traverse.

Vertex-level access control is straightforward: each vertex has attributes (like classification, project, sensitivity), and the authorization layer evaluates whether the subject can read vertices with those attributes. This mirrors wiki page permissions.

Edge-level access control is more subtle. An agent might have permission to read vertex A and vertex B individually, but not permission to know about the relationship between them. The edge represents a connection—A employs B, A acquired B, A communicates with B. Some relationships are sensitive. The knowledge graph must control whether the subject can see the edge, not just the vertices.

In practice, knowledge graph access control requires:

- Vertex attributes that describe sensitivity and access requirements
- Edge attributes that describe relationship sensitivity
- A traversal policy that respects both vertex and edge permissions
- A query mechanism that filters during traversal, not after

When the agent queries "who worked on the authentication system," the knowledge graph traversal starts at vertex=authentication_system and follows edges to related vertices. Each edge traversal checks whether the subject has permission to see that relationship. Each vertex retrieval checks whether the subject has permission to read that entity. The result is a permission-aware subgraph that the agent can safely use as context.

This is more complex than wiki or database permissions because the graph structure creates implicit data. Knowing that A connects to B is itself sensitive information. Knowledge graph access control must account for this.

## Beat 7: Provenance for Audit

Authorization without audit is incomplete. Every piece of context the model receives should be traceable: where did this data come from, who authorized access, what attributes were evaluated?

Provenance tracks this. It's the metadata layer that records, for each retrieved context item:

- The source store (wiki, database, knowledge graph)
- The specific resource (page, table, vertex)
- The subject attributes that authorized access
- The timestamp of access
- The query that triggered retrieval

This metadata serves two purposes. First, it enables security auditing. When investigating a potential data leak, you can trace what data the model received and why it was permitted. Second, it enables reliability debugging. When the model produces incorrect output, you can trace what context it received and whether authorization filtering behaved as expected.

Provenance is not optional in production systems. It's the difference between "the model somehow accessed this data" and "the model accessed this data because subject=agent-7 had attribute project=aurora and the row had attribute project=aurora, satisfying the policy."

The context engineering system must record provenance for every retrieval operation. This adds overhead, but it's essential for operating secure, auditable AI systems at scale.

## Beat 8: Bridging Permission Models

The central challenge of authorization across data stores is bridging different permission models into a coherent system. Wikis use page-level permissions. Databases use row-level security. Knowledge graphs use vertex and edge permissions. Each store has its own model, its own API, its own enforcement mechanism.

Context engineering must abstract across these models. The approach is to define a canonical authorization model—ABAC at the context engineering layer—and translate it to each store's native model:

- For wikis: translate subject attributes to page permission checks
- For databases: translate to RLS policies that execute with subject identity
- For knowledge graphs: translate to vertex and edge filters during traversal

The subject attributes flow through the entire pipeline. The context engineering system carries the agent's identity and attributes to each store, applies the store-specific authorization mechanism, and aggregates the permitted results with provenance metadata.

This is the data platform approach to authorization. Instead of hoping the prompt enforces security, you build a coherent authorization layer across heterogeneous stores. The model receives only permitted context. Every retrieval is auditable. Every decision is traceable.

The reliability question—who is allowed to see what?—has a precise answer: the data platform decides, store by store, and the context engineering system bridges the models.

This is not a prompt engineering problem. It's a systems engineering problem. And it's one you must solve before your AI system operates on sensitive data.

(End of file)