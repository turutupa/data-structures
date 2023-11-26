import { Graph } from './graph';

export class GraphPrinter<T> {
  graph: Graph<T>;

  constructor(graph: Graph<T>) {
    this.graph = graph;
  }

  print() {}
}
