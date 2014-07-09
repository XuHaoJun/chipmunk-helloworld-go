package main

import (
  "fmt"

  "github.com/vova616/chipmunk"
  "github.com/vova616/chipmunk/vect"
)

func main() {
  // vect.Vect is a 2D vector
  gravity := vect.Vect{0, -100}

  // Create an empty space.
  space := chipmunk.NewSpace()
  space.Gravity = gravity

  // Add a static line segment shape for the ground.
  // We'll make it slightly tilted so the ball will roll off.
  // We attach it to staticBody to tell Chipmunk it shouldn't be movable.
  staticBody := chipmunk.NewBodyStatic()
  ground := chipmunk.NewSegment(vect.Vect{-20, 5}, vect.Vect{20, -5}, 0)
  ground.SetFriction(1)
  staticBody.AddShape(ground)
  space.AddBody(staticBody)

  // Now let's make a ball that falls onto the line and rolls off.
  // First we need to make a Chipmunk's body to hold physical properties of object.
  // Then we attach collision shapes to the Chipmunk's body to give it size and shape.

  radius := vect.Float(5.0)
  mass := vect.Float(1.0)

  // Now we create the collision shape for the ball.
  // You can create multiple collsion shapes that point to the same body.
  // They will all be attached to the body and move around to follow it.
  ballShape := chipmunk.NewCircle(vect.Vector_Zero, float32(radius))
  ballShape.SetFriction(0.7)

  // The moment of inertia is like mass for rotation
  // Use the circleShape.Moment() function to help you approximate it.
  ballBody := chipmunk.NewBody(mass, ballShape.Moment(float32(mass)))
  ballBody.SetPosition(vect.Vect{0, 15})
  ballBody.AddShape(ballShape)

  // The space.AddBody() functions return the thing that you are adding.
  // It's convenient to create and add an object in one line.
  space.AddBody(ballBody)

  // Now that it's all set up, we simulate all the objects in the space by 
  // stepping sorward through time in small increments called steps.
  // It is *highly* recommended to use a fixed size time step.
  timeStep := vect.Float(1.0/60.0)
  for time := vect.Float(0); time < 2; time += timeStep{
    pos := ballBody.Position()
    vel := ballBody.Velocity()
    fmt.Printf(
      "Time is %5.2f. ballBody is at (%5.2f, %5.2f). It's velocity is (%5.2f, %5.2f)\n",
    time, pos.X, pos.Y, vel.X, vel.Y)
    space.Step(timeStep)
  }
}
