import { Graph } from './undirectedGraph';

describe('Graph', () => {
  it('should add edges correctly', () => {
    const graph = new Graph<string>();
    graph.addEdge('a', 'b', 1);
    expect(graph.getEdges()).toEqual([['a', 'b', 1]]);
  });

  it('should handle edge cases', () => {
    // Write more test cases to cover different scenarios
  });
});
