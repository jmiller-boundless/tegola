package intersect

import (
	"fmt"

	"github.com/terranodo/tegola/container/list"
	ptList "github.com/terranodo/tegola/container/list/point/list"

	"github.com/terranodo/tegola/maths"
)

type Point struct {
	ptList.Pt

	Inward bool
	Seen   bool

	subject list.Sentinel
	region  list.Sentinel
}

func NewPt(pt maths.Pt, inward bool) *Point {
	return &Point{Pt: *ptList.NewPt(pt), Inward: inward}
}
func NewPoint(x, y float64, inward bool) *Point {
	return &Point{Pt: *ptList.NewPoint(x, y), Inward: inward}
}

func (i *Point) String() string {
	return fmt.Sprintf("Intersec{ X: %v, Y: %v, Inward: %v}", i.Pt.X, i.Pt.Y, i.Inward)
}
func (i *Point) AsSubjectPoint() *SubjectPoint { return (*SubjectPoint)(i) }
func (i *Point) AsRegionPoint() *RegionPoint   { return (*RegionPoint)(i) }

/*
func (i *Point) Walk() (w Walker) {
	var ele list.Elementer
	var ok bool
	if i.Inward {
		ele = i.subject.Next()
	}
	ele = i.region.Next()
	for w, ok = ele.(Walker); ele != nil && !ok; ele = ele.Next() {
	}
	if ele != nil {
		return w
	}
	return nil
}
*/

// RegionPoint causes an intersect point to "act" like a region point so that it can be inserted into a region list.
type RegionPoint Point

func (i *RegionPoint) Prev() list.Elementer { return i.region.Prev() }
func (i *RegionPoint) Next() list.Elementer { return i.region.Next() }
func (i *RegionPoint) SetNext(e list.Elementer) list.Elementer {
	return i.region.SetNext(e)
}
func (i *RegionPoint) SetPrev(e list.Elementer) list.Elementer {
	return i.region.SetPrev(e)
}
func (i *RegionPoint) List() *list.List                { return i.region.List() }
func (i *RegionPoint) SetList(l *list.List) *list.List { return i.region.SetList(l) }
func (i *RegionPoint) AsSubjectPoint() *SubjectPoint {
	return (*SubjectPoint)(i)
}
func (i *RegionPoint) AsIntersectPoint() *Point { return (*Point)(i) }
func (i *RegionPoint) Point() maths.Pt          { return i.Pt.Point() }

// SubjectPoing causes an intersect point to "act" like a subject point so that it can be inserted into a subject list.
type SubjectPoint Point

func (i *SubjectPoint) Prev() list.Elementer { return i.subject.Prev() }
func (i *SubjectPoint) Next() list.Elementer { return i.subject.Next() }
func (i *SubjectPoint) SetNext(e list.Elementer) list.Elementer {
	return i.subject.SetNext(e)
}
func (i *SubjectPoint) SetPrev(e list.Elementer) list.Elementer {
	return i.subject.SetPrev(e)
}
func (i *SubjectPoint) List() *list.List                { return i.subject.List() }
func (i *SubjectPoint) SetList(l *list.List) *list.List { return i.subject.SetList(l) }
func (i *SubjectPoint) AsRegionPoint() *RegionPoint {
	return (*RegionPoint)(i)
}
func (i *SubjectPoint) AsIntersectPoint() *Point { return (*Point)(i) }
func (i *SubjectPoint) Point() maths.Pt          { return i.Pt.Point() }
