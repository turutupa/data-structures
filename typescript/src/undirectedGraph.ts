type Edge<T> = [T, T, number];

export class Graph<T extends string | number> {
  // Stores each node and what node it is connected to plus the weight of the edge.
  // Example:
  //{
  //  a: { b: 1, c: 4},
  //  b: { a: 1, d: 5},
  //}
  private adjacencyList: Map<T, Map<T, number>> = new Map();

  private addVertex(value: T): void {
    if (!this.adjacencyList.has(value)) {
      this.adjacencyList.set(value, new Map());
    }
  }

  public addEdge(vertexA: T, vertexB: T, weight: number): void {
    this.addVertex(vertexA);
    this.addVertex(vertexB);

    this.adjacencyList.get(vertexA)!.set(vertexB, weight);
    this.adjacencyList.get(vertexB)!.set(vertexA, weight);
  }

  public addEdgeFrom(edges: Edge<T>[]) {
    for (const edge of edges) {
      this.addEdge(...edge);
    }
  }

  public getEdges(): Edge<T>[] {
    const edges: Edge<T>[] = [];
    const visitedEdges = new Set<string>();
    for (const [vertex, neighbors] of this.adjacencyList.entries()) {
      for (const [neighbor, weight] of neighbors.entries()) {
        const vertices: [T, T] = [vertex, neighbor].sort() as [T, T];
        const visitedKey = vertices.join('-');
        if (!visitedEdges.has(visitedKey)) {
          visitedEdges.add(visitedKey);
          edges.push([...vertices, weight]);
        }
      }
    }
    return edges;
  }
}
