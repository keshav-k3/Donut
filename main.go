package main

import (
	"fmt"
	"math"
	"time"
)

const (
	screenWidth  = 80
	screenHeight = 24
	donutWidth   = 3
	donutHeight  = 3
	delay        = 33 * time.Millisecond
)

func main() {
	A := 0.0
	B := 0.0
	for {
		screen := make([]byte, screenWidth*screenHeight)
		zBuffer := make([]float64, screenWidth*screenHeight)

		for i := range screen {
			screen[i] = ' '
			zBuffer[i] = 0
		}

		cosA, sinA := math.Cos(A), math.Sin(A)
		cosB, sinB := math.Cos(B), math.Sin(B)

		// Donut Rendering
		for theta := 0.0; theta < 2*math.Pi; theta += 0.07 {
			costheta := math.Cos(theta)
			sintheta := math.Sin(theta)

			for phi := 0.0; phi < 2*math.Pi; phi += 0.02 {
				cosphi := math.Cos(phi)
				sinphi := math.Sin(phi)

				// Torus Calculation
				x := (donutWidth + costheta) * cosphi
				y := (donutWidth + costheta) * sinphi
				z := sintheta

				// Y-axis Rotation
				x1 := x*cosB - z*sinB
				z1 := x*sinB + z*cosB

				// X-axis Rotation
				y1 := y*cosA - z1*sinA
				z2 := y*sinA + z1*cosA

				// Projection
				ooz := 1.0 / (6 + z2)
				projX := int(screenWidth/2 + 25*x1*ooz)
				projY := int(screenHeight/2 - 12*y1*ooz)

				index := projY*screenWidth + projX
				if index >= 0 && index < len(screen) && ooz > zBuffer[index] {
					zBuffer[index] = ooz
					luminance := costheta*0.5 + cosphi*0.5
					shadeIndex := int(math.Max(0, math.Min(11, luminance*8+2)))
					screen[index] = ".,-~:;=!*#$@"[shadeIndex]
				}
			}
		}

		// Magic
		fmt.Print("\x1b[H")
		for i := 0; i < len(screen); i++ {
			if i%screenWidth == 0 && i != 0 {
				fmt.Println()
			}
			fmt.Printf("%c", screen[i])
		}

		// Speed Adjustment
		A += 0.04
		B += 0.01
		time.Sleep(delay)
	}
}
