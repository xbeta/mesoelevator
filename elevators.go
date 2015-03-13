package mesoelevator

import(
  "sort"
)

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ById []*Elevator

func (a ById) Len() int           { return len(a) }
func (a ById) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ById) Less(i, j int) bool { return a[i].id < a[j].id }

type Elevators []*Elevator

// NOTE: not thread-safe, but it does not need to be thread-safe.
// It's simple for loop to initialize the list.  This should get call at the very beginning to setup
// the elevator list.
func NewElevators(size int)(Elevators){
  var elevators  Elevators
  for i := 1; i <= size; i++  {
    elevator := NewElevator(i)
    elevators = append(elevators, elevator)
  }
  return elevators
}

// It does a binary search to find the elevator element by its Id and return the index position
func (es Elevators) FindId(id int)(int){
  sort.Sort(ById(es)) //assuming that slice is not always sorted
  return sort.Search(len(es), func(i int) bool {
    return es[i].id >= id
  })
}

// NOTE: not thread-safe, but it does not need to be thread-safe.  It's a read-only function and
// should only be used for reporting purpose.
func (es Elevators) Flatten()([][]int){
  var slice [][]int
  for _, elevator := range es {
      // type assertion for slices from interface{} to int
      goals := make([]int, elevator.goalFloors.Size())
      for i, floor := range elevator.goalFloors.List() { goals[i] = floor.(int) }

      elevatorStatus := []int{elevator.id,  elevator.currentFloor, elevator.direction}

      //flatten goal floors into 1 slice and added
      elevatorStatus = append(elevatorStatus, goals...)
      slice = append(slice, elevatorStatus )
  }
  return slice
}
