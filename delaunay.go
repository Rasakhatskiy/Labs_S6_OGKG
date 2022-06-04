package main

import "errors"

// Тріангуляція Делоне

type Delaunay struct {
	dc DelaunayCel
	g  Graph
}

func (d *Delaunay) create(points []Point) error {
	err := d.dc.init(points)
	if err != nil {
		return err
	}

	d.initGraph()

	for i := 1; i < len(d.dc.vertices); i++ {

	}

	return nil
}

type PositionType int

const (
	interior PositionType = 0
	boundary              = 1
	outside               = 2
)

type Direction int

const (
	positive  Direction = 0
	negative            = 1
	collinear           = 2
)

func (dc DelaunayCel) getDirection(p Point, srcPointId, dstPointId int) Direction {
	if srcPointId > 0 && dstPointId > 0 {
		// нормальне джерело і звичайна точка призначення
		// напрямок негативний - точка не в трикутнику
		// напрямок позитивний - точка може бути в трикутнику (потрібно перевірити інші ребра)
		// напрямок колінеарний - точка знаходиться на відрізку ребра трикутника
		return dc.vertices[srcPointId-1].point.getDirectionTo(dc.vertices[dstPointId-1].point, p)
	}

	if (srcPointId > 0 && dstPointId == -2) {
		// якщо p знаходиться вище вихідної точки або на тому ж рівні і праворуч від вихідної точки,
		// тоді напрямок негативний, інакше позитивний
		//
		//        -2 x                     -2 x
		//            \     x p                \
		//             \   /                    \
		//              \ /                      \
		//        source x                 source x ---- x p
		if p.greaterThan(dc.vertices[srcPointId-1].point) {
			return negative
		} else {
			return positive
		}
	}

}

func (d Delaunay) getPositionType(p Point, nodeIndex int) {
	// індекси точок трикутника
	vertices := d.g.nodes[nodeIndex].verticesIds

}

func (d *Delaunay) findNode(p Point) int {
	currentIndex := 0
	for !d.g.nodes[currentIndex].isLeaf() {
		children := d.g.nodes[currentIndex].childrenNodesId
		foundId := -1
		for i, child := range children {

		}
	}

}

/*
int delaunay::find_node(util::point point) const
{
    int current_index = 0;
    while (!m_graph[current_index].leaf()) {
        const auto& children = m_graph[current_index].children();
        auto it = std::find_if(children.cbegin(), children.cend(), [&](int child_index) {
                      return get_position(point, child_index) != position::outside;
                  });
        // there should be a triangle that constains the point at each level
        assert(it != children.end());
        current_index = *it;
    }

    return current_index;
}
*/
func (d *Delaunay) addPoint(i int) {
	point := d.dc.vertices[i].point
	// get node index that contains the point

}

/*
void delaunay::add_point(int point_index)
{
    auto point = m_dcel.vertex(point_index).point();
    // get node index that contains the point
    int node_index = find_node(point);
    auto position = get_position(point, node_index);
    assert(position != position::outside);

    if (position == position::strictly_interior) {
        // point is strictly in the triangle
        split_triangle_interior(point_index, node_index);
    }
    else {
        // point is on the triangle edge
        split_triangle_boundary(point_index, node_index);
    }
}
*/

// initGraph додаємо першу ноду на граф - найбільший трикутник (1,-2,-1) що містить всі точки
// і представляє собою першу внутрішню грань 1
func (d *Delaunay) initGraph() {
	d.g.add(Node{
		verticesIds:     []int{1, -2, -1},
		childrenNodesId: []int{},
		faceId:          1,
	})
}

type DelaunayCel struct {
	vertices []Vertex
	edges    []Edge
	faces    []Face
}

// setHighestFirst знаходимо найбільшу вершину і ставимо першою
func (dc *DelaunayCel) setHighestFirst() error {
	if len(dc.edges) != 0 || len(dc.faces) != 0 || len(dc.vertices) == 0 {
		return errors.New("setHighestFirst invalid input")
	}

	maxVertex := dc.vertices[0]
	maxId := 0
	for i, vertex := range dc.vertices {
		if vertex.greaterThan(maxVertex) {
			maxVertex = vertex
			maxId = i
		}
	}

	dc.vertices[0], dc.vertices[maxId] = dc.vertices[maxId], dc.vertices[0]

	return nil
}

func (dc *DelaunayCel) init(points []Point) error {
	if len(points) < 3 {
		return errors.New("minimum 3 points needed")
	}
	for _, point := range points {
		dc.vertices = append(dc.vertices, Vertex{point, -1})
	}

	err := dc.setHighestFirst()
	if err != nil {
		return err
	}

	dc.vertices[0].incidentEdgeId = 1

	dc.edges = append(dc.edges, Edge{1, 4, 3, 2, 1})  // 1-4
	dc.edges = append(dc.edges, Edge{-2, 6, 1, 3, 1}) // 2-6
	dc.edges = append(dc.edges, Edge{-1, 5, 2, 1, 1}) // 3-5
	dc.edges = append(dc.edges, Edge{-2, 1, 6, 5, 0}) // 4-1
	dc.edges = append(dc.edges, Edge{1, 3, 4, 6, 0})  // 5-3
	dc.edges = append(dc.edges, Edge{-1, 2, 5, 4, 0}) // 6-2

	// зовнішня грань
	dc.faces = append(dc.faces, Face{4})

	// перша внутрішня грань
	dc.faces = append(dc.faces, Face{1})

	return nil
}

type Node struct {
	verticesIds     []int
	childrenNodesId []int
	faceId          int
}

func (n Node) isLeaf() bool {
	return len(n.childrenNodesId) == 0
}

type Graph struct {
	// map: face id -> node id
	// last node id with face id
	faceToNode map[int]int
	nodes      []Node
}

func (g *Graph) add(n Node) {
	if n.faceId > 0 {
		g.faceToNode[n.faceId] = len(g.nodes)
	}
	g.nodes = append(g.nodes, n)
}
