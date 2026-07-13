# Lucene and Elasticsearch Notes

## Apache Lucene

**Overview**: Free and open-source search engine software library, originally written in Java by Doug Cutting in 1999.

### Key Facts
- Released: 1999 (27+ years old)
- Current version: 10.5.0 (June 2026)
- License: Apache License 2.0
- Website: lucene.apache.org

### History
- Doug Cutting wrote Lucene in 1999 (his 5th search engine)
- Joined Apache Software Foundation's Jakarta family in September 2001
- Became top-level Apache project in February 2005
- Name comes from Doug Cutting's wife's middle name and her maternal grandmother's first name

### Features
- Full-text indexing and searching
- Fuzzy search based on edit distance (Levenshtein distance)
- Recommendation systems (MoreLikeThis class)
- Cross-platform (Java, Python, Ruby, C#, C++, PHP, Perl)

### Lucene-based Projects
- **Apache Solr** - Enterprise search server
- **Elasticsearch** - Distributed search and analytics engine
- **OpenSearch** - AWS fork of Elasticsearch
- **Apache Nutch** - Web crawling and HTML parsing
- **MongoDB Atlas Search** - Cloud-native enterprise search
- **CrateDB** - Distributed SQL database built on Lucene
- **Kinosearch** - Perl/C search engine

---

## Elasticsearch

**Overview**: Source-available search engine based on Apache Lucene, developed by Elastic NV.

### Key Facts
- Release: February 8, 2010 (16 years old)
- Current versions: 8.x (8.19.18), 9.x (9.4.3)
- License: Triple-licensed (Elastic License, SSPL, AGPL as of v8.16.0)
- Developer: Elastic NV (founded 2012)
- IPO: October 5, 2018 (NYSE)

### History
- Created by Shay Banon in 2004 as Compass
- Rewrote Compass to create Elasticsearch (2010)
- Company renamed from Elasticsearch to Elastic (March 2015)
- Acquired Swiftype (2017) - basis for App Search and Site Search
- License change controversy (2021) - relicensed from Apache 2.0 to SSPL/Elastic License
- AWS forked to create OpenSearch (2021)
- Added AGPL license option (August 2024) - back to open source

### Architecture
- Built on Apache Lucene
- JSON documents stored in indices
- Indices divided into primary shards with replicas
- Distributed, multitenant-capable
- Schema-free JSON documents
- HTTP web interface

### Features
- Full-text search
- Faceted search
- Real-time search
- Multitenancy
- Percolation (prospective search)
- NoSQL datastore (no distributed transactions)
- Security (TLS, RBAC)
- SIEM capabilities
- Machine learning

### Elastic Stack (ELK Stack)
- **Elasticsearch** - Search and analytics
- **Logstash** - Data collection and log parsing
- **Kibana** - Analytics and visualization
- **Beats** - Lightweight data shippers

### Official Clients
- Java, C# (.NET), PHP, Python, Ruby, Go, JavaScript, and others

### Market Position
- Most popular enterprise search engine (per DB-Engines ranking)

---

## Relevance to Context Engineering

Lucene and Elasticsearch represent the mature, battle-tested foundation for retrieval systems in AI applications:

1. **Lexical Retrieval**: Traditional inverted index-based search that remains relevant for exact matches, keyword search, and structured queries

2. **Hybrid Systems**: Modern context engineering combines vector search with lexical search (Elasticsearch supports both through plugins and dense vector support)

3. **Distributed Architecture**: Elasticsearch's sharding and replication patterns inform how to build scalable context retrieval systems

4. **Production Reliability**: Both projects have extensive operational history - patterns for monitoring, scaling, and failure handling apply to context engineering platforms

5. **Query Understanding**: Faceted search, percolation, and advanced query DSL demonstrate sophisticated retrieval strategies beyond simple vector similarity

---

## References
- https://en.wikipedia.org/wiki/Apache_Lucene
- https://en.wikipedia.org/wiki/Elasticsearch
- https://lucene.apache.org/
- https://www.elastic.co/elasticsearch/