# gosearch

### Diagram of the Architecture

```mermaid
graph LR
    A[Client Application] --> B[API Layer - HTTP Server];
    B --> C{Request Type?};
    C -- Indexing Request --> D[Document Ingestion];
    D --> E[Text Processing];
    E --> F{Distributed?};
    F -- Yes --> G[Sharding & Routing];
    G --> H[Data Node 1];
    G --> I[Data Node 2];
    H --> J[Inverted Index];
    I --> K[Inverted Index];
    J -- Update --> L[Persistence Disk];
    K -- Update --> M[Persistence Disk];
    F -- No --> J;
    J -- Update --> L;
    C -- Search Query --> N[Query Parsing];
    N --> O[Text Processing];
    O --> P{Distributed?};
    P -- Yes --> Q[Routing to Data Nodes];
    Q --> H;
    Q --> I;
    H --> R[Inverted Index Lookup];
    I --> S[Inverted Index Lookup];
    R --> T[Ranking Algorithm];
    S --> U[Ranking Algorithm];
    T --> V[Result Merging & Ranking];
    U --> V;
    V --> B;
    P -- No --> R;
    R --> T;
    T --> B;
    C -- CRUD Operations --> W{Distributed?};
    W -- Yes --> X[Routing & Coordination];
    X --> H;
    X --> I;
    H -- CRUD --> L;
    I -- CRUD --> M;
    W -- No --> J;
    J -- CRUD --> L;
    B --> A;
    subgraph Distributed Cluster
        direction LR
        H -- Data Storage & Search --> J
        I -- Data Storage & Search --> K
        subgraph Master Node
            AA[Cluster Management]
            BB[Node Discovery]
        end
        AA -- Manages --> H
        AA -- Manages --> I
        BB -- Connects --> H
        BB -- Connects --> I
        B -- Routes to --> AA
        B -- Routes to --> BB
        B -- Routes to --> H
        B -- Routes to --> I
    end
    style A fill:#fff
    style B fill:#f9f,stroke:#333,stroke-width:2px
    style C fill:#ccf
    style D fill:#fff
    style E fill:#fff
    style F fill:#ccf
    style G fill:#fff
    style H fill:#fff
    style I fill:#fff
    style J fill:#fff
    style K fill:#fff
    style L fill:#fff
    style M fill:#fff
    style N fill:#fff
    style O fill:#fff
    style P fill:#ccf
    style Q fill:#fff
    style R fill:#fff
    style S fill:#fff
    style T fill:#fff
    style U fill:#fff
    style V fill:#fff
    style W fill:#ccf
    style X fill:#fff
    style AA fill:#fff
    style BB fill:#fff
```

#### Erläuterung des Diagramms:

1. Client Application: Dies ist die Anwendung oder der Benutzer, der mit der Suchmaschine interagiert.
2. API Layer - Go HTTP Server: Dies ist die Schnittstelle der Suchmaschine, implementiert als HTTP-Server in Go (z.B. mit actix-web, warp oder hyper). Er empfängt Anfragen zum Indizieren, Suchen und für CRUD-Operationen.
3. Request Type?: Eine Entscheidungsstelle, die die Art der eingehenden Anfrage (Indizierung, Suche oder CRUD) unterscheidet.
4. Indexing Request:
    - Document Ingestion: Hier werden die zu indexierenden Dokumente entgegengenommen.
    - Text Processing - Go: Die Dokumente werden in Go verarbeitet (Tokenisierung, Stoppwortentfernung, Stemming/Lemmatisierung).
    - Distributed?: Eine Entscheidung, ob die Architektur verteilt ist.
    - Sharding & Routing: Wenn verteilt, werden die Dokumente auf verschiedene Datenknoten verteilt (Sharding).
    - Data Node 1/2: Repräsentieren einzelne Knoten im verteilten System.
    - Inverted Index - Go: Der Invertierte Index wird in Go auf den jeweiligen Datenknoten erstellt oder aktualisiert.
    - Persistence (Disk): Der Invertierte Index wird auf Festplatte gespeichert.
5. Search Query:
    - Query Parsing: Die Suchanfrage wird analysiert.
    - Text Processing - Go: Die Suchbegriffe werden analog zu den Dokumenten verarbeitet.
    - Distributed?: Entscheidung, ob die Suche über mehrere Knoten verteilt werden muss.
    - Routing to Data Nodes: Die Anfrage wird an die relevanten Datenknoten weitergeleitet.
    - Inverted Index Lookup: Die Datenknoten suchen im lokalen Invertierten Index nach passenden Dokumenten.
    - Ranking Algorithm - Go: Die Relevanz der gefundenen Dokumente wird mit einem Ranking-Algorithmus (z.B. TF-IDF oder BM25) in Go berechnet.
    - Result Merging & Ranking: Die Ergebnisse von verschiedenen Datenknoten werden zusammengeführt und erneut nach Relevanz sortiert.
6. CRUD Operations:
    - Distributed?: Entscheidung, ob die Operation im verteilten System durchgeführt werden muss.
    - Routing & Coordination: Die Operation wird an die entsprechenden Datenknoten geleitet und koordiniert.
    - CRUD: Die eigentlichen Erstell-, Lese-, Aktualisierungs- oder Löschoperationen auf dem Invertierten Index und der Persistenzebene.
7. Distributed Cluster:
    - Master Node: Verantwortlich für die Verwaltung des Clusters (Hinzufügen/Entfernen von Knoten, Überwachung des Zustands).
    - Node Discovery: Mechanismus, mit dem sich die Knoten im Cluster finden und miteinander kommunizieren können.
    - Data Node 1/2: Speichern die Daten und führen Suchoperationen aus.
8. Persistence (Disk): Speichert den Invertierten Index dauerhaft.

Dieses Diagramm bietet eine High-Level-Übersicht über die Architektur einer Elasticsearch-ähnlichen Suchmaschine, die in Go entwickelt werden könnte. Die Details der einzelnen Komponenten und deren Implementierung können je nach spezifischen Anforderungen variieren.
