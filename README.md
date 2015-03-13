Mesosphere Challenge
---

###Problem Specification

Design and implement an elevator control system. What data structures, interfaces and algorithms will you need? Your elevator control system should be able to handle a few elevators -- up to 16.

You can use the language of your choice to implement an elevator control system. In the end, your control system should provide an interface for:

* Querying the state of the elevators (what floor are they on and where they are going),

* receiving an update about the status of an elevator,

* receiving a pickup request,

* time-stepping the simulation.

For example, we could imagine in Scala an interface like this:

    trait ElevatorControlSystem {
      def status(): Seq[(Int, Int, Int)]
      def update(Int, Int, Int)
      def pickup(Int, Int)
      def step()
    }

Here we have chosen to represent elevator state as 3 integers:

    Elevator ID, Floor Number, Goal Floor Number

An update alters these numbers for one elevator. A pickup request is two integers:

    Pickup Floor, Direction (negative for down, positive for up)

This is not a particularly nice interface, and leaves some questions open. For example, the elevator state only has one goal floor; but it is conceivable that an elevator holds more than one person, and each person wants to go to a different floor, so there could be a few goal floors queued up. Please feel free to improve upon this interface!

The most interesting part of this challenge is the scheduling problem. The simplest implementation would be to serve requests in FCFS (first-come, first-served) order. This is clearly bad; imagine riding such an elevator! Please discuss how your algorithm improves on FCFS in your write-up.

Please provide a source tarball (or link to a GitHub repository) containing code in the language of your choice, as well as a README discussing your solution (and providing build instructions). The accompanying documentation is an important part of your submission. It counts to show your work.

----
###Exercise

####Prerequisites

Choose Go as the language to be used for this exercise as it is the most familiar language I have used with system programming in the past few months.  Follow this [official guide](https://golang.org/doc/install) to get Go installed.


####Install

Go does not come with a built-in package manager for managing its libraries.  We do have few libs here for this exercise.  We choose [godep](https://github.com/tools/godep) to keep our Go dependencies libraries, so we will have to restore them before we run any step.

    $ go get github.com/tools/godep
    $ godep restore

####Test

For all running all unit tests

    $ godep go test

Test whether tests have any race condition

    $ godep go test -race

####Run

We created one individual unit test `TestStep()` as a demo on how the Elevator Control System is simulated.

    $ godep go test -run TestStep

####Algorithm

The simplest implementation would be to serve requests in FCFS (first-come, first-served) order. This is clearly bad.  We need to do better!

[Destination dispatch](http://en.wikipedia.org/wiki/Destination_dispatch) is an optimization technique we can use when design for our elevator system assuming we have more than one elevator in the system.  This groups all pickup requests for the same destinations into the same elevator.  However, our pickup request was only given `Pickup Floor, Direction`, it does not contain the destination floor.   Given the limited time for this exercise, we only going to group all the same direction into the same elevator, reducing waiting and travel times.  Ideally we should also have a pick up request containing the destination.  The implementation detail can be found at `func Unload()`.

Another optimization technique we can add on group dispatch system is to divide our group of elevators into more than one group.  The first group goes to floor from 0 to said 20, the second elevator group skips all floor until after 20 to said 40 and the ground floor.   This optimization is designed to go very fast from first to the 20nd floor, making it very efficient. More on [this](http://www.tinyepiphany.com/2009/12/elevator-algorithms.html).  However, our implementation did not have this optimized design because we are not clear on the requirement given to `step()` function whether `step()` movement is stepping at each floor or stepping at each iteration of an elevator system.  The simulation of an elevator system can simply jump from floor 0 to floor 20 at each iteration because it is simulated, not floor-by-floor stepping.  This exercise do not have it clearly stated, hence we chose the latter.

Given a pickup request is made:

1. Added to the pickup request queue waiting to be process. This simulate the keypad pressing at the Destination dispatch system.
1. Each `step()` calls process the pickup request queue if it is not empty, calculate and allocate them to each elevator.  It allocates base on the criteria that:
    * if the elevator is at ground floor, and the elevator has no goal floors
    * they are going to the same direction and the elevator is on its path to pickup request's floor
1. Once the allocation has completed, it is actually then move the elevator given each elevator's goal floors queue.
    1. If an elevator has no goal floors :
        * if current floor is not at ground floor, move to ground floor.
        * if current is already at ground floor, stays in idle.
    1. If an elevator still has a goal floor:
        * move to that goal floor and change the current floor to that

This algorithm is at least better than FCFS because it takes pickup request's direction into considering whether an elevator should pick up the request or not.  It is also distributed in allocating pickup requests meaning that with more elevators in the system, the more "concurrently" our elevators can process these pickup requests. However, we do believe that if we have given more than the limited timeframe of 4 hours, we could also modify our current design to allow pickup requests to set destination floor along with grouping the elevators by serving floors.

####Assumptions and Limitations

* Each step (or tick) allows the elevator to reach the next floor next floor if it has any existing goal floors.
* Lack of true load-balance, first elevator will almost always pick up as much as it can, but it does added intelligence that it can allocate pickup requests with same direction as much as it can.

####Improvements

* Use goroutine for simulation, so it can be truly simulating the multi-elevators system having goroutine handling all the pickup requests.  And this can also better demostrated using Go's channel to communicate via message instead of using share objects.
* Elevators can be more intelligent by separating them by groups.
* Elevators can be more intelligent by grouping destination as much as it can into same elevators.
* Adding true load-balance mechanism to utilize as much elevator as possible while not always using the first elevator.
