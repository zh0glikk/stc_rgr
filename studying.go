package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"github.com/sirupsen/logrus"
	"math"
	"time"
)

const amountOfInputs = 5

var expectedValues []float64
var dispersions []float64

var tR = map[int]float64{3: 2.353363, 4: 2.131847,5:2.015048,6:1.943180,7:1.894579,8:1.859548,9:1.833113,10:1.812461}

func logKey(length int) (string, []time.Duration) {
	var result string
	var durations []time.Duration

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for i := 0; i < length; i++ {
		before := time.Now()
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		dur := time.Since(before)
		durations = append(durations, dur)

		logrus.WithFields(logrus.Fields{
			"char" : string(char),
			"time" : dur,
		}).Info("Logging")

		if key == keyboard.KeyEnter {
			break
		}

		result += string(char)
	}

	return result, durations
}

func calcExpectedValue(dur []time.Duration, i int) float64{
	sum := 0.0

	for j := 0; j < len(dur); j++ {
		if i != j {
			sum += float64(dur[j])
		}
	}

	return sum / float64(len(dur) - 1)
}

func calcDispersion(dur []time.Duration, expectedValue float64, j int) float64 {
	tmp := 0.0

	for i := 0; i < len(dur); i++ {
		if i != j {
			tmp += math.Pow(float64(dur[i]) - expectedValue, 2)
		}
	}
	return tmp / float64(len(dur) - 2)
}

func calcStudentCoefficient(expectedValues []float64, dispersions []float64, durations [][]time.Duration) {
	studentCoefficient := 0.0

	for i := 0; i < len(durations[0]); i++ {
		for j := 0; j < len(durations); j++ {
			studentCoefficient = math.Abs((float64(durations[j][i]) - expectedValues[j]) / math.Sqrt(dispersions[j]))

			if studentCoefficient > tR[len(durations)-2] {
				durations[j][i] = 0
			}
		}
	}

	for i := range durations {
		logrus.WithFields(logrus.Fields{
			"dur" : durations[i],
		}).Info("duration")
	}
}

func calcReference(dur [][]time.Duration) ([]float64, []float64){
	passwordLength := len(dur[0])

	var tMin = make([]float64, passwordLength)
	var tMax = make([]float64, passwordLength)

	var tempI []int
	var mE []float64

	for i := 0; i < passwordLength; i++ {
		temp := 0.0
		t := 0

		for j := 0; j < len(dur); j++ {
			if float64(dur[j][i]) != 0.0 {
				temp += float64(dur[j][i])
				t += 1
			}
		}
		tempI = append(tempI, t)
		mE = append(mE, temp / float64(t))
	}

	var sE []float64

	for i := 0; i < passwordLength; i++ {
		temp := 0.0
		t := 0

		for j := 0; j < len(dur); j++ {
			if float64(dur[j][i]) != 0.0 {
				temp += math.Pow(float64(dur[j][i]) - mE[i], 2)
				t += 1
			}
		}
		sE = append(sE, math.Sqrt(temp / float64(t)))
	}

	logrus.WithFields(logrus.Fields{
		"tempI" : tempI,
		"mE" : mE,
		"sE" : sE,
	}).Info("middle")


	for i := 0; i < len(mE); i++ {
		m := mE[i]
		t := tR[tempI[i]-1]
		s := sE[i]

		ts := t * s

		max := m + ts
		min := m - ts

		tMin[i] = min
		tMax[i] = max
	}

	logrus.WithFields(logrus.Fields{
		"tMax" : tMax,
		"tMin" : tMin,
	}).Info("ref")


	return tMin, tMax
}


func Study(userPassword string) ([]float64, []float64) {
	var password string
	var dur [][]time.Duration

	dur = make([][]time.Duration, amountOfInputs)

	fmt.Println("Now would be program learning\n" +
		"You should enter your password several times.\n" +
		"Press 'Enter' after input ")

	for i := 0; i < amountOfInputs; {
		fmt.Println("Input your password: ")

		password, dur[i] = logKey(len(userPassword))

		expectedValue := calcExpectedValue(dur[i], i)

		expectedValues = append(expectedValues, expectedValue)

		if password == userPassword {
			i += 1
		} else {
			fmt.Println("Wrong pass")
		}
	}

	logrus.WithFields(logrus.Fields{
		"exp Values" : expectedValues,
	}).Info("ExpV")

	for i := range expectedValues {
		dispersion := calcDispersion(dur[i], expectedValues[i], i)

		dispersions = append(dispersions, dispersion)
	}

	logrus.WithFields(logrus.Fields{
		"dispersions" : dispersions,
	}).Info("Dispersions")

	logrus.WithFields(logrus.Fields{
		"durations" : dur,
	}).Info("durations")

	calcStudentCoefficient(expectedValues, dispersions, dur)

	return calcReference(dur)
}

