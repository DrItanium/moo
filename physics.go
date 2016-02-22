// physics engine
package moo

import "github.com/DrItanium/moo/cseries"

type physicsModel int16

const (
	// models
	modelGameWalking physicsModel = iota
	modelGameRunning
	numberOfPhysicsModels
)

type physicsConstants struct {
	maximum_forward_velocity, maximum_backward_velocity, maximum_perpendicular_velocity cseries.Fixed
	acceleration, deceleration, airborne_deceleration                                   cseries.Fixed /* forward, backward and perpendicular */
	gravitational_acceleration, climbing_acceleration, terminal_velocity                cseries.Fixed
	external_deceleration                                                               cseries.Fixed

	angular_acceleration, angular_deceleration, maximum_angular_velocity, angular_recentering_velocity cseries.Fixed
	fast_angular_velocity, fast_angular_maximum                                                        cseries.Fixed /* for head movements */
	maximum_elevation                                                                                  cseries.Fixed /* positive and negative */
	external_angular_deceleration                                                                      cseries.Fixed

	/* step_length is distance between adjacent nodes in the actorÕs phase */
	step_delta, step_amplitude                                cseries.Fixed
	radius, height, dead_height, camera_height, splash_height cseries.Fixed

	half_camera_separation cseries.Fixed
}
