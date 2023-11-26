// Directed Graph data structure. Uses adjacency list as it will perform better in most cases
// reference: https://ricardoborges.dev/data-structures-in-typescript-graph

// TODO: finish implementation - realized I would also implement adjacency list for
// undirected graph so left this unfinished as this was not high priority at the given
// time.

export class DirectedGraph<T> {
  vertices: Map<T, Vertex<T>> = new Map();

  addNode(value: T) {}

  addEdge(source: T, destination: T) {}
}

export class Vertex<T> {
  value: T;
  adjacents: Vertex<T>[];
  description: string;
  comparator: (a: T, b: T) => number;

  constructor(value: T) {
    this.value = value;
    this.adjacents = [];
    this.description = '';
  }

  addAdjacent(vertex: Vertex<T>): void {
    if (!this.adjacents.includes(vertex)) this.adjacents.push(vertex);
  }

  removeAdjacent(value: T) {
    let index = 0;
    for (let adjacent of this.adjacents) {
      if (adjacent.comparator(adjacent.value, value) === 0) {
        return this.adjacents.splice(index, 1).pop();
      }
      index++;
    }
    return null;
  }
}

export class Edge<T> {}
