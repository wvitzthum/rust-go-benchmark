use std::collections::VecDeque;

pub struct Graph {
    adj_list: Vec<Vec<usize>>, // Adjacency list representation of the graph
}

impl Graph {
    fn new(num_vertices: usize) -> Self {
        Graph {
            adj_list: vec![Vec::new(); num_vertices],
        }
    }

    fn add_edge(&mut self, u: usize, v: usize) {
        // Add an undirected edge between nodes u and v
        self.adj_list[u].push(v);
        self.adj_list[v].push(u);
    }

    pub fn bfs(&self, start: usize) -> Vec<usize> {
        let mut visited = vec![false; self.adj_list.len()]; // Keep track of visited nodes
        let mut queue = VecDeque::new(); // FIFO queue for BFS traversal
        let mut traversal = Vec::new(); // Store the traversal order

        visited[start] = true;
        queue.push_back(start);

        while let Some(node) = queue.pop_front() {
            traversal.push(node);
            // Visit each neighbor of the current node that has not been visited yet
            for neighbor in &self.adj_list[node] {
                if !visited[*neighbor] {
                    visited[*neighbor] = true;
                    queue.push_back(*neighbor);
                }
            }
        }

        traversal
    }
}

pub fn example_graph() -> Graph {
    let mut graph = Graph::new(6);
    graph.add_edge(0, 1);
    graph.add_edge(0, 2);
    graph.add_edge(1, 3);
    graph.add_edge(2, 3);
    graph.add_edge(2, 4);
    graph.add_edge(3, 4);
    graph.add_edge(3, 5);
    return graph;
}