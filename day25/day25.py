from igraph import Graph

with open('input.txt') as f:
    lines = f.readlines()

split_lines = [l.split(':') for l in lines]
G = {v: e.split() for v, e in split_lines}

# Create the graph
graph = Graph.ListDict(G)

# Find the cut
cut = graph.mincut()

print('Cut:', cut.cut)
print('ans:', len(cut.partition[0]) * len(cut.partition[1]))

#print out the cuts
for e in cut.cut:
    source = graph.es[e].source
    target = graph.es[e].target
    print("cut ", graph.vs[source]['name'], "-", graph.vs[target]['name'])

